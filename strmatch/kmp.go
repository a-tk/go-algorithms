package strmatch

// pi maps 1-m in the pattern to 0-m-1
// pi[q] is the length of the longest prefix of pattern that is a proper suffix of P[:q]
func computePrefixFunction(pattern string) []int {

	var pi = make([]int, len(pattern))
	pi[0] = -1
	k := -1
	for q := 1; q < len(pattern); q++ {
		for k >= 0 && pattern[k+1] != pattern[q] {
			k = pi[k]
		}
		if pattern[k+1] == pattern[q] {
			k = k + 1
		}
		pi[q] = k
	}
	return pi
}

func StrMatchKmpFirst(input []byte, pattern string) (i int, found bool) {
	pi := computePrefixFunction(pattern)
	q := 0                           // number of characters matched
	for i = 0; i < len(input); i++ { //scan from left to right
		for q >= 0 && pattern[q+1] != input[i] {
			q = pi[q] //next character does not match
		}
		if pattern[q+1] == input[i] {
			q = q + 1 //next character matches
		}
		if q == len(pattern)-1 { //is all the pattern matched?
			// no need to search further, return i
			q = pi[q] //look for next match
			return i - len(pattern), true
		}
	}
	return i, false
}

func StrMatchKmpAll(input []byte, pattern string) (is []int, found bool) {
	pi := computePrefixFunction(pattern)
	q := 0                            // number of characters matched
	for i := 0; i < len(input); i++ { //scan from left to right
		for q >= 0 && pattern[q+1] != input[i] {
			q = pi[q] //next character does not match
		}
		if pattern[q+1] == input[i] {
			q = q + 1 //next character matches
		}
		if q == len(pattern)-1 { //is all the pattern matched?
			// no need to search further, return i
			is = append(is, i-len(pattern))
			q = pi[q] //look for next match
		}
	}
	return is, false
}
