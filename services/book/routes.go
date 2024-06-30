package book

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/theabrahamaudu/zeep-lend/types"
	"github.com/theabrahamaudu/zeep-lend/utils"
)

type Handler struct {
	store types.BookStore
}

func NewHandler(store types.BookStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/books/all", h.handleGetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/createbook", h.handleCreateBook).Methods(http.MethodPost)
}

func (h *Handler) handleCreateBook(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateBookPayload
	if err := utils.ParseJSON(r, &payload); err!= nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// valiate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invalid payload %v", errors),
		)
		return
	}

	// check if book exists
	_, err := h.store.GetBookByName(payload.Name)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"book with name %s already exists", payload.Name,
		))
		return		
	}

	// create book if book doesn't exist
	err = h.store.CreateBook(types.Book{
		UserID: payload.UserID,
		Name: payload.Name,
		Description: payload.Description,
		Image: payload.Image,
		Fee: payload.Fee,
		Duration: payload.Duration,
		Status: payload.Status,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetBook(w http.ResponseWriter, r *http.Request) {
	bs, err := h.store.GetBooks()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, bs)
}