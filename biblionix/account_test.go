package biblionix_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Account(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rt, client := mockedClient(t)

		rt.
			EXPECT().
			RoundTrip(mock.MatchedBy(func(r *http.Request) bool {
				// TODO: Read body and make sure it has session id
				return r.URL.Path == "/catalog/ajax_backend/account.xml.pl"
			})).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(minimalSuccessXML)),
			}, nil).
			Once()

		account, err := client.Account("session-id")
		assert.NoError(t, err)
		assert.Equal(t, "12345", account.Card)
		assert.Equal(t, "12346", account.LinkedCard[0].CardNumber)
		// check deobfuscation
		assert.Equal(t, "Real Japanese cooking: the only Japanese cookbook you will ever need!", account.Item[0].Title)
		assert.Equal(t, "Itoh, Makiko", account.Item[0].Author)
	})

	t.Run("http error", func(t *testing.T) {
		rt, client := mockedClient(t)

		rt.
			EXPECT().
			RoundTrip(mock.Anything).
			Return(&http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewBufferString(minimalSuccessXML)),
			}, nil).
			Once()

		_, err := client.Account("session-id")
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

		_, err := client.Account("session-id")
		assert.Error(t, err)
	})
}

//nolint:all
const minimalSuccessXML string = `<root card="12345" card_printable="P12345">
	<linked_card id="12345678" card_number="12346" />
	<item
		id="1234567890"
		title_sort="real japanese cooking"
		title="Real Ja­pane­se ­cook­ing: the on­ly Ja­pane­se ­cook­book y­ou will ev­er need!"
		medium=""
		author="Ito­h, Makiko"
		author_sort="Itoh, Makiko"
		out_raw="2025-10-17"
		due_raw="2025-11-11"
		out="10&#8209;17&#8209;2025"
		due="11&#8209;11&#8209;2025"
		renewable="1"
		indicate_reserve="0"
		overdue="0"
		biblio="2541757192" />
</root>`
