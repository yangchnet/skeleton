package biz

import "context"

// Echo is a domain object that handle an echo.
type Echo struct {
	ID          int64
	Message     string
	EchoMessage string
}

// EchoRepo is a interface to access dao.
type EchoRepo interface {
	CreateEcho(ctx context.Context, echo *Echo) (*Echo, error)
	ListEcho(ctx context.Context, offset, limit int64) (*[]Echo, error)
	UpdateEcho(ctx context.Context, message string) (*Echo, error)
	DeleteEcho(ctx context.Context, ID int64) error
}

// EchoCase handle EchoRepo interface.
type EchoCase struct {
	repo EchoRepo
}

// NewEchoCase creates a new EchoCase.
func NewEchoCase(repo EchoRepo) *EchoCase {
	return &EchoCase{repo: repo}
}
