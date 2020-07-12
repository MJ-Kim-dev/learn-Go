package members

import "errors"

// Member struct
type Member struct {
	name    string
	mobile  string
	value   int
	address string
}

var errNoValue = errors.New("Can't Extend MemberShip")

// AddMember create Member
func AddMember(name, mobile string, value int) *Member {
	member := Member{name: name, mobile: mobile, value: value}
	return &member
}

// SetAddress of your information
func (a *Member) SetAddress(address string) {
	a.address = address
}

// ExtMemberShip of your licence
func (a *Member) ExtMemberShip() (int, error) {
	if a.value < 30 {
		return a.value, errNoValue
	}
	a.value -= 30
	return a.value, nil
}
