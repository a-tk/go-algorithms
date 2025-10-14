package math

import "testing"

func Test_RepeatedSquaringRec_Simple(t *testing.T) {
	got := RepeatedSquaringRec(3, 3)
	if got != 27 {
		t.Errorf("expected 3^3=27 got %d\n", got)
	}

	got = RepeatedSquaringRec(3, 4)
	if got != 81 {
		t.Errorf("expected 3^4=81 got %d\n", got)
	}
}
