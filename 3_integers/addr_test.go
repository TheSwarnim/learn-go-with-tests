package integers

import (
	"testing"
	"fmt"
)

func TestAddr(t *testing.T) {
	sum := Add(4, 5)
	expected := 9

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
