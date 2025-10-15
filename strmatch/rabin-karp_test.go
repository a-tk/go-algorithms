package strmatch

import "testing"

func Test_rKSimpleWiki(t *testing.T) {
	got := RabinKarpMatcher("abaa", "aa")

	if got != 2 {
		t.Errorf("Fail!\n")
	}
}

func Test_rKSimple(t *testing.T) {
	rabinKarpMatcher("aa aa", "aa", 256, 16777619) // 16777216 is 2^32-1 / 256. 16777619 is a close prime
}

func Test_rKIntSimple(t *testing.T) {
	text := []byte{3, 1, 4, 1, 5}
	pattern := []byte{1, 5}

	rabinKarpIntMatcher(text, pattern)
}
