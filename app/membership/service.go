package membership

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateMembership(req CreateRequest) CreateResponse {
	id := uuid.NewString()
	membership := NewMembershipBuilder().SetID(id).
		SetUserName(req.UserName).
		SetMembershipType(req.MembershipType).
		GetMembership()

	res, err := s.repository.AddMembership(*membership)
	if err != nil {
		return CreateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return CreateResponse{
		Code:           http.StatusCreated,
		Message:        http.StatusText(http.StatusCreated),
		ID:             res.ID,
		MembershipType: res.MembershipType,
	}
}

func (s *Service) UpdateMembership(req UpdateRequest) UpdateResponse {
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return UpdateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	membership := NewMembershipBuilder().SetID(req.ID).
		SetUserName(req.UserName).
		SetMembershipType(req.MembershipType).
		GetMembership()

	res, err := s.repository.UpdateMembership(*membership)
	if err != nil {
		return UpdateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return UpdateResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             res.ID,
		UserName:       res.UserName,
		MembershipType: res.MembershipType,
	}
}

func (s *Service) DeleteMembership(id string) DeleteResponse {
	err := s.repository.DeleteMembership(id)
	if err != nil {
		return DeleteResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return DeleteResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}

func (s *Service) ReadMembershipById(id string) ReadResponse {
	res, err := s.repository.ReadMembership(id)
	if err != nil {
		return ReadResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return ReadResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             res.ID,
		UserName:       res.UserName,
		MembershipType: res.MembershipType,
	}
}

func (s *Service) ReadAllMembership(req ReadAllRequest) ReadAllResponse {

	limit, err := parseLimit(&req)
	if err != nil {
		return ReadAllResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	offset, err := parseOffset(&req)
	if err != nil {
		return ReadAllResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	res, err := s.repository.ReadAllMemberships(limit, offset)
	if err != nil {
		return ReadAllResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return ReadAllResponse{
		Code:         http.StatusOK,
		Message:      http.StatusText(http.StatusOK),
		ReadResponse: res,
	}
}

func (s *Service) FindMemebershipByName(name string) (Membership, error) {
	membership, err := s.repository.ReadMembershipByName(name)
	if err != nil {
		return Membership{}, err
	}
	return membership, nil
}

func parseLimit(req *ReadAllRequest) (int, error) {
	if req.Limit == "" {
		return 0, nil
	}
	limit, err := strconv.Atoi(req.Limit)
	if err != nil {
		return 0, errors.New("wrong limit type")
	}
	return limit, nil
}

func parseOffset(req *ReadAllRequest) (int, error) {
	if req.Offset == "" {
		return 0, nil
	}
	offset, err := strconv.Atoi(req.Offset)
	if err != nil {
		return 0, errors.New("wrong offset type")
	}
	return offset, nil
}

func CheckIdAndPw(name string, pw string) bool {
	if name != pw {
		return false
	}
	return true
}
