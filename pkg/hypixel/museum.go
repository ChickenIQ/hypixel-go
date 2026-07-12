package hypixel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	sc "github.com/DuckySoLucky/SkyCrypt-Types"
)

func (c *Client) GetMuseum(ctx context.Context, profileID, memberID string) (*sc.Museum, error) {
	body, err := c.get(ctx, "/skyblock/museum?profile="+url.QueryEscape(profileID))
	if err != nil {
		return nil, err
	}

	var data struct {
		Members map[string]sc.Museum `json:"members"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("decode Hypixel museum response: %w", err)
	}

	museum, ok := data.Members[parseUUID(memberID)]
	if !ok {
		return nil, nil
	}

	return &museum, nil
}
