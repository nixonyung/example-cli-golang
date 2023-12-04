package algorithms

func Exponential(base, power int) int {
	// (ref.) [Binary Exponentiation](https://cp-algorithms.com/algebra/binary-exp.html)

	result := 1
	for power > 0 {
		if power&1 == 1 {
			result *= base
		}
		base *= base
		power >>= 1
	}
	return result
}
