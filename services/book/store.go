package book

import (
	"database/sql"

	"github.com/theabrahamaudu/zeep-lend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateBook(book types.Book) error {
	_, err := s.db.Exec(
		`INSERT INTO books (
			userId, name, description, image, fee, duration, status
		) 
		VALUES (?,?,?,?,?,?,?)`,
		book.UserID,
		book.Name,
		book.Description,
		book.Image,
		book.Fee,
		book.Duration,
		book.Status,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetBooks() ([]types.Book, error) {
	rows, err := s.db.Query("SELECT * FROM books")

	if err != nil {
		return nil, err
	}

	books := make([]types.Book, 0)

	for rows.Next() {
		b, err := scanRowIntoBook(rows)

		if err != nil {
			return nil, err
		}

		books = append(books, *b)
	}

	return books, nil
}

func (s *Store) GetBookByName(name string) (*types.Book, error) {
	rows, err := s.db.Query("SELECT * FROM books WHERE name = ?", name)

	if err != nil {
		return nil, err
	}

	b := new(types.Book)
	for rows.Next() {
		b, err = scanRowIntoBook(rows)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func (s *Store) GetBookByID(id int) (*types.Book, error) {
	rows, err := s.db.Query("SELECT * FROM books WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	b := new(types.Book)
	for rows.Next() {
		b, err = scanRowIntoBook(rows)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func scanRowIntoBook(rows *sql.Rows) (*types.Book, error) {
	book := new(types.Book)

	err := rows.Scan(
		&book.ID,
		&book.UserID,
		&book.Name,
		&book.Description,
		&book.Image,
		&book.Fee,
		&book.Duration,
		&book.Status,
		&book.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return book, nil
}