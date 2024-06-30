package transaction

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

func (s *Store) CreateTransaction(transaction types.CreateTransactionPayload) (int, error) {
	res, err := s.db.Exec(
		`INSERT INTO transactions (
			lenderId, borrowerId, bookId, status
		) VALUES (?,?,?,?)`, 
	transaction.LenderID,
	transaction.BorrowerID,
	transaction.BookID,
	transaction.Status,)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CloseTransaction(transaction types.CloseTransactionPayload) (int, error) {
	res, err := s.db.Exec(
		`UPDATE transactions SET status = 0 WHERE id = ?
		`, transaction.ID,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}