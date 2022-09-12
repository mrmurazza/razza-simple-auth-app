package user

import (
	"time"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

func GetAllRoles() map[Role]bool {
	return map[Role]bool{
		RoleAdmin: true,
		RoleUser:  true,
	}
}

type User struct {
	Id        int
	Username  string
	Name      string
	Password  string
	Address   string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
