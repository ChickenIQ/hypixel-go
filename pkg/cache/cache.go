package cache

import (
	"context"
	"time"
)

type Provider interface {
	Put(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, bool, error)
}

type Cache struct {
	Provider Provider
	TTL      time.Duration
}

func Do(ctx context.Context, c *Cache, key string, fn func(context.Context) ([]byte, error)) ([]byte, error) {
	if c == nil || c.Provider == nil {
		return fn(ctx)
	}

	value, found, err := c.Provider.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if found {
		return value, nil
	}

	data, err := fn(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.Provider.Put(ctx, key, data, c.TTL); err != nil {
		return nil, err
	}

	return data, nil
}
