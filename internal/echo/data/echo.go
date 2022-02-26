package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yangchnet/skeleton/internal/echo/biz"
	"github.com/yangchnet/skeleton/internal/echo/data/ent"
	"github.com/yangchnet/skeleton/internal/echo/data/ent/echo"
	"github.com/yangchnet/skeleton/pkg/cache"
	"github.com/yangchnet/skeleton/pkg/logger"
)

var _ biz.EchoRepo = (*Data)(nil)

var echoCacheKey = func(id int) string {
	return fmt.Sprintf("%s%d", "echo_record_key_", id)
}

// CreateEcho creates a new echo record and cache it by key id.
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

	cacheKey := echoCacheKey(echoModel.ID)
	r.setEchoCache(ctx, echoModel, cacheKey)

	return &biz.Echo{
		ID:          echoModel.ID,
		Message:     echo.Message,
		EchoMessage: echo.EchoMessage,
	}, nil
}

// ListEcho returns a list of Echo.
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

// UpdateEcho updates the echo and set cache.
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

	cacheKey := echoCacheKey(echo.ID)
	r.setEchoCache(ctx, e, cacheKey)

	return &biz.Echo{
		ID:          e.ID,
		Message:     e.Message,
		EchoMessage: e.EchoMessage,
	}, nil
}

// DeleteEcho delete echo record.
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
	e, err := r.db.Echo.
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

	cacheKey := echoCacheKey(e.ID)
	r.delEchoCache(ctx, cacheKey)

	return nil
}

// GetEcho get echo record by id.
func (r *Data) GetEcho(ctx context.Context, ID int64) (*biz.Echo, error) {
	var e *ent.Echo
	var err error
	cacheKey := echoCacheKey(int(ID))
	e, found := r.getEchoCache(ctx, cacheKey)
	if found {
		return &biz.Echo{
			ID:          e.ID,
			Message:     e.Message,
			EchoMessage: e.EchoMessage,
			CreateTime:  &e.CreateTime,
		}, nil
	}

	e, err = r.db.Echo.
		Query().
		Where(echo.ID(int(ID))).
		Only(ctx)
	if err != nil {
		logger.Error("error querying echo record: ", err)
		if ent.IsNotFound(err) {
			return nil, biz.ErrNotFound
		}

		return nil, err
	}

	r.setEchoCache(ctx, e, cacheKey)

	return &biz.Echo{
		ID:          e.ID,
		Message:     e.Message,
		EchoMessage: e.EchoMessage,
		CreateTime:  &e.CreateTime,
	}, nil
}

func (r *Data) setEchoCache(ctx context.Context, echo *ent.Echo, key string) {
	marshal, err := json.Marshal(echo)
	if err != nil {
		logger.Error("error marshal echo record: ", err)

		return
	}

	err = r.cache.Set(ctx, key, marshal, &cache.Options{
		Expiration: time.Hour * 2,
	})
	if err != nil {
		logger.Error("error set echo record: ", err)

		return
	}
}

func (r *Data) delEchoCache(ctx context.Context, key string) {
	_ = r.cache.Delete(ctx, key)
}

func (r *Data) getEchoCache(ctx context.Context, key string) (echo *ent.Echo, found bool) {
	e, found := r.cache.Get(ctx, key)
	if !found {
		return nil, false
	}

	e_ := e.(string)

	if err := json.Unmarshal([]byte(e_), &echo); err != nil {
		logger.Warn("error unmarshalling echo from cache ", err)

		return nil, false
	}

	return echo, true
}
