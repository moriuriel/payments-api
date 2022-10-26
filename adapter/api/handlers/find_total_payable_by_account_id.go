package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moriuriel/go-payments/adapter/api/response"
	"github.com/moriuriel/go-payments/usecase"
)

type FindTotalPayableByAccountIDHandler struct {
	uc usecase.FindTotalPayableByAccountIDUsecase
}

func NewFindTotalPayableByAccountIDHandler(uc usecase.FindTotalPayableByAccountIDUsecase) FindTotalPayableByAccountIDHandler {
	return FindTotalPayableByAccountIDHandler{
		uc: uc,
	}
}

func (h FindTotalPayableByAccountIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["account_id"]

	output, err := h.uc.Execute(r.Context(), accountId)
	if err != nil {
		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)

}
