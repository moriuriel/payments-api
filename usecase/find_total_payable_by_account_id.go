package usecase

import (
	"context"
	"time"

	"github.com/moriuriel/go-payments/domain"
)

type (
	FindTotalPayableByAccountIDInput struct {
		AccountID string
	}

	FindTotalPayableByAccountIDOutput struct {
		TotalPaid         float64 `json:"total_paid"`
		TotalWaitingFunds float64 `json:"total_waiting_funds"`
	}

	FindTotalPayableByAccountIDUsecase interface {
		Execute(context.Context, string) (FindTotalPayableByAccountIDOutput, error)
	}

	FindTotalPayableByAccountIDPresenter interface {
		Output(float64, float64) FindTotalPayableByAccountIDOutput
	}

	FindTotalPayableByAccountIDContainer struct {
		repo       domain.PayableRepository
		ctxTimeout time.Duration
		presenter  FindTotalPayableByAccountIDPresenter
	}
)

func NewFindTotalPayableByAccountIDContainer(r domain.PayableRepository, t time.Duration, p FindTotalPayableByAccountIDPresenter) FindTotalPayableByAccountIDContainer {
	return FindTotalPayableByAccountIDContainer{
		repo:       r,
		ctxTimeout: t,
		presenter:  p,
	}
}

func (uc FindTotalPayableByAccountIDContainer) Execute(ctx context.Context, accountID string) (FindTotalPayableByAccountIDOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	totalPaid, err := uc.repo.SumAmountPaidByStatus(ctx, "paid", accountID)
	if err != nil {
		return uc.presenter.Output(0, 0), err
	}

	totalWaitingFunds, err := uc.repo.SumAmountPaidByStatus(ctx, "waiting_funds", accountID)
	if err != nil {
		return uc.presenter.Output(0, 0), err
	}

	return uc.presenter.Output(totalPaid, totalWaitingFunds), nil
}
