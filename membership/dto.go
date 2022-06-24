package membership

type CreateRequest struct {
	UserName       string `valid:"required"`
	MembershipType string `valid:"required,membershipType"`
}

type CreateResponse struct {
	ID             string `json:"id"`
	MembershipType string `json:"membershipType"`
}

type UpdateRequest struct {
	ID             string `valid:"required"`
	UserName       string `valid:"required"`
	MembershipType string `valid:"required,membershipType"`
}

type UpdateResponse struct {
	ID             string
	UserName       string
	MembershipType string
}

type ReadRequest struct {
	Limit  int
	Offset int
}

type ReadResponse struct {
	ID             string `valid:"required"`
	UserName       string `valid:"required"`
	MembershipType string `valid:"required,membershipType"`
}
