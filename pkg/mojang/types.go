package mojang

import (
	"net/http"

	"github.com/chickeniq/hypixel-go/pkg/cache"
)

type Client struct {
	cache *cache.Cache
	http  http.Client
}

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
