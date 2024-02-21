package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	coonPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new SQLStore
func NewStore(coonPool *pgxpool.Pool) Store {
	return &SQLStore{
		coonPool: coonPool,
		Queries:  New(coonPool),
	}
}
