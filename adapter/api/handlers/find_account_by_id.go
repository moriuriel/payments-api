package handlers

import (
	"net/http"

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

	output, err := h.uc.Execute(r.Context(), id)
	if err != nil {
		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
