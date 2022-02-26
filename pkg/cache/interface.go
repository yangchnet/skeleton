package cache

import "context"

// CacheInterface is the interface for all available stores.
type CacheInterface interface {
	// Get returns the object stored in cache if it exists
	Get(ctx context.Context, key string) (interface{}, bool)
	// Set populates the cache item using the given key
	Set(ctx context.Context, key string, value interface{}, options *Options) error
	// Delete removes the cache item using the given key
	Delete(ctx context.Context, key string) error
	// Invalidate invalidates cache item from given options
	Invalidate(ctx context.Context, options InvalidateOptions) error
	// Clear resets all cache data
	Clear(ctx context.Context) error
	// GetType returns the cache type
	GetType(ctx context.Context) string
}
