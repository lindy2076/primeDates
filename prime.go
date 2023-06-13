package main

import "math"

func IsPrime(num uint32) bool {
	if num < 2 {
		return false
	}
	numSqrt := math.Sqrt(float64(num))
	numSqrtInt := uint32(math.Floor(numSqrt))
	var i uint32 = 2
	for ; i <= numSqrtInt; i++ {
		if num%i == 0 && num != i {
			return false
		}
	}

	return true
}
