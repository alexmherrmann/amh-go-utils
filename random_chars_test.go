package amh_go_utils

import (
	"strings"
	"testing"
)

func TestAllCharsRightLen(t *testing.T) {
	expectedNumOfChars := 26 + 26 + 10
	if len(allChars) != expectedNumOfChars {
		t.Fatalf("Length of allchars was off: %3d != %3d", expectedNumOfChars, len(allChars))
	}
}

func TestAllCharsHasBounds(t *testing.T) {

	for _, runeToCheck := range []rune{'a', 'z', 'A', 'Z', '0', '9'} {
		if !strings.ContainsRune(string(allChars), runeToCheck) {
			t.Fatalf("allchars did not contain: %c", runeToCheck)
		}
	}
}
