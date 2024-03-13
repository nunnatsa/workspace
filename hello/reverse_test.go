package hello

import "testing"

func TestReverse(t *testing.T) {
	if Reverse("abc") != "cba" {
		t.Error("should be cba")
	}

	if Reverse("תובל") != "לבות" {
		t.Error("shold be לבות")
	}
}
