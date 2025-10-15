package strmatch

func NaiveStringMatcher[T string | []byte](text, pattern T) (int, bool) {
	n := len(text)
	m := len(pattern)
	for s := 0; s < n-m+1; s++ {
		if string(pattern) == string(text[s:s+m]) {
			return s, true
		}
	}
	return -1, false
}
