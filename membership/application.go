package membership

import "github.com/google/uuid"

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	membership := NewMembershipBuilder()
	id := uuid.NewString()

	membership.SetID(id).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType)

	app.repository.data[id] = membership.GetMembership()
	//app.repository.data[time.Now().String()] =
	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{}, nil
}
func (app *Application) Delete(id string) error {
	return nil
}
