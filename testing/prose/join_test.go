package prose

import (
	"fmt"
	"testing"
)

func TestJoinWithCommasNullElement(t *testing.T) {
	list := []string{}
	want := ""
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestJoinWithCommasOneElement(t *testing.T) {
	list := []string{"apples"}
	want := "apples"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestJoinWithCommasTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestJoinWithCommasThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
