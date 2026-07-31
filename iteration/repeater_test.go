package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	char := "a"
	repeatCount := 5
	repeated := Repeat(char, repeatCount)
	expected := strings.Repeat(char, repeatCount)

	if repeated != expected {
		t.Errorf("expected: %s, got: %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
