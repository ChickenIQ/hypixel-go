package mojang

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func NewClient() *Client {
	return &Client{http: http.Client{Timeout: 10 * time.Second}}
}

func (c *Client) GetProfile(ctx context.Context, username string) (*Profile, error) {
	url := "https://api.minecraftservices.com/minecraft/profile/lookup/name/" + username

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mojang returned %s", resp.Status)
	}

	var profile Profile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}
