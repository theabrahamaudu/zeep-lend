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

type BookStore interface {
	GetBooks() ([]Book, error)
	GetBookByName(name string) (*Book, error)
	GetBookByID(id int) (*Book, error)
	CreateBook(Book) error
}

type TransactionStore interface {
	CreateTransaction(transaction CreateTransactionPayload) (int, error)
	CloseTransaction(transaction CloseTransactionPayload) (int, error)
}

type Transaction struct {
	ID			int		`json:"id"`
	LenderID	int		`json:"lenderId"`
	BorrowerID	int		`json:"borrowerId"`
	BookID		int		`json:"bookId"`
	Status		int		`json:"status"`
	CreatedAt	int		`json:"createdAt"`
}

type CreateTransactionPayload struct {
	LenderID	int		`json:"lenderId" validate:"required"`
	BorrowerID	int		`json:"borrowerId" validate:"required"`
	BookID		int		`json:"bookId" validate:"required"`
	Status		int		`json:"status" validate:"required"`
}

type CloseTransactionPayload struct {
	ID			int		`json:"id" validate:"required"`
	LenderID	int		`json:"lenderId" validate:"required"`
	BorrowerID	int		`json:"borrowerId" validate:"required"`
}

type Book struct {
	ID			int			`json:"id"`
	UserID		int			`json:"userId"`
	Name 		string		`json:"name"`
	Description	string		`json:"description"`
	Image		string		`json:"image"`
	Fee			int			`json:"fee"`
	Duration 	int			`json:"duration"`
	Status  	int			`json:"status"`
	CreatedAt	time.Time 	`json:"createdAt"`
}

type CreateBookPayload struct {
	UserID		int			`json:"userId" validate:"required"`
	Name 		string		`json:"name" validate:"required"`
	Description	string		`json:"description" validate:"required"`
	Image		string		`json:"image" validate:"required"`
	Fee			int			`json:"fee" validate:"required"`
	Duration 	int			`json:"duration" validate:"required"`
	Status  	int			`json:"status" validate:"required"`
}

type User struct {
	ID			int			`json:"id"`
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


