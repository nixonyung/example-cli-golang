package algorithms

import "math"

func PrimesUpTo(upperBound int) []int {
	// (ref.) [Sieve of Eratosthenes](https://www.wikiwand.com/en/Sieve_of_Eratosthenes)

	isPrime := make([]bool, upperBound+1)
	{
		for i := 2; i <= upperBound; i++ {
			isPrime[i] = true
		}
	}

	result := []int{}
	upperBoundSqrt := int(math.Sqrt(float64(upperBound)))
	for i := 2; i <= upperBoundSqrt; i++ {
		if isPrime[i] {
			for j := i * i; j <= upperBound; j += i {
				isPrime[j] = false
			}
			result = append(result, i)
		}
	}
	for i := upperBoundSqrt + 1; i <= upperBound; i++ {
		if isPrime[i] {
			result = append(result, i)
		}
	}
	return result
}
