package primes

import (
	"fmt"
	"math"
)

// Checks if num is prime.
func IsPrime(num uint32) bool {
	if num < 2 {
		return false
	}

	numSqrt := math.Sqrt(float64(num))
	numSqrtInt := uint32(math.Floor(numSqrt))
	var div uint32 = 2
	for ; div <= numSqrtInt; div++ {
		if num%div == 0 && num != div {
			return false
		}
	}

	return true
}

// Returns a string in format "*num* is [not] prime"
func PrimalityString(num uint32) string {
	if IsPrime(num) {
		return fmt.Sprintf("%d is prime", num)
	}
	return fmt.Sprintf("%d is not prime", num)
}
