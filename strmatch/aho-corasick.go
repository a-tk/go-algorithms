package strmatch

import "fmt"

// use {he, she, his, hers}
// string x to be searched is ushers

type Matcher[T string | []byte] struct {
	g        func(state int, a T) (next int, fail bool)
	f        func(state int) (prev int)
	output   func(state int) (pattern T, fail bool)
	patterns []T
}

func constructGoto[T string | []byte](patterns []T) (
	g func(state int, a T) (next int, fail bool),
	output func(state int) (pattern T, fail bool),
) {

}

func NewMatcher[T string | []byte](patterns []T) *Matcher[T] {
	return nil
}

func (m *Matcher[T]) Match(x T) {
	state := 0
	for i := 0; i < len(x); i++ {
		for _, fail := m.g(state, x[i]); fail; state = m.f(state) {
		}
		state, _ = m.g(state, x[i])
		if p, matched := m.output(state); matched {
			fmt.Printf("Match! %d: %s", i, p)
		}
	}
}
