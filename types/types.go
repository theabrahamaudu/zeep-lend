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
	ID			uuid.UUID	`json:"id"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Email		string		`json:"email"`
	Password	string		`json:"-"`
	CreatedAt	time.Time	`json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
}


