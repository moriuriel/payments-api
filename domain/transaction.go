package domain

import (
	"context"
	"time"
)

type (
	TransactionRepository interface {
		Create(ctx context.Context, transaction Transaction) (Transaction, error)
		ExecuteWithTransaction(ctx context.Context, fn func(ctxFn context.Context) error) error
	}
	Transaction struct {
		id                 string
		accountID          string
		description        string
		amount             float64
		cardOwner          string
		cardNumber         string
		cardExpirationDate string
		cardCvv            int64
		paymentMethod      string
		createdAt          time.Time
	}
)
type CardNumber string

func (Card CardNumber) LastFourDigit() string {
	return string(Card[len(Card)-4:])
}

func NewTransaction(ID string, accountID string, description string, amount float64, cardOwner string, cardNumber string, cardExpirationDate string, cardCvv int64, paymentMethod string, createdAt time.Time) Transaction {
	return Transaction{
		id:                 ID,
		accountID:          accountID,
		description:        description,
		amount:             amount,
		cardOwner:          cardOwner,
		cardNumber:         cardNumber,
		cardExpirationDate: cardExpirationDate,
		cardCvv:            cardCvv,
		paymentMethod:      paymentMethod,
		createdAt:          createdAt,
	}
}

func (t Transaction) CalculateAmountPaid() (float64, int64) {
	switch t.paymentMethod {
	case "credit_card":
		return t.amount - (t.amount / 100 * 5), 5
	case "debit_card":
		return t.amount - (t.amount / 100 * 3), 3
	default:
		return t.amount, 0
	}
}

func (t Transaction) ID() string {
	return t.id
}

func (t Transaction) AccountID() string {
	return t.accountID
}

func (t Transaction) Description() string {
	return t.description
}

func (t Transaction) Amount() float64 {
	return t.amount
}

func (t Transaction) CardOwner() string {
	return t.cardOwner
}

func (t Transaction) CardNumber() string {
	return t.cardNumber
}

func (t Transaction) CardExpirationDate() string {
	return t.cardExpirationDate
}

func (t Transaction) CardCvv() int64 {
	return t.cardCvv
}

func (t Transaction) CreatedAt() time.Time {
	return t.createdAt
}

func (t Transaction) PaymentMethod() string {
	return t.paymentMethod
}
