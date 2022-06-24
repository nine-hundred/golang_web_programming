package membership

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	StartValidator()
	return &Application{repository: repository}
}

func StartValidator() {
	govalidator.TagMap["membershipType"] = govalidator.Validator(func(str string) bool {
		if str == "payco" {
			return true
		}
		if str == "naver" {
			return true
		}
		if str == "toss" {
			return true
		}
		return false
	})
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	_, err := govalidator.ValidateStruct(request)

	if err != nil {
		return CreateResponse{}, err
	}
	membershipBuilder := NewMembershipBuilder()
	id := uuid.NewString()

	membershipBuilder.SetID(id).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType)

	membership, err := membershipBuilder.GetMembership()
	if err != nil {
		return CreateResponse{}, err
	}

	_, err = app.repository.
		AddMembership(*membership)
	if err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return UpdateResponse{}, err
	}
	newMembership, err := NewMembershipBuilder().
		SetID(request.ID).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType).
		GetMembership()
	if err != nil {
		return UpdateResponse{}, err
	}

	_, err = app.repository.UpdateMembership(*newMembership)
	if err != nil {
		return UpdateResponse{}, err
	}

	return UpdateResponse{
		ID:             newMembership.ID,
		UserName:       newMembership.UserName,
		MembershipType: newMembership.MembershipType,
	}, nil
}

func (app *Application) Delete(id string) error {
	if id == "" {
		return errors.New("there is no id")
	}
	m := app.repository.data[id]
	membership, err := NewMembershipBuilder().
		SetID(m.ID).
		SetUserName(m.UserName).
		SetMembershipType(m.MembershipType).
		GetMembership()
	if err != nil {
		return err
	}
	err = app.repository.DeleteMembership(*membership)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) Read(id string) (ReadResponse, error) {
	membership, err := app.repository.ReadMembership(id)
	if err != nil {
		return ReadResponse{}, err
	}

	return ReadResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (app *Application) ReadAll(request ReadRequest) ([]ReadResponse, error) {
	memberships := app.repository.ReadAllMemberships()
	memberships = splitMemberships(request.Limit, request.Offset, memberships)
	res := make([]ReadResponse, 0)
	for _, membership := range memberships {
		res = append(res, ReadResponse{
			ID:             membership.ID,
			UserName:       membership.UserName,
			MembershipType: membership.MembershipType,
		})
	}
	return res, nil
}

func splitMemberships(limit int, offset int, memberships []Membership) []Membership {
	if offset == 0 && limit == 0 {
		return memberships
	}
	membershipSlice := make([][]Membership, 0)
	j := 0
	for i := 0; i < len(memberships); i += limit {
		j += limit
		if j > len(memberships) {
			j = len(memberships)
		}
		membershipSlice = append(membershipSlice, memberships[i:j])
	}
	if offset >= len(membershipSlice) {
		return []Membership{}
	}
	return membershipSlice[offset]
}
