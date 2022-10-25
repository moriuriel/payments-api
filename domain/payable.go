package domain

import (
	"context"
	"time"
)

type (
	PayableRepository interface {
		Create(ctx context.Context, payable Payable) (Payable, error)
		ExecuteWithTransaction(ctx context.Context, fn func(ctxFn context.Context) error) error
	}
	Payable struct {
		id            string
		transactionId string
		accountId     string
		amountPaid    float64
		status        string
		fee           int64
		paymentDate   time.Time
		createdAt     time.Time
	}
)

func NewPayable(
	id string,
	transactionId string,
	accountId string,
	amountPaid float64,
	status string,
	fee int64,
	paymentDate time.Time,
	createdAt time.Time) Payable {
	return Payable{
		id:            id,
		transactionId: transactionId,
		accountId:     accountId,
		amountPaid:    amountPaid,
		status:        status,
		fee:           fee,
		paymentDate:   paymentDate,
		createdAt:     createdAt,
	}
}

func (p Payable) ID() string {
	return p.id
}

func (p Payable) AccountID() string {
	return p.accountId
}

func (p Payable) TransactionID() string {
	return p.transactionId
}

func (p Payable) AmountPaid() float64 {
	return p.amountPaid
}

func (p Payable) Status() string {
	return p.status
}

func (p Payable) Fee() int64 {
	return p.fee
}

func (p Payable) PaymentDate() time.Time {
	return p.paymentDate
}

func (p Payable) CreatedAt() time.Time {
	return p.createdAt
}
