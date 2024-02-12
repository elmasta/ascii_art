package asciiArt

import (
    "testing"
	"strconv"
)

func TestTerminalSize(t *testing.T) {
	want := 139
	width := TerminalSize()

    if want != width {
        t.Fatalf(`We want %q but we get %#q`, strconv.Itoa(want), strconv.Itoa(width))
    }
}
