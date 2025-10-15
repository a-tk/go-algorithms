package strmatch

import "testing"

func TestNaiveStringMatcher(t *testing.T) {
	i, got := NaiveStringMatcher("abaa", "aa")

	if i != 2 || got == false {
		t.Errorf("Fail, i = %d, got = %t\n", i, got)
	}

	i, got = NaiveStringMatcher("abab", "aa")

	if i != -1 || got {
		t.Errorf("Fail, i = %d, got = %t\n", i, got)
	}
}
