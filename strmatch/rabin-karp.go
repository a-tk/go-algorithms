package strmatch

import (
	"fmt"
	"github.com/a-tk/go-algorithms/math"
)

func RabinKarpMatcher[T string | []byte](text, pattern T) (int, bool) {
	return rabinKarpMatcher(text, pattern, 256, 16777619)
}

// rabinKarpMatcher is an impl taken from CLRS. It requires knowledge of which modulo to use
//
//	q should be positive for % to correctly be modulo
//	d is the radix to use, typically the size of the alphabet (e.g. 256)
//	with a d-ary alphabet (0,1..d-1) choose q so that dq fits in a computer word
func rabinKarpMatcher[T string | []byte](text, pattern T, d, q uint32) (int, bool) {

	n := len(text)
	m := len(pattern)
	h := math.RepeatedSquaring(d, uint32(m-1)) % q
	var p, t uint32 = 0, 0

	for i := 0; i < m; i++ {
		p = (d*p + uint32(pattern[i])) % q
		t = (d*t + uint32(text[i])) % q
	}
	for s := 0; s < n-(m-1); s++ {
		if p == t {
			if string(pattern) == string(text[s:s+m]) {
				return s, true
			}
		}
		if s < n-m {
			t = (d*(t-uint32(text[s])*h) + uint32(text[s+m])) % q
		}
	}
	return -1, false
}

// base 10 for testing
func rabinKarpIntMatcher(T, P []byte) {

	n := len(T)
	m := len(P)
	h := math.RepeatedSquaringRec(10, m-1)
	p := 0
	t := 0

	for i := 0; i < m; i++ {
		p = 10*p + int(P[i])
		t = 10*t + int(T[i])
	}
	for s := 0; s < n-(m-1); s++ {
		if p == t {
			if string(P[0:m]) == string(T[s:s+m]) {
				fmt.Printf("pattern occurs with shift %d\n", s)
			}
		}
		if s < n-m {
			t = 10*(t-int(T[s])*h) + int(T[s+m])
		}
	}
}

func simpleRKHash[T string | []byte](s T) int {

	var h int

	for i := 0; i < len(s); i++ {
		h += int(s[i])
	}

	return h
}

func RabinKarpSimpleMatcher[T string](s, p T) (i int, found bool) {

	n := len(s)
	m := len(p)
	hpattern := simpleRKHash(p)
	hs := simpleRKHash(s[0:m]) // usually m-1, however, go slices are open interval [0:m)
	for j := 0; j < n-m; j++ {
		hs = hs - int(s[j]) + int(s[j+m])
		if hs == hpattern {
			if s[j:j+m] == p {
				return j, true
			}
		}
	}
	return -1, false
}
