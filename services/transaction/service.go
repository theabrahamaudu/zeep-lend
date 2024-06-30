package transaction

import (
	"fmt"

	"github.com/theabrahamaudu/zeep-lend/types"
)

func (h *Handler) getBookStatus(
		transaction types.CreateTransactionPayload,
	) (error) {

	book, err := h.bookStore.GetBookByID(transaction.BookID)
	if err != nil {
		return fmt.Errorf(
			"book with id %v not found", transaction.BookID,
		)
	}
	
	if book.Status==0 {
		return fmt.Errorf(
			"book with id %d unavailable", book.ID,
		)
	}

	return nil
}