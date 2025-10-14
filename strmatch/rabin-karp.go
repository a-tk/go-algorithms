package strmatch

import (
	"fmt"
	"github.com/a-tk/go-algorithms/math"
)

// rabinKarpMatcher is an impl taken from CLRS. It requires knowledge of which modulo to use
// q should be positive for % to correctly be modulo
// d is the radix to use, typically the size of the alphabet
// with a d-ary alphabet (0,1..d-1) choose q so that dq fits in a computer word
func rabinKarpMatcher(T, P string, d, q int) {

	n := len(T)
	m := len(P)
	h := math.RepeatedSquaringRec(d, m-1)
	p := 0
	t := 0

	for i := 0; i < m; i++ {
		p = (d*p + int(P[i])) % q // int cast
		t = (d*t + int(T[i])) % q
	}
	for s := 0; s < n-(m-1); s++ {
		if p == t {
			if P == T[s:s+m] {
				fmt.Printf("pattern occurs with shift %d\n", s)
			}
		}
		if s < n-m {
			t = (d*(t-int(T[s+1])*h) + int(T[s+m])) % q
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

func RabinKarpMatcher[T string](s, p T) (i int, found bool) {

	n := len(s)
	m := len(p)
	hpattern := simpleRKHash(p)
	hs := simpleRKHash(s[0 : m-1])
	for j := 0; j < n-m; j++ {
		if hs == hpattern {
			if s[j:j+m-1] == p[0:m-1] {
				return j, true
			}
		}
		hs = hs - int(s[j]) + int(s[j+m])
	}
	return -1, false
}
