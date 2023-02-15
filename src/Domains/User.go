package Domains

type User struct {
	id        UserId
	firstName string
	lastName  string
}

func NewUser(id UserId, firstName string, lastName string) *User {
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
	}
}

func (user *User) GetId() UserId {
	return user.id
}

func (user *User) GetFirstName() string {
	return user.firstName
}

func (user *User) GetLastName() string {
	return user.lastName
}
