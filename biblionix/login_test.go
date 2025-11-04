package biblionix_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Login(t *testing.T) {
	t.Run("successful login", func(t *testing.T) {
		rt, client := mockedClient(t)

		rt.
			EXPECT().
			RoundTrip(mock.Anything).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`<root session="session-id" />`)),
			}, nil).
			Once()

		sessionId, err := client.Login("account", "password")
		assert.NoError(t, err)
		assert.Equal(t, "session-id", sessionId)
	})

	t.Run("http error", func(t *testing.T) {
		rt, client := mockedClient(t)

		rt.
			EXPECT().
			RoundTrip(mock.Anything).
			Return(&http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewBufferString(`<root session="session-id" />`)),
			}, nil).
			Once()

		_, err := client.Login("account", "password")
		assert.Error(t, err)
	})

	t.Run("bad response", func(t *testing.T) {
		rt, client := mockedClient(t)

		rt.
			EXPECT().
			RoundTrip(mock.Anything).
			Return(&http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewBufferString(`oh no invalid xml`)),
			}, nil).
			Once()

		_, err := client.Login("account", "password")
		assert.Error(t, err)
	})
}
