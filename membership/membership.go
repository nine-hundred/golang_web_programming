package membership

import "errors"

type Membership struct {
	ID             string
	UserName       string
	MembershipType string
}

func (m *Membership) SetID(id string) *Membership {
	m.ID = id
	return m
}

func (m *Membership) SetUserName(userName string) *Membership {
	m.UserName = userName
	return m
}

func (m *Membership) SetMembershipType(membershipType string) *Membership {
	m.MembershipType = membershipType
	return m
}

func (m *Membership) GetMembership() (Membership, error) {
	if err := m.validateMembership(); err != nil {
		return Membership{}, err
	}
	return Membership{
		ID:             m.ID,
		UserName:       m.UserName,
		MembershipType: m.MembershipType,
	}, nil
}

func (m *Membership) validateMembership() error {
	if m.UserName == "" {
		return errors.New("there is no user name")
	}
	if m.MembershipType == "" {
		return errors.New("there is no membership type")
	}
	return nil
}

func NewMembershipBuilder() *Membership {
	return &Membership{}
}
