package data

import (
	"context"

	"github.com/yangchnet/skeleton/internal/echo/biz"
	"github.com/yangchnet/skeleton/pkg/logger"
)

var _ biz.EchoRepo = (*Data)(nil)

func (r *Data) CreateEcho(ctx context.Context, echo *biz.Echo) (*biz.Echo, error) {
	echoModel, err := r.db.Echo.
		Create().
		SetMessage(echo.Message).
		SetEchoMessage(echo.EchoMessage).
		Save(ctx)
	if err != nil {
		logger.Errorf("create echo failed: ", err)
		return nil, err
	}

	return &biz.Echo{
		ID:          echoModel.ID,
		Message:     echo.Message,
		EchoMessage: echo.EchoMessage,
	}, nil
}

func (r *Data) ListEcho(ctx context.Context, offset int64, limit int64) (*[]biz.Echo, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Data) UpdateEcho(ctx context.Context, message string) (*biz.Echo, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Data) DeleteEcho(ctx context.Context, ID int64) error {
	panic("not implemented") // TODO: Implement
}
