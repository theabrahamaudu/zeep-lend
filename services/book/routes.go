package book

import (
	"net/http"

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
	router.HandleFunc("/books", h.handleGetBook).Methods(http.MethodGet)
}

func (h *Handler) handleGetBook(w http.ResponseWriter, r *http.Request) {
	bs, err := h.store.GetBooks()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, bs)
}