package membership

import "errors"

type Membership struct {
	ID             string
	UserName       string
	MembershipType string
}

func (m *MembershipBuilder) SetID(id string) *MembershipBuilder {
	m.Membership.ID = id
	return m
}

func (m *MembershipBuilder) SetUserName(userName string) *MembershipBuilder {
	m.Membership.UserName = userName
	return m
}

func (m *MembershipBuilder) SetMembershipType(membershipType string) *MembershipBuilder {
	m.Membership.MembershipType = membershipType
	return m
}

func (m *MembershipBuilder) GetMembership() (*Membership, error) {
	if err := m.validateMembership(); err != nil {
		return &Membership{}, err
	}
	return m.Membership, nil
}

func (m *MembershipBuilder) validateMembership() error {
	if m.Membership.ID == "" {
		return errors.New("there is no id")
	}
	if m.Membership.UserName == "" {
		return errors.New("there is no user name")
	}
	if m.Membership.MembershipType == "" {
		return errors.New("there is no membership type")
	}
	if !(m.Membership.MembershipType == "naver" || m.Membership.MembershipType == "toss" || m.Membership.MembershipType == "payco") {
		return errors.New("not supported membership")
	}
	return nil
}

type MembershipBuilder struct {
	Membership *Membership
}

func NewMembershipBuilder() *MembershipBuilder {
	return &MembershipBuilder{Membership: &Membership{}}
}
