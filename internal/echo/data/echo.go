package data

import (
	"context"

	"github.com/yangchnet/skeleton/internal/echo/biz"
)

var _ biz.EchoRepo = (*echoRepo)(nil)

type echoRepo struct{} // FIXME

func NewEchoRepo() *echoRepo {
	return nil // FIXME
}

func (r *echoRepo) CreateEcho(ctx context.Context, echo *biz.Echo) (*biz.Echo, error) {
	panic("not implemented") // TODO: Implement
}

func (r *echoRepo) ListEcho(ctx context.Context, offset int64, limit int64) (*[]biz.Echo, error) {
	panic("not implemented") // TODO: Implement
}

func (r *echoRepo) UpdateEcho(ctx context.Context, message string) (*biz.Echo, error) {
	panic("not implemented") // TODO: Implement
}

func (r *echoRepo) DeleteEcho(ctx context.Context, ID int64) error {
	panic("not implemented") // TODO: Implement
}
