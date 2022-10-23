package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/moriuriel/go-payments/adapter/api/response"
	"github.com/moriuriel/go-payments/usecase"
)

type CreateAccountHandler struct {
	uc usecase.CreateAccountUsecase
}

func NewCreateAccountHandler(uc usecase.CreateAccountUsecase) CreateAccountHandler {
	return CreateAccountHandler{
		uc: uc,
	}
}

func (h CreateAccountHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateAccountInput

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
