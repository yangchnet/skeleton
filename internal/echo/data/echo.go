package data

import (
	"context"
	"time"

	"github.com/yangchnet/skeleton/internal/echo/biz"
	"github.com/yangchnet/skeleton/internal/echo/data/ent"
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

func (r *Data) ListEcho(ctx context.Context, offset int64, limit int64) ([]*biz.Echo, error) {
	echoList, err := r.db.Echo.
		Query().
		Offset(int(offset)).
		Limit(int(limit)).
		All(ctx)
	if err != nil {
		logger.Error("list echo failed: ", err)
	}

	list := make([]*biz.Echo, len(echoList))
	for i, e := range echoList {
		list[i] = &biz.Echo{
			ID:          e.ID,
			Message:     e.Message,
			EchoMessage: e.EchoMessage,
			CreateTime:  &e.CreateTime,
		}
	}

	return list, nil
}

func (r *Data) UpdateEcho(ctx context.Context, echo *biz.Echo) (*biz.Echo, error) {
	e, err := r.db.Echo.
		UpdateOneID(echo.ID).
		SetMessage(echo.Message).
		SetEchoMessage(echo.Message).
		SetUpdateTime(time.Now()).
		Save(ctx)
	if err != nil {
		logger.Error("update echo failed: ", err)
		if ent.IsNotFound(err) {
			return nil, biz.ErrNotFound
		}
	}

	return &biz.Echo{
		ID:          e.ID,
		Message:     e.Message,
		EchoMessage: e.EchoMessage,
	}, nil
}

func (r *Data) DeleteEcho(ctx context.Context, ID int64) error {
	// 直接删除
	// if err := r.db.Echo.
	// 	DeleteOneID(int(ID)).
	// 	Exec(ctx); err != nil {
	// 	logger.Error("error delete echo record: ", err)

	// 	return err
	// }

	// return nil

	// 暂时标记为删除，等待服务空闲时删除
	_, err := r.db.Echo.
		UpdateOneID(int(ID)).
		SetDeleted(true).
		SetDeleteTime(time.Now()).
		Save(ctx)
	if err != nil {
		logger.Error("error delete echo record: ", err)

		if ent.IsNotFound(err) {
			return biz.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *Data) GetEcho(ctx context.Context, ID int64) (*biz.Echo, error) {
	panic("not implemented") // TODO: Implement
}
