package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/moriuriel/go-payments/adapter/api/response"
	"github.com/moriuriel/go-payments/usecase"
)

type FindTAllPayableByAccountIDHandler struct {
	uc usecase.FindAllPayableByAccountIDUsecase
}

func NewFindAllPayableByAccountIDHandler(uc usecase.FindAllPayableByAccountIDUsecase) FindTAllPayableByAccountIDHandler {
	return FindTAllPayableByAccountIDHandler{
		uc: uc,
	}
}

func (h FindTAllPayableByAccountIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["account_id"]

	_, err := uuid.Parse(accountId)
	if err != nil {
		error := errors.New("invalid account_id")
		response.NewError(error, http.StatusBadRequest).Send(w)
		return
	}

	output, err := h.uc.Execute(r.Context(), accountId)
	if err != nil {
		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)

}
