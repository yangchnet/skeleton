package biz

import (
	"context"
	"errors"
	"time"
)

var ErrNotFound = errors.New("echo record not found")

// Echo is a domain object that handle an echo.
type Echo struct {
	ID          int
	Message     string
	EchoMessage string
	CreateTime  *time.Time
}

// EchoRepo is a interface to access dao.
type EchoRepo interface {
	CreateEcho(ctx context.Context, echo *Echo) (*Echo, error)
	ListEcho(ctx context.Context, offset, limit int64) ([]*Echo, error)
	UpdateEcho(ctx context.Context, echo *Echo) (*Echo, error)
	DeleteEcho(ctx context.Context, ID int64) error
	GetEcho(ctx context.Context, ID int64) (*Echo, error)
}

// EchoCase handle EchoRepo interface.
type EchoCase struct {
	repo EchoRepo
}

// NewEchoCase creates a new EchoCase.
func NewEchoCase(ctx context.Context, repo EchoRepo) *EchoCase {
	return &EchoCase{repo: repo}
}

func (uc *EchoCase) CreateEcho(ctx context.Context, message string) (*Echo, error) {
	return uc.repo.CreateEcho(ctx, &Echo{
		Message:     message,
		EchoMessage: message,
	})
}

func (uc *EchoCase) ListEcho(ctx context.Context, offset, limit int64) ([]*Echo, error) {
	return uc.repo.ListEcho(ctx, offset, limit)
}

func (uc *EchoCase) UpdateEcho(ctx context.Context, id int64, message string) (*Echo, error) {
	echo, err := uc.repo.UpdateEcho(ctx, &Echo{
		ID:      int(id),
		Message: message,
	})
	if err != nil {
		return nil, err
	}

	return echo, nil
}

func (uc *EchoCase) DeleteEcho(ctx context.Context, id int64) error {
	if err := uc.repo.DeleteEcho(ctx, id); err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}

		return err
	}

	return nil
}

func (uc *EchoCase) GetEcho(ctx context.Context, id int64) (*Echo, error) {
	return uc.repo.GetEcho(ctx, id)
}
