package math

// IntPowN is the brute force calulcation for x^n
func IntPowN(x, n int) int {
	for i := 0; i < n; i++ {
		x = x * x
	}
	return x
}

func ModularExponentiation(a, b, n int) int {
	if b == 0 {
		return 1
	} else if b%2 == 0 {
		d := ModularExponentiation(a, b/2, n)
		return (d * d) % n
	} else {
		d := ModularExponentiation(a, b-1, n)
		return (a * d) % n
	}
}

// RepeatedSquaringRec is common repeated squaring method of finding a^b
func RepeatedSquaringRec(a, b int) int {
	if b == 0 {
		return 1
	} else if b%2 == 0 {
		d := RepeatedSquaringRec(a, b/2)
		return d * d
	} else {
		d := RepeatedSquaringRec(a, b-1)
		return a * d
	}
}

func RepeatedSquaring(a, b uint32) uint32 {
	var y uint32 = 1
	for b > 1 {
		// if b is even, a^2, and shift b >> 1
		if b&0x1 == 1 {
			b = b - 1
			y = a * y
		}
		a = a * a
		b = b >> 1
	}
	return a * y
}
