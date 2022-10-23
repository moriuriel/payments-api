package domain

import (
	"context"
	"time"
)

type (
	AccountRepository interface {
		Create(context.Context, Account) (Account, error)
		FindByID(ctx context.Context, string2 string) (Account, error)
	}

	Account struct {
		id        string
		name      string
		document  AccountDocument
		createdAt time.Time
	}
)

type AccountDocument string

func (Document AccountDocument) LastFourDocumentDigit() string {
	return string(Document[len(Document)-4:])
}

func (Document AccountDocument) String() string {
	return string(Document)
}

func NewAccount(ID string, name string, document AccountDocument, createdAt time.Time) Account {
	return Account{
		id:        ID,
		name:      name,
		document:  document,
		createdAt: createdAt,
	}
}

func (a Account) ID() string {
	return a.id
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Document() AccountDocument {
	return a.document
}

func (a Account) CreatedAt() time.Time {
	return a.createdAt
}
