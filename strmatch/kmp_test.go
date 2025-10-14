package strmatch

import (
	"reflect"
	"testing"
)

func TestComputePrefixFunction(t *testing.T) {
	tests := []struct {
		pattern string
		want    []int
	}{
		{"a", []int{-1}},
		{"aa", []int{-1, 0}},
		{"ab", []int{-1, -1}},
		{"ababaca", []int{-1, -1, 0, 1, 2, -1, 0}},
		{"aaaa", []int{-1, 0, 1, 2}},
	}

	for _, tt := range tests {
		got := computePrefixFunction(tt.pattern)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("computePrefixFunction(%q) = %v, want %v", tt.pattern, got, tt.want)
		}
	}
}

func TestStrMatchKmpFirst(t *testing.T) {
	tests := []struct {
		input   string
		pattern string
		wantPos int
		wantOk  bool
	}{
		{"hello world", "world", 6, true},
		{"aaaaa", "aa", 0, true},
		{"abc", "d", 3, false},
		{"abcabcabc", "abc", 0, true},
		{"mississippi", "issi", 1, true},
	}

	for _, tt := range tests {
		gotPos, gotOk := StrMatchKmpFirst([]byte(tt.input), tt.pattern)
		if gotPos != tt.wantPos || gotOk != tt.wantOk {
			t.Errorf("StrMatchKmpFirst(%q,%q) = (%d,%v), want (%d,%v)",
				tt.input, tt.pattern, gotPos, gotOk, tt.wantPos, tt.wantOk)
		}
	}
}

func TestStrMatchKmpAll(t *testing.T) {
	tests := []struct {
		input   string
		pattern string
		wantPos []int
	}{
		{"hello world", "world", []int{6}},
		{"aaaaa", "aa", []int{0, 1, 2, 3}},
		{"abcabcabc", "abc", []int{0, 3, 6}},
		{"mississippi", "issi", []int{1, 4}},
		{"abcdef", "gh", []int{}},
	}

	for _, tt := range tests {
		gotPos, _ := StrMatchKmpAll([]byte(tt.input), tt.pattern)
		if !reflect.DeepEqual(gotPos, tt.wantPos) {
			t.Errorf("StrMatchKmpAll(%q,%q) = %v, want %v",
				tt.input, tt.pattern, gotPos, tt.wantPos)
		}
	}
}
