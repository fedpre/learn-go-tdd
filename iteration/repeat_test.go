package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q received %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 6)
	fmt.Println(result)
	// Output: aaaaaa
}

func TestStrings(t *testing.T) {
	t.Run("count each ocurrence of a substring in a string", func(t *testing.T) {
		counter := strings.Count("hello world", "o")
		expected := 2

		if counter != expected {
			t.Errorf("expected %d got %d", expected, counter)
		}
	})
}
