package biblionix

import (
	"encoding/xml"
	"testing"
)

func Test_LoginDecoding(t *testing.T) {
	var login LoginResponse
	err := xml.Unmarshal([]byte(`<root session="session-id" />`), &login)
	if err != nil {
		t.Fatal(err)
	}
	if login.Session != "session-id" {
		t.Errorf("expected session-id, got %s", login.Session)
	}
}
