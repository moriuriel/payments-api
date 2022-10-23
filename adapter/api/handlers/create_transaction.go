package handlers

import (
	"encoding/json"
	"github.com/moriuriel/go-payments/adapter/api/response"
	"github.com/moriuriel/go-payments/usecase"
	"net/http"
)

type CreateTransactionHandler struct {
	uc usecase.CreateTransactionContainer
}

func NewCreateTransactionHandler(uc usecase.CreateTransactionContainer) CreateTransactionHandler {
	return CreateTransactionHandler{
		uc: uc,
	}
}

func (h CreateTransactionHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateTransactionInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	output, err := h.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
