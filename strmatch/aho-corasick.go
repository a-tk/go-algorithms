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
	f      []int // from a state go to another state
	output map[int][]T
}

func constructFailure[T string | []byte](g *graph.GraphAutomata[int, byte], output map[int][]T) ([]int, map[int][]T) {
	f := make([]int, g.Len()) // an array with a spot for every state
	q := queue.New[int](100)  // TODO queue can hold as many letters as transitions from state zero
	// what should that number be?

	// silly, this is just getting all the valid transitions from 0
	//as, _ := g.GetTransitionsW(0)
	//for _, a := range as {
	//	s, _ := g.GetTransition(0, a)
	//	q.Enqueue(s)
	//	f[s] = 0
	//}

	states, _ := g.StateTransitions(0)
	for _, s := range states {
		if s != 0 {
			q.Enqueue(s)
			f[s] = 0
		}
	}

	for !q.Empty() {
		r, _ := q.Dequeue()
		// for each a such that g(r, a) = s != fail do

		as, _ := g.WTransitions(r)
		for _, a := range as {
			// have a, now get s
			s, _ := g.GetTransition(r, a)
			q.Enqueue(s)
			state := f[r]

			for _, ok := g.GetTransition(state, a); !ok; state = f[state] {
			}
			t, _ := g.GetTransition(state, a)
			f[s] = t
			o, ok := output[f[s]]
			if ok {
				output[s] = append(output[s], o...)
			}
		}
	}
	return f, output
}

func constructGoto[T string | []byte](patterns []T) (
	g *graph.GraphAutomata[int, byte],
	output map[int][]T,
) {
	g = graph.NewGraphAutomata[int, byte]()
	g.AddState(0)
	output = make(map[int][]T)
	newState := 0

	enter := func(pattern T) {
		state := 0
		j := 0
		// walk to a new state. For example if "he" already exists and we are adding "hers" or "his"

		ok := true
		for ok {
			var s int
			s, ok = g.GetTransition(state, pattern[j])
			if ok {
				state = s
				j++
			}
		}
		for p := j; p < len(pattern); p++ {
			newState = newState + 1
			g.AddState(newState)
			g.AddTransition(state, newState, pattern[p])
			state = newState

			// for all a such that g(0,a) == fail, do g(0, a) = 0
			if _, ok = g.GetTransition(0, pattern[p]); !ok {
				g.AddTransition(0, 0, pattern[p])
			}
		}
		// add this to the output list
		output[state] = append(output[state], pattern)
	}
	for i := 0; i < len(patterns); i++ {
		enter(patterns[i])
	}

	return
}

func NewMatcher[T string | []byte](patterns []T) *Matcher[T] {
	g, output := constructGoto(patterns)
	f, output := constructFailure(g, output)

	return &Matcher[T]{
		g:      g,
		output: output,
		f:      f,
	}
}

func (m *Matcher[T]) Match(x T) {
	state, prev := 0, 0
	for i := 0; i < len(x); i++ {
		s, ok := m.g.GetTransition(state, x[i])
		// without adding a negative transition in the automata, we should check to see if the previous state was also not zero
		for (!ok && s != 0) || (!ok && prev != 0) {
			prev = state
			state = m.f[state]
			s, ok = m.g.GetTransition(state, x[i])
		}
		prev = state
		state, _ = m.g.GetTransition(state, x[i])
		if p, matched := m.output[state]; matched {
			fmt.Printf("Match! %d: %s\n", i, p)
		}
	}
}
