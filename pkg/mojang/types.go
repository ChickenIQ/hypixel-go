package mojang

import "net/http"

type Client struct {
	http http.Client
}

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
