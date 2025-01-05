package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}

func TestSubstract(t *testing.T) {
	if Substract(5, 3) != 2 {
		t.Error("5 - 3 did not equal 2")
	}
}
