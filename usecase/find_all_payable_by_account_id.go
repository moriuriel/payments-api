package usecase

import (
	"context"
	"time"

	"github.com/moriuriel/go-payments/domain"
)

type (
	FindAllPayableByAccountIDOutput struct {
		ID            string  `json:"id"`
		AccountID     string  `json:"account_id"`
		TransactionID string  `json:"transaction_id"`
		AmountPaid    float64 `json:"amount_paid"`
		Fee           int64   `json:"fee"`
		Status        string  `json:"status"`
		PaymentDate   string  `json:"payment_date"`
		CreatedAt     string  `json:"created_at"`
	}

	FindAllPayableByAccountIDUsecase interface {
		Execute(context.Context, string) ([]FindAllPayableByAccountIDOutput, error)
	}

	FindAllPayableByAccountIDPresenter interface {
		Output([]domain.Payable) []FindAllPayableByAccountIDOutput
	}

	FindAllPayableByAccountIDContainer struct {
		repo       domain.PayableRepository
		ctxTimeout time.Duration
		presenter  FindAllPayableByAccountIDPresenter
	}
)

func NewFindAllPayableByAccountIDContainer(r domain.PayableRepository, t time.Duration, p FindAllPayableByAccountIDPresenter) FindAllPayableByAccountIDContainer {
	return FindAllPayableByAccountIDContainer{
		repo:       r,
		ctxTimeout: t,
		presenter:  p,
	}
}

func (uc FindAllPayableByAccountIDContainer) Execute(ctx context.Context, accountID string) ([]FindAllPayableByAccountIDOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	payables, err := uc.repo.FindAllByAccountID(ctx, accountID)
	if err != nil {
		return uc.presenter.Output([]domain.Payable{}), err
	}

	return uc.presenter.Output(payables), nil
}
