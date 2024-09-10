package user

type UserRepository interface {
	CreateUser(user *User) error
	GetUserById(id int) (*User, error)
	GetAllUsers() ([]*User, error)
	DeleteUserById(id int) error
}
