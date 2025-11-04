package biblionix

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

type LoginResponse struct {
	XMLName xml.Name `xml:"root"`
	Session string   `xml:"session,attr"`
}

func (c *Client) Login(account, password string) (string, error) {
	resp, err := c.httpClient().PostForm(c.URL("/catalog/ajax_backend/login.xml.pl"), url.Values{
		"username": {account},
		"password": {password},
	})
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed: %s", resp.Status)
	}

	defer func() { _ = resp.Body.Close() }()

	var login LoginResponse
	err = xml.NewDecoder(resp.Body).Decode(&login)
	if err != nil {
		return "", err
	}

	return login.Session, nil
}
