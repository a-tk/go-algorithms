package strmatch

import "testing"

func Test_Ushers(t *testing.T) {

	// use {he, she, his, hers}
	// string x to be searched is ushers

	m := NewMatcher([]string{"he", "she", "his", "hers"})

	m.Match("ushersx")
}
