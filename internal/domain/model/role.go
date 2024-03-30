package model

type Role uint

const (
	Unknown Role = iota + 1
	Admin
	Member
)

func NewRole(i uint) Role {
	switch i {
	case Admin.Uint(), Member.Uint():
		return Role(i)
	default:
		return Unknown
	}
}

func (r Role) Uint() uint {
	return uint(r)
}
