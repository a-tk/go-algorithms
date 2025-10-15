package math

import (
	"container/list"
	"math"
	"math/rand/v2"
)

// PseudoPrime determines if n is likely prime or composite (false)
func PseudoPrime(n int) bool {
	if ModularExponentiation(2, n-1, n) != 1 {
		return false
	} else {
		return true
	}
}

func witness(a, n int) bool {
	// let t and u be such that t>=1, u is odd, and n-1 = 2^t * u
	// the binary representation of n-1 is same as u followed by t zeros
	// n is odd, so n-1 has at least 1 0 in the lowest order bit
	// count the >> of n-1, and this is t. The remainder is u
	u := n - 1
	t := 0
	// must execute at least once, per above
	for u&1 == 0 {
		u >>= 1
		t++
	}

	x := ModularExponentiation(a, u, n)
	for i := 0; i < t; i++ {
		xnext := x * x % n
		if xnext == 1 && x != 1 && x != n-1 {
			return true
		}
		x = xnext
	}
	if x != 1 {
		return true
	} else {
		return false
	}

}

// MillerRabin performs pseudoprimality testing with much higher certainty
//
//	n is odd number to check, and s is number of trials
//
// TODO: check n for being odd
func MillerRabin(n, s int) bool {
	for j := 0; j < s; j++ {
		a := 2 + rand.N(n-4) // interval of 2 -> n-2
		//a := 7 // test witness according to CLRS with a=7
		if witness(a, n) {
			return false
		}
	}
	return true
}

// EratosthenesSieve tests primality of n using the ancient method
func EratosthenesSieve(n int) bool {
	// seed a list with 2 (the first and only even prime)
	// enumerate each number i from 3 -> sqrt(n)
	// for any element k of the list, if k % i != 0, i is prime, add it to the list
	l := list.New()
	l.PushBack(2)
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		e := l.Front()
		for e != nil && e.Value.(int)%i != 0 {
			e = e.Next()
		}
		if e == nil {
			l.PushBack(i)
		}
	}

	e := l.Front()
	for e != nil {
		if n%e.Value.(int) == 0 {
			return false
		}
		e = e.Next()
	}
	return true
}
