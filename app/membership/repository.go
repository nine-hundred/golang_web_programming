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

func (r *Repository) DeleteMembership(membership string) error {
	if _, ok := r.data[membership]; ok {
		delete(r.data, membership)
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

func (r *Repository) ReadAllMemberships(limit int, offset int) (memberships []Membership, err error) {
	for _, membership := range r.data {
		memberships = append(memberships, membership)
	}
	sort.Slice(memberships, func(i, j int) bool {
		return memberships[i].ID > memberships[j].ID
	})
	memberships = splitMemberships(limit, offset, memberships)
	return memberships, nil
}

func (r *Repository) ReadMembershipByName(name string) (Membership, error) {
	for _, membership := range r.data {
		if membership.UserName == name {
			return membership, nil
		}
	}
	return Membership{}, errors.New("there is no membership")
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
