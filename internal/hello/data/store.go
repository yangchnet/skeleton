package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/wire"
)

// ProviderSet provide Store.
var ProviderSet = wire.NewSet(NewStore)

// Store hold db connection and Querier.
type Store struct {
	db *sql.DB
	*Queries
}

// HelloInterface is the interface for operating the database for hello server.
type HelloInterface interface {
	Querier
}

// NewStore return new HelloInterface by given db connection.
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
			return fmt.Errorf("tx err: %v, rb err: %w", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}
