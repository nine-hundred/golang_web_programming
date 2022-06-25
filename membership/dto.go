package membership

type CreateRequest struct {
	UserName       string `valid:"required"`
	MembershipType string `valid:"required,membershipType"`
}

type CreateResponse struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	ID             string `json:"id"`
	MembershipType string `json:"membership_type"`
}

type UpdateRequest struct {
	ID             string `valid:"required"`
	UserName       string `valid:"required"`
	MembershipType string `valid:"required,membershipType"`
}

type UpdateResponse struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type ReadRequest struct {
	Limit  int
	Offset int
}

type ReadResponse struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	ID             string `json:"id"valid:"required"`
	UserName       string `json:"user_name"valid:"required"`
	MembershipType string `json:"membership_type"valid:"required,membershipType"`
}

type DeleteResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
