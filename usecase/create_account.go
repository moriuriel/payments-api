package usecase

import (
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/moriuriel/go-payments/domain"
)

type (
	CreateAccountUsecase interface {
		Execute(context.Context, CreateAccountInput) (CreateAccountOutput, error)
	}

	CreateAccountPresenter interface {
		Output(domain.Account) CreateAccountOutput
	}

	CreateAccountInput struct {
		Name     string `json:"name"`
		Document string `json:"document"`
	}

	CreateAccountOutput struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Document  string `json:"document"`
		CreatedAt string `json:"created_at"`
	}

	CreateAccountContainer struct {
		presenter  CreateAccountPresenter
		repo       domain.AccountRepository
		ctxTimeout time.Duration
	}
)

func NewCreateAccountContainer(r domain.AccountRepository, t time.Duration, p CreateAccountPresenter) CreateAccountContainer {
	return CreateAccountContainer{
		presenter:  p,
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc CreateAccountContainer) Execute(ctx context.Context, input CreateAccountInput) (CreateAccountOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	document := domain.AccountDocument(input.Document).LastFourDocumentDigit()

	var account = domain.NewAccount(
		uuid.New().String(),
		input.Name,
		domain.AccountDocument(document),
		time.Now(),
	)

	account, err := uc.repo.Create(ctx, account)

	if err != nil {
		return uc.presenter.Output(domain.Account{}), err
	}

	return uc.presenter.Output(account), nil
}
