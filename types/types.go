package types

import (
	"time"

	"github.com/google/uuid"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID			int	`json:"id"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Email		string		`json:"email"`
	Password	string		`json:"-"`
	CreatedAt	time.Time	`json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName	string	`json:"firstname" validate:"required"`
	LastName	string	`json:"lastname" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Password 	string	`json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email		string	`json:"email" validate:"required,email"`
	Password 	string	`json:"password" validate:"required"`
}


