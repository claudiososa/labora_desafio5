package Models

type User struct {
	id    int
	name  string
	email string
}

func (user *User) GetId() int {
	return user.id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) SetId(id int) *User {
	user.id = id
	return user
}

func (user *User) SetName(name string) *User {
	user.name = name
	return user
}

func (user *User) SetEmail(email string) *User {
	user.email = email
	return user
}
