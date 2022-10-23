package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/moriuriel/go-payments/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return AccountRepository{
		db: db,
	}
}

func (r AccountRepository) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	var query = `
		INSERT INTO 
			accounts (id, name, document, created_at) 
		VALUES 
			($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, account.ID(), account.Name(), account.Document(), account.CreatedAt())
	if err != nil {
		return domain.Account{}, errors.Wrap(err, "error to create account")
	}
	return account, nil
}

func (r AccountRepository) FindByID(ctx context.Context, ID string) (domain.Account, error) {
	var (
		query     = "SELECT * FROM accounts WHERE id = $1 LIMIT 1"
		id        string
		name      string
		document  string
		createdAt time.Time
	)

	err := r.db.QueryRowContext(ctx, query, ID).Scan(&id, &name, &document, &createdAt)

	if err != nil {
		return domain.Account{}, errors.Wrap(err, "erro to find account")
	}

	return domain.NewAccount(id, name, domain.AccountDocument(document), createdAt), nil
}
