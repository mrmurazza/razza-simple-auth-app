package user

import (
	"dealljobs/domain/request"
)

type Repository interface {
	Persist(u *User) (*User, error)
	Update(u *User) error
	Delete(id int) error

	GetUser(id int) *User
	GetUserByUserPass(username, password string) *User
	GetUserByUsername(u User) (*User, error)
	GetAllUsers() []*User
}

type Service interface {
	CreateUserIfNotAny(req request.CreateUserRequest) (*User, error)
	UpdateUser(req request.UpdateUserRequest) error
	DeleteUser(id int) error

	Login(username, password string) (*User, string, error)
	GetAllUsers() []*User
	GetUser(id int) *User
}
