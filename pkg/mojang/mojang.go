package mojang

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chickeniq/hypixel-go/pkg/cache"
)

func NewClient(cache *cache.Cache) *Client {
	return &Client{
		cache: cache,
		http:  http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) GetProfile(ctx context.Context, username string) (*Profile, error) {
	url := "https://api.minecraftservices.com/minecraft/profile/lookup/name/" + username

	body, err := cache.Do(ctx, c.cache, url, func(ctx context.Context) ([]byte, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

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
			return nil, fmt.Errorf("mojang returned %s: %s", resp.Status, body)
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	var profile Profile
	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}
