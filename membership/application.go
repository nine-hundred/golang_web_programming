package membership

import "github.com/google/uuid"

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
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
		AddMembership(membership)
	if err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	newMembership, err := NewMembershipBuilder().
		SetID(request.ID).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType).GetMembership()
	if err != nil {
		return UpdateResponse{}, err
	}
	app.repository.data[newMembership.ID] = newMembership

	return UpdateResponse{
		ID:             newMembership.ID,
		UserName:       newMembership.UserName,
		MembershipType: newMembership.MembershipType,
	}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}
