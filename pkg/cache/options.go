package cache

import "time"

// Options represents the cache store available options.
type Options struct {
	// Expiration allows to specify an expiration time when setting a value
	Expiration time.Duration
	// Tags allows to specify associated tags to the current value
	Tags []string
}

// ExpirationValue returns the expiration option value.
func (o Options) ExpirationValue() time.Duration {
	return o.Expiration
}

// TagsValue returns the tags option value.
func (o Options) TagsValue() []string {
	return o.Tags
}

// InvalidateOptions represents the cache invalidation available options.
type InvalidateOptions struct {
	// Tags allows to specify associated tags to the current value
	Tags []string
}

// TagsValue returns the tags option value.
func (o InvalidateOptions) TagsValue() []string {
	return o.Tags
}
