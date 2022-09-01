package membership

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"net/http"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
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

	membership := membershipBuilder.GetMembership()
	if err != nil {
		return CreateResponse{}, err
	}

	_, err = app.repository.
		AddMembership(*membership)
	if err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{
		Code:           http.StatusCreated,
		Message:        http.StatusText(http.StatusCreated),
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return UpdateResponse{}, err
	}
	newMembership := NewMembershipBuilder().
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
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             newMembership.ID,
		UserName:       newMembership.UserName,
		MembershipType: newMembership.MembershipType,
	}, nil
}

func (app *Application) Delete(id string) error {
	if id == "" {
		return errors.New("there is no id")
	}
	err := app.repository.DeleteMembership(id)
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
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (app *Application) ReadAll(request ReadAllRequest) ([]ReadResponse, error) {
	memberships, _ := app.repository.ReadAllMemberships(0, 0)

	//memberships = splitMemberships(request.Limit, request.Offset, memberships)
	res := make([]ReadResponse, 0)
	for _, membership := range memberships {
		res = append(res, ReadResponse{
			Code:           http.StatusOK,
			Message:        http.StatusText(http.StatusOK),
			ID:             membership.ID,
			UserName:       membership.UserName,
			MembershipType: membership.MembershipType,
		})
	}
	return res, nil
}
