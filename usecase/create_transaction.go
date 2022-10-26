package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/moriuriel/go-payments/domain"
)

type (
	CreateTransactionInput struct {
		Amount             float64 `json:"amount"`
		AccountID          string  `json:"account_id"`
		Description        string  `json:"description"`
		CardOwner          string  `json:"card_owner"`
		CardNumber         string  `json:"card_number"`
		CardExpirationDate string  `json:"card_expiration_date"`
		CardCvv            int64   `json:"card_cvv"`
		PaymentMethod      string  `json:"payment_method"`
	}

	CreateTransactionOutput struct {
		ID                 string    `json:"id"`
		Amount             float64   `json:"amount"`
		AccountID          string    `json:"account_id"`
		Description        string    `json:"description"`
		CardOwner          string    `json:"card_owner"`
		CardNumber         string    `json:"card_number"`
		CardExpirationDate string    `json:"card_expiration_date"`
		CardCvv            int64     `json:"card_cvv"`
		PaymentMethod      string    `json:"payment_method"`
		PayableID          string    `json:"payable_id"`
		CreatedAt          time.Time `json:"created_at"`
	}

	CreateTransactionPresenter interface {
		Output(transaction domain.Transaction, PayableID string) CreateTransactionOutput
	}

	CreateTransactionUsecase interface {
		Execute(ctx context.Context, input CreateTransactionInput) (CreateTransactionInput, error)
	}

	CreateTransactionContainer struct {
		presenter       CreateTransactionPresenter
		accountRepo     domain.AccountRepository
		ctxTimeout      time.Duration
		transactionRepo domain.TransactionRepository
		payableRepo     domain.PayableRepository
	}
)

func NewCreateTransactionContainer(p CreateTransactionPresenter, aRepo domain.AccountRepository, t time.Duration, tRepo domain.TransactionRepository, pRepo domain.PayableRepository) CreateTransactionContainer {
	return CreateTransactionContainer{
		presenter:       p,
		accountRepo:     aRepo,
		transactionRepo: tRepo,
		ctxTimeout:      t,
		payableRepo:     pRepo,
	}
}

func (uc CreateTransactionContainer) Execute(ctx context.Context, input CreateTransactionInput) (CreateTransactionOutput, error) {
	var (
		transaction   domain.Transaction
		transactionID = uuid.New().String()
		payable       domain.Payable
		err           error
	)

	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	err = uc.payableRepo.ExecuteWithTransaction(ctx, func(ctxFn context.Context) error {
		_, err := uc.accountRepo.FindByID(ctxFn, input.AccountID)
		if err != nil {
			return err
		}

		transaction = domain.NewTransaction(
			transactionID,
			input.AccountID,
			input.Description,
			input.Amount,
			input.CardOwner,
			domain.CardNumber(input.CardNumber).LastFourDigit(),
			input.CardExpirationDate,
			input.CardCvv,
			input.PaymentMethod,
			time.Now(),
		)

		_, err = uc.transactionRepo.Create(ctxFn, transaction)
		if err != nil {
			return err
		}

		paymentMethod := domain.NewMethod(domain.PaymentMethod(input.PaymentMethod))
		amountPaid, fee := paymentMethod.CalculateAmountPaid(transaction.Amount())

		payable = domain.NewPayable(
			uuid.New().String(),
			transactionID,
			input.AccountID,
			amountPaid,
			paymentMethod.PaymentStatus(),
			fee,
			paymentMethod.PaymentDate(),
			time.Now(),
		)

		_, err = uc.payableRepo.Create(ctxFn, payable)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return uc.presenter.Output(domain.Transaction{}, ""), err
	}

	return uc.presenter.Output(transaction, payable.ID()), nil
}
