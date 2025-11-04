package biblionix

import "strings"

func deobfuscate(input string) string {
	return strings.ReplaceAll(input, "\u00ad", "")
}
