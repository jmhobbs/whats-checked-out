package biblionix

import "testing"

func Test_Deobfuscate(t *testing.T) {
	deobfuscated := Deobfuscate("Real Ja­pane­se ­cook­ing")
	if deobfuscated != "Real Japanese cooking" {
		t.Errorf("expected 'Real Japanese cooking', got %q", deobfuscated)
	}
}
