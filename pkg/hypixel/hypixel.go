package hypixel

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	http   http.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http:   http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	u := fmt.Sprintf("%s/%s", "https://api.hypixel.net/v2", strings.TrimPrefix(path, "/"))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("API-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hypixel returned %s: %s", resp.Status, body)
	}

	return body, nil
}
