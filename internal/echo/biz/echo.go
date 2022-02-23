package biz

import "context"

type Echo struct {
	ID          int64
	Message     string
	EchoMessage string
}

type EchoRepo interface {
	CreateEcho(ctx context.Context, echo *Echo) (*Echo, error)
	ListEcho(ctx context.Context, offset, limit int64) (*[]Echo, error)
	UpdateEcho(ctx context.Context, message string) (*Echo, error)
	DeleteEcho(ctx context.Context, ID int64) error
}

type EchoCase struct {
	repo EchoRepo
}

func NewEchoCase(repo EchoRepo) *EchoCase {
	return nil // TODO: implementation
}
