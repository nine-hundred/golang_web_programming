package membership

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
	return m.Membership, nil
}

type MembershipBuilder struct {
	Membership *Membership
}

func NewMembershipBuilder() *MembershipBuilder {
	return &MembershipBuilder{Membership: &Membership{}}
}
