package valkey

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/valkey-io/valkey-go"
)

type Client struct {
	client valkey.Client
}

func NewClient(opts valkey.ClientOption) (*Client, error) {
	client, err := valkey.NewClient(opts)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) Put(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	if ttl < time.Millisecond {
		return fmt.Errorf("invalid ttl: %v", ttl)
	}

	cmd := c.client.B().Set().Key(key).Value(valkey.BinaryString(value)).Px(ttl).Build()

	return c.client.Do(ctx, cmd).Error()
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, bool, error) {
	result := c.client.Do(ctx, c.client.B().Get().Key(key).Build())

	value, err := result.AsBytes()
	if valkey.IsValkeyNil(err) {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	return bytes.Clone(value), true, nil
}

func (c *Client) Close() {
	c.client.Close()
}
