package math

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkPseudoPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PseudoPrime(rand.Int())
	}
}

func BenchmarkMillerRabin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MillerRabin(rand.Int(), 3)
	}
}

func BenchmarkEratosthenesSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EratosthenesSieve(rand.Int())
	}
}

func TestPseudoPrime_Simple(t *testing.T) {

	got := PseudoPrime(31)
	if !got {
		t.Errorf("Fail, found 31 composite")
	}

	got = PseudoPrime(63)
	if got {
		t.Errorf("Fail, found 63 prime")
	}

	got = PseudoPrime(341) // actually composite
	if !got {
		t.Errorf("Fail, found 341 composite")
	}

	got = PseudoPrime(561) // carmichael number
	if !got {
		t.Errorf("Fail, found 561 composite")
	}
}

func TestMillerRabin(t *testing.T) {

	got := MillerRabin(31, 1)
	if !got {
		t.Errorf("Fail, found 31 composite")
	}

	got = MillerRabin(63, 1)
	if got {
		t.Errorf("Fail, found 63 prime")
	}

	got = MillerRabin(341, 1) // actually composite
	if got {
		t.Errorf("Fail, found 341 prime")
	}

	got = MillerRabin(561, 1) // carmichael number
	if got {
		t.Errorf("Fail, found 561 prime")
	}

	got = MillerRabin(1001, 1) // composite
	if got {
		t.Errorf("Fail, found 1001 prime")
	}
}

func TestEratosthenesSieve(t *testing.T) {

	got := EratosthenesSieve(31)
	if !got {
		t.Errorf("Fail, found 31 composite")
	}

	got = EratosthenesSieve(63)
	if got {
		t.Errorf("Fail, found 63 prime")
	}

	got = EratosthenesSieve(341)
	if got {
		t.Errorf("Fail, found 341 prime")
	}

	got = EratosthenesSieve(561)
	if got {
		t.Errorf("Fail, found 561 prime")
	}
}
