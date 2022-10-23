package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/moriuriel/go-payments/domain"
	"time"
)

type (
	CreateTransactionInput struct {
		Amount             int64  `json:"amount"`
		AccountID          string `json:"account_id"`
		Description        string `json:"description"`
		CardOwner          string `json:"card_owner"`
		CardNumber         string `json:"card_number"`
		CardExpirationDate string `json:"card_expiration_date"`
		CardCvv            int64  `json:"card_cvv"`
		PaymentMethod      string `json:"payment_method"`
	}

	CreateTransactionOutput struct {
		ID                 string    `json:"id"`
		Amount             int64     `json:"amount"`
		AccountID          string    `json:"account_id"`
		Description        string    `json:"description"`
		CardOwner          string    `json:"card_owner"`
		CardNumber         string    `json:"card_number"`
		CardExpirationDate string    `json:"card_expiration_date"`
		CardCvv            int64     `json:"card_cvv"`
		PaymentMethod      string    `json:"payment_method"`
		CreatedAt          time.Time `json:"created_at"`
	}

	CreateTransactionPresenter interface {
		Output(transaction domain.Transaction) CreateTransactionOutput
	}

	CreateTransactionUsecase interface {
		Execute(ctx context.Context, input CreateTransactionInput) (CreateTransactionInput, error)
	}

	CreateTransactionContainer struct {
		presenter       CreateTransactionPresenter
		accountRepo     domain.AccountRepository
		ctxTimeout      time.Duration
		transactionRepo domain.TransactionRepository
	}
)

func NewCreateTransactionContainer(p CreateTransactionPresenter, aRepo domain.AccountRepository, t time.Duration, tRepo domain.TransactionRepository) CreateTransactionContainer {
	return CreateTransactionContainer{
		presenter:       p,
		accountRepo:     aRepo,
		transactionRepo: tRepo,
		ctxTimeout:      t,
	}
}

func (uc CreateTransactionContainer) Execute(ctx context.Context, input CreateTransactionInput) (CreateTransactionOutput, error) {
	var (
		transaction domain.Transaction
		err         error
	)

	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	err = uc.transactionRepo.ExecuteWithTransaction(ctx, func(ctxFn context.Context) error {
		_, err := uc.accountRepo.FindByID(ctxFn, input.AccountID)
		if err != nil {
			return err
		}

		transaction = domain.NewTransaction(
			uuid.New().String(),
			input.AccountID,
			input.Description,
			input.Amount,
			input.CardOwner,
			input.CardNumber,
			input.CardExpirationDate,
			input.CardCvv,
			input.PaymentMethod,
			time.Now(),
		)

		_, err = uc.transactionRepo.Create(ctxFn, transaction)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return uc.presenter.Output(domain.Transaction{}), err
	}

	return uc.presenter.Output(transaction), nil
}
