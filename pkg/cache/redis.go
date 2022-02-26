package cache

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yangchnet/skeleton/pkg/logger"
)

const (
	// RedisType represents the storage type as a string value.
	RedisType = "redis"
	// RedisTagPattern represents the tag pattern to be used as a key in specified storage.
	RedisTagPattern = "cache_tag_%s"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr, passwd string, db int) *RedisStore {
	return &RedisStore{client: redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})}
}

var _ CacheInterface = (*RedisStore)(nil)

// Get returns the object stored in cache if it exists.
func (r *RedisStore) Get(ctx context.Context, key string) (interface{}, bool) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Error("error redis get :", err)
		}

		return nil, false
	}

	return value, true
}

// Set populates the cache item using the given key.
func (r *RedisStore) Set(ctx context.Context, key string, value interface{}, options *Options) error {
	if err := r.client.Set(ctx, key, value, options.ExpirationValue()).Err(); err != nil {
		logger.Error(err)

		return err
	}

	if tags := options.TagsValue(); len(tags) > 0 {
		r.setTags(ctx, key, tags)
	}

	return nil
}

// Delete removes the cache item using the given key.
func (r *RedisStore) Delete(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()

	return err
}

// Invalidate invalidates cache item from given options.
func (r *RedisStore) Invalidate(ctx context.Context, options InvalidateOptions) error {
	if tags := options.TagsValue(); len(tags) > 0 {
		for _, tag := range tags {
			tagKey := fmt.Sprintf(RedisTagPattern, tag)
			cacheKeys := r.getCacheKeysForTag(ctx, tagKey)
			for _, cacheKey := range cacheKeys {
				if err := r.Delete(ctx, cacheKey); err != nil {
					logger.Error(err.Error())
				}
			}
			if err := r.Delete(ctx, tagKey); err != nil {
				logger.Error(err.Error())
			}
		}
	}

	return nil
}

// Clear resets all cache data.
func (r *RedisStore) Clear(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

// GetType returns the cache type.
func (r *RedisStore) GetType(ctx context.Context) string {
	panic("not implemented") // TODO: Implement
}

func (r *RedisStore) setTags(ctx context.Context, key string, tags []string) {
	for _, tag := range tags {
		tagKey := fmt.Sprintf(RedisTagPattern, tag)
		cacheKeys := r.getCacheKeysForTag(ctx, tagKey)
		alreadyInserted := false
		for _, cacheKey := range cacheKeys {
			if cacheKey == key {
				alreadyInserted = true
				break
			}
		}
		if !alreadyInserted {
			cacheKeys = append(cacheKeys, key)
		}
		if err := r.Set(ctx, tagKey, strings.Join(cacheKeys, ","), &Options{Expiration: 720 * time.Hour}); err != nil {
			logger.Error(ctx, err.Error())
		}
	}
}

func (r *RedisStore) getCacheKeysForTag(ctx context.Context, tagKey string) []string {
	cacheKeys := []string{}
	if result, get := r.Get(ctx, tagKey); get {
		if str, ok := result.(string); ok {
			cacheKeys = strings.Split(str, ",")
		}
	}

	return cacheKeys
}
