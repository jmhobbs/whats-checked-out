package biblionix

import (
	"fmt"
	"net/http"
)

type Client struct {
	http      *http.Client
	subdomain string
}

type ClientOption func(*Client)

func New(subdomain string, fns ...ClientOption) *Client {
	c := &Client{subdomain: subdomain}
	for _, fn := range fns {
		fn(c)
	}
	return c
}

func (c *Client) httpClient() *http.Client {
	if c.http != nil {
		return c.http
	}
	return http.DefaultClient
}

func (c *Client) URL(path string) string {
	return fmt.Sprintf("https://%s.biblionix.com%s", c.subdomain, path)
}

func WithHTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.http = client
	}
}
