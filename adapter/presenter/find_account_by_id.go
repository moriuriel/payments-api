package presenter

import (
	"github.com/moriuriel/go-payments/domain"
	"github.com/moriuriel/go-payments/usecase"
	"time"
)

type findAccountByIDPresenter struct {
}

func NewFindAccountByIDPresenter() usecase.FindAccountByIDPresenter {
	return findAccountByIDPresenter{}
}

func (c findAccountByIDPresenter) Output(account domain.Account) usecase.FindAccountByIDOutput {
	return usecase.FindAccountByIDOutput{
		ID:        account.ID(),
		Name:      account.Name(),
		Document:  account.Document().String(),
		CreatedAt: account.CreatedAt().Format(time.RFC3339),
	}
}
