package presenter

import (
	"time"

	"github.com/moriuriel/go-payments/domain"
	"github.com/moriuriel/go-payments/usecase"
)

type createAccountPresenter struct{}

func NewCreateAccountPresenter() usecase.CreateAccountPresenter {
	return createAccountPresenter{}
}

func (c createAccountPresenter) Output(account domain.Account) usecase.CreateAccountOutput {
	return usecase.CreateAccountOutput{
		ID:        account.ID(),
		Name:      account.Name(),
		Document:  account.Document().String(),
		CreatedAt: account.CreatedAt().Format(time.RFC3339),
	}
}
