package hypixel

import (
	"context"
	"net/url"
	"strings"
)

func (c *Client) getWithUUID(ctx context.Context, path, uuid string) ([]byte, error) {
	return c.get(ctx, path+"?uuid="+url.QueryEscape(uuid))
}

func parseUUID(uuid string) string {
	return strings.ToLower(strings.ReplaceAll(uuid, "-", ""))
}
