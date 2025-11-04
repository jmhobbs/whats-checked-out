package biblionix_test

import (
	"net/http"
	"testing"

	"github.com/jmhobbs/whats-checked-out/biblionix"
	mockHttp "github.com/jmhobbs/whats-checked-out/mocks/http"

	"github.com/stretchr/testify/assert"
)

func mockedClient(t *testing.T) (*mockHttp.MockRoundTripper, *biblionix.Client) {
	rt := mockHttp.NewMockRoundTripper(t)
	httpClient := &http.Client{Transport: rt}
	client := biblionix.New("mydomain", biblionix.WithHTTPClient(httpClient))
	return rt, client
}

func Test_Client_URL(t *testing.T) {
	assert.Equal(
		t,
		"https://mydomain.biblionix.com/test/path",
		biblionix.New("mydomain").URL("/test/path"),
	)
}
