package presenter

import (
	"time"

	"github.com/moriuriel/go-payments/domain"
	"github.com/moriuriel/go-payments/usecase"
)

type findAllPayableByAccountIDPresenter struct{}

func NewFindAllPayableByAccountIDPresenter() usecase.FindAllPayableByAccountIDPresenter {
	return findAllPayableByAccountIDPresenter{}
}

func (p findAllPayableByAccountIDPresenter) Output(payables []domain.Payable) []usecase.FindAllPayableByAccountIDOutput {
	var output = make([]usecase.FindAllPayableByAccountIDOutput, 0)

	for _, payable := range payables {
		output = append(output, usecase.FindAllPayableByAccountIDOutput{
			ID:            payable.ID(),
			AccountID:     payable.AccountID(),
			TransactionID: payable.TransactionID(),
			AmountPaid:    payable.AmountPaid(),
			Fee:           payable.Fee(),
			Status:        payable.Status(),
			PaymentDate:   payable.PaymentDate().Format(time.RFC3339),
			CreatedAt:     payable.CreatedAt().Format(time.RFC3339),
		})
	}
	return output
}
