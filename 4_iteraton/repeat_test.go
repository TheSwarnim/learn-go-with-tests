package iteraton

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("p", 5)
	expected := "ppppp"

	if repeat != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("p", 5)
	}
}