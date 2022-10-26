package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/moriuriel/go-payments/domain"
	"github.com/pkg/errors"
)

type PayableRepository struct {
	db *sql.DB
}

func NewPayableRepository(db *sql.DB) PayableRepository {
	return PayableRepository{
		db: db,
	}
}

func (r PayableRepository) Create(ctx context.Context, payable domain.Payable) (domain.Payable, error) {
	var query = `
		INSERT INTO 
    		payables (id, account_id, transaction_id, amount_paid, status, fee, payment_date, created_at)
		VALUES 
		    ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	tx, ok := ctx.Value("TxKey").(*sql.Tx)
	if !ok {
		var err error
		tx, err = r.db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return domain.Payable{}, errors.Wrap(err, "unknown error")
		}
	}

	_, err := tx.ExecContext(
		ctx,
		query,
		payable.ID(),
		payable.AccountID(),
		payable.TransactionID(),
		payable.AmountPaid(),
		payable.Status(),
		payable.Fee(),
		payable.PaymentDate(),
		payable.CreatedAt(),
	)

	if err != nil {
		return domain.Payable{}, errors.Wrap(err, "error to create payable")
	}
	return payable, nil
}

func (t PayableRepository) ExecuteWithTransaction(ctx context.Context, fn func(ctxFn context.Context) error) error {
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

func (t PayableRepository) SumAmountPaidByStatus(ctx context.Context, status string, accountID string) (float64, error) {

	var (
		query = "SELECT SUM(amount_paid) AS total FROM payables WHERE status = $1 and account_id = $2;"
		total float64
	)

	err := t.db.QueryRowContext(ctx, query, status, accountID).Scan(&total)
	if err != nil {
		return 0, errors.Wrap(err, "error to sum amount_paid")
	}
	return total, nil
}

func (t PayableRepository) FindAllByAccountID(ctx context.Context, accountID string) ([]domain.Payable, error) {
	var query = "SELECT * FROM payables WHERE account_id = $1"

	rows, err := t.db.QueryContext(ctx, query, accountID)
	if err != nil {
		return []domain.Payable{}, errors.Wrap(err, "error to listing payable")
	}

	payables := make([]domain.Payable, 0)
	for rows.Next() {
		var (
			ID            string
			accountID     string
			transactionID string
			amountPaid    float64
			status        string
			fee           int64
			paymentDate   time.Time
			createdAt     time.Time
		)

		err = rows.Scan(&ID, &accountID, &transactionID, &status, &fee, &paymentDate, &amountPaid, &createdAt)
		if err != nil {
			return []domain.Payable{}, errors.Wrap(err, "error to listing payable")
		}

		payables = append(payables, domain.NewPayable(
			ID,
			accountID,
			transactionID,
			amountPaid,
			status,
			fee,
			paymentDate,
			createdAt,
		))
	}
	defer rows.Close()
	if err = rows.Err(); err != nil {
		return []domain.Payable{}, err
	}

	return payables, nil

}
