package usecase

import (
	"context"
	"github.com/moriuriel/go-payments/domain"
	"time"
)

type (
	FindAccountByIDUsecase interface {
		Execute(context.Context, string) (FindAccountByIDOutput, error)
	}

	FindAccountByIDPresenter interface {
		Output(domain.Account) FindAccountByIDOutput
	}

	FindAccountByIDOutput struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Document  string `json:"document"`
		CreatedAt string `json:"created_at"`
	}

	FindAccountByIDContainer struct {
		presenter  FindAccountByIDPresenter
		repo       domain.AccountRepository
		ctxTimeout time.Duration
	}
)

func NewFindAccountByIDContainer(r domain.AccountRepository, t time.Duration, p FindAccountByIDPresenter) FindAccountByIDContainer {
	return FindAccountByIDContainer{
		presenter:  p,
		repo:       r,
		ctxTimeout: t,
	}

}

func (uc FindAccountByIDContainer) Execute(ctx context.Context, ID string) (FindAccountByIDOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	account, err := uc.repo.FindByID(ctx, ID)
	if err != nil {
		return uc.presenter.Output(domain.Account{}), err
	}

	return uc.presenter.Output(account), nil
}
