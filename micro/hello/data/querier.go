// Code generated by sqlc. DO NOT EDIT.

package data

import (
	"context"
)

type Querier interface {
	CreateMessage(ctx context.Context) error
}

var _ Querier = (*Queries)(nil)
