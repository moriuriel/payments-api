package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/moriuriel/go-payments/adapter/api/response"
	"github.com/moriuriel/go-payments/usecase"
)

type FindAccountByIDHandler struct {
	uc usecase.FindAccountByIDUsecase
}

func NewFindAccountByIDHandler(uc usecase.FindAccountByIDUsecase) FindAccountByIDHandler {
	return FindAccountByIDHandler{
		uc: uc,
	}
}

func (h FindAccountByIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := uuid.Parse(id)
	if err != nil {
		error := errors.New("invalid account_id")
		response.NewError(error, http.StatusBadRequest).Send(w)
		return
	}

	output, err := h.uc.Execute(r.Context(), id)
	if err != nil {
		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
