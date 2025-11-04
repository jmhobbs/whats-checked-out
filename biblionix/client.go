package biblionix

import (
	"fmt"
	"strings"
)

type Client struct {
	subdomain string
}

func New(subdomain string) *Client {
	return &Client{subdomain: subdomain}
}

func (c *Client) URL(path string) string {
	return fmt.Sprintf("https://%s.biblionix.com%s", c.subdomain, path)
}

func Deobfuscate(input string) string {
	return strings.ReplaceAll(input, "­", "")
}
