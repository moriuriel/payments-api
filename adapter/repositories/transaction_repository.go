package repositories

import (
	"context"
	"database/sql"
	"github.com/moriuriel/go-payments/domain"
	"github.com/pkg/errors"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return TransactionRepository{
		db: db,
	}
}

func (t TransactionRepository) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	var query = `
		INSERT INTO 
			transactions (id, account_id, description, amount, card_owner, card_number, card_expiration_date, card_cvv, payment_method, created_at) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	tx, ok := ctx.Value("TxKey").(*sql.Tx)
	if !ok {
		var err error
		tx, err = t.db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return domain.Transaction{}, errors.Wrap(err, "unknown error")
		}
	}

	_, err := tx.ExecContext(
		ctx,
		query,
		transaction.ID(),
		transaction.AccountID(),
		transaction.Description(),
		transaction.Amount(),
		transaction.CardOwner(),
		transaction.CardNumber(),
		transaction.CardExpirationDate(),
		transaction.CardCvv(),
		transaction.PaymentMethod(),
		transaction.CreatedAt(),
	)
	if err != nil {
		return domain.Transaction{}, errors.Wrap(err, "error to create transaction")
	}

	return transaction, nil
}

func (t TransactionRepository) ExecuteWithTransaction(ctx context.Context, fn func(ctxFn context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "error to execute transaction")
	}

	ctxTx := context.WithValue(ctx, "TxKey", tx)
	err = fn(ctxTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Wrap(err, "rollback error")
		}
		return err
	}

	return tx.Commit()
}
