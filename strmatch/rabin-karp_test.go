package strmatch

import "testing"

func Test_rKSimple(t *testing.T) {
	rabinKarpMatcher("hello world", "world", 256, 16777619) // stealing 16777619 from go std_lib. See bytealg.go
}
