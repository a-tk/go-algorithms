package strmatch

import "testing"

func Test_rKSimpleWiki(t *testing.T) {
	_, got := RabinKarpMatcher("abaa", "aa")

	if got != true {
		t.Errorf("error!\n")
	}
}

func Test_rKSimple(t *testing.T) {
	rabinKarpMatcher("aa aa", "aa", 256, 16777619) // stealing 16777619 from go's byte alg

}
