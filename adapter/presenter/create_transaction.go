package presenter

import (
	"github.com/moriuriel/go-payments/domain"
	"github.com/moriuriel/go-payments/usecase"
)

type createTransactionPresenter struct{}

func NewCreateTransactionPresenter() usecase.CreateTransactionPresenter {
	return createTransactionPresenter{}
}

func (c createTransactionPresenter) Output(transaction domain.Transaction, payableID string) usecase.CreateTransactionOutput {
	return usecase.CreateTransactionOutput{
		ID:                 transaction.ID(),
		AccountID:          transaction.AccountID(),
		Amount:             transaction.Amount(),
		Description:        transaction.Description(),
		CardOwner:          transaction.CardOwner(),
		CardNumber:         transaction.CardNumber(),
		CardCvv:            transaction.CardCvv(),
		CardExpirationDate: transaction.CardExpirationDate(),
		PaymentMethod:      transaction.PaymentMethod(),
		PayableID:          payableID,
		CreatedAt:          transaction.CreatedAt(),
	}
}
