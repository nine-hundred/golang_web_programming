package membership

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

func (m *Membership) GetMembership() Membership {
	return Membership{
		ID:             m.ID,
		UserName:       m.UserName,
		MembershipType: m.MembershipType,
	}
}

func NewMembershipBuilder() *Membership {
	return &Membership{}
}
