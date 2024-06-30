package transaction

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/theabrahamaudu/zeep-lend/types"
	"github.com/theabrahamaudu/zeep-lend/utils"
)

type Handler struct {
	store types.TransactionStore
	bookStore types.BookStore
}

func NewHandler(store types.TransactionStore,
				bookStore types.BookStore) *Handler {
	return &Handler{store: store, bookStore: bookStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc(
		"/transaction/lend",
		h.handleLend).Methods(http.MethodPost)
	router.HandleFunc(
		"/transaction/return",
		h.handleReturn).Methods(http.MethodPost)
}

func (h *Handler) handleLend(w http.ResponseWriter, r *http.Request) {
	var transaction types.CreateTransactionPayload

	if err:= utils.ParseJSON(r, &transaction); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	if err := utils.Validate.Struct(transaction); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"invalid payload: %v", errors, 
		))
		return
	}

	if err := h.getBookStatus(transaction); err != nil {
		utils.WriteError(w, http.StatusConflict, err)
		return
	}

	transaction_id, err := h.store.CreateTransaction(transaction)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, transaction_id)

}

func (h *Handler) handleReturn(w http.ResponseWriter, r *http.Request) {
	var transaction types.CloseTransactionPayload

	if err:= utils.ParseJSON(r, &transaction); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	if err := utils.Validate.Struct(transaction); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"invalid payload: %v", errors, 
		))
		return
	}

	rows_affected, err := h.store.CloseTransaction(transaction)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, rows_affected)
}