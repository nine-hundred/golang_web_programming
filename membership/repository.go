package membership

import (
	"errors"
	"sort"
)

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
		if membership.ID == m.ID {
			continue
		}
		if membership.UserName == m.UserName {
			return Membership{}, errors.New("already existed name")
		}
	}
	r.data[m.ID] = m
	return m, nil
}

func (r *Repository) DeleteMembership(membership Membership) error {
	if _, ok := r.data[membership.ID]; ok {
		delete(r.data, membership.ID)
		return nil
	}
	return errors.New("not existed id")
}

func (r *Repository) ReadMembership(id string) (Membership, error) {
	if _, ok := r.data[id]; ok {
		return r.data[id], nil
	}
	return Membership{}, errors.New("there is no user id")
}

func (r *Repository) ReadAllMemberships() (memberships []Membership) {
	for _, membership := range r.data {
		memberships = append(memberships, membership)
	}
	sort.Slice(memberships, func(i, j int) bool {
		return memberships[i].ID > memberships[j].ID
	})
	return memberships
}
