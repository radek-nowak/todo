package reverser

import (
	"testing"
)

func TestReverse(t *testing.T) {
	input := "abc123"
	expected := "321cba"
	actual := Reverse(input)

	if expected != actual {
		t.Errorf("actual %s, expected %s", actual, expected)
	}

}
