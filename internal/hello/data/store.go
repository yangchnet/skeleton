package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewStore)

type Store struct {
	db *sql.DB
	*Queries
}

type HelloInterface interface {
	Querier
}

func NewStore(db *sql.DB) HelloInterface {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
