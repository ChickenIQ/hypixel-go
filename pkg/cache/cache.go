package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeyaside"
)

type Cache struct {
	client  valkeyaside.CacheAsideClient
	ttl     time.Duration
	timeout time.Duration
}

func NewCache(connectionString string, ttl, timeout time.Duration) (*Cache, error) {
	if ttl < time.Millisecond {
		return nil, fmt.Errorf("invalid cache TTL: %v", ttl)
	}

	if timeout < time.Millisecond {
		return nil, fmt.Errorf("invalid cache timeout: %v", timeout)
	}

	options, err := valkey.ParseURL(connectionString)
	if err != nil {
		return nil, fmt.Errorf("parse Valkey URL: %w", err)
	}

	client, err := valkeyaside.NewClient(valkeyaside.ClientOption{ClientOption: options})
	if err != nil {
		return nil, fmt.Errorf("create valkey client: %w", err)
	}

	return &Cache{client: client, ttl: ttl, timeout: timeout}, nil
}

func (c *Cache) Do(ctx context.Context, key string, fn func(context.Context) ([]byte, error)) ([]byte, error) {
	if c == nil || c.client == nil {
		return fn(ctx)
	}

	cacheCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	value, err := c.client.Get(cacheCtx, c.ttl, key, func(ctx context.Context, _ string) (string, error) {
		data, err := fn(ctx)
		return string(data), err
	})
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (c *Cache) Close() {
	if c != nil && c.client != nil {
		c.client.Close()
	}
}
