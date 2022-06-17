package membership

import "errors"

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) AddMembership(m Membership) (Membership, error) {
	for _, membership := range r.data {
		if membership.UserName == m.UserName {
			return Membership{}, errors.New("already existed name")
		}
	}
	r.data[m.ID] = m
	return m, nil
}

func (r *Repository) UpdateMembership(m Membership) (Membership, error) {
	for _, membership := range r.data {
		if membership.UserName == m.UserName {
			return Membership{}, errors.New("already existed name")
		}
	}
	return m, nil
}
