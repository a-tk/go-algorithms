package strmatch

import (
	"fmt"
	"github.com/a-tk/go-datastructures/graph"
	"github.com/a-tk/go-datastructures/queue"
)

// use {he, she, his, hers}
// string x to be searched is ushers

type Matcher[T string | []byte] struct {
	g      *graph.GraphAutomata[int, byte]
	f      map[int]int // from a state go to another state
	output map[int][]T
}

func constructFailure[T string | []byte](g *graph.GraphAutomata[int, byte], output map[int][]T) (f map[int]int) {
	f = make(map[int]int)
	q := queue.New[int](100) // TODO queue can hold as many letters as transitions from state zero
	// what should that number be?

	// silly, this is just getting all the valid transitions from 0
	//as, _ := g.GetTransitionsW(0)
	//for _, a := range as {
	//	s, _ := g.GetTransition(0, a)
	//	q.Enqueue(s)
	//	f[s] = 0
	//}

	states, _ := g.GetTransitionsStates(0)
	for _, s := range states {
		q.Enqueue(s)
		f[s] = 0
	}

	for !q.Empty() {
		r, _ := q.Dequeue()
		// for each a such that g(r, a) = s != fail do

		as, _ := g.GetTransitionsW(r)
		for _, a := range as {
			// have a, now get s
			s, _ := g.GetTransition(r, a)
			q.Enqueue(s)
			state := f[r]

			for _, ok := g.GetTransition(state, a); !ok; state = f[state] {
			}
			t, _ := g.GetTransition(state, a)
			f[s] = t
			output[s] = append(output[s], output[f[s]]...)
		}
	}
	return f
}

func constructGoto[T string | []byte](patterns []T) (
	g *graph.GraphAutomata[int, byte],
	output map[int][]T,
) {
	g = graph.NewGraphAutomata[int, byte]()
	output = make(map[int][]T)
	newState := 0

	enter := func(pattern T) {
		state := 0
		j := 0
		// walk to a new state. For example if "he" already exists and we are adding "hers"
		for _, ok := g.GetTransition(state, pattern[j]); ok; j++ {
			state, ok = g.GetTransition(state, pattern[j]) // TODO: yuck
		}
		for p := j; p < len(pattern); p++ {
			newState = newState + 1
			g.AddTransition(state, newState, pattern[p])
			state = newState
		}
		// add this to the output list
		output[state] = append(output[state], pattern)
	}
	for i := 0; i < len(patterns); i++ {
		enter(patterns[i])
	}
	// for all a such that g(0,a) == fail, do g(0, a) = 0
	// this is basically saying for all letters that would fail to
	//  transition from 0, make them instead transition to zero
	// I think this is unnecessary, as the GraphAutomata will simply fail, and the state won't change

	return
}

func NewMatcher[T string | []byte](patterns []T) *Matcher[T] {
	g, output := constructGoto(patterns)
	return &Matcher[T]{
		g:      g,
		output: output,
	}
}

func (m *Matcher[T]) Match(x T) {
	state := 0
	for i := 0; i < len(x); i++ {
		for _, fail := m.g.GetTransition(state, x[i]); fail; state = m.f[state] {
		}
		state, _ = m.g.GetTransition(state, x[i])
		if p, matched := m.output[state]; matched {
			fmt.Printf("Match! %d: %s", i, p) // TODO: p is actually a slice of strings
		}
	}
}
