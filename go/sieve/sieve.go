package sieve

import "math"

func Sieve(limit int) []int {
	var primes []int
	var primesC []int
	primeCheck := 0
	countP := 1

	if limit <= 1 {
		return []int{}
	} else if limit == 2 {
		return []int{2}
	}
	for i := 2; i < limit+1; i++ {
		primes = append(primes, i)
	}
	primesC = primes
	for i := primeCheck; i <= 9; i++ {
		for j := countP; j < limit-1; j++ {
			if primes[j]%primes[i] == 0 || (10*(math.Sqrt(float64(primes[j]))) == 10*math.Floor(math.Sqrt(float64(primes[j]))) &&
				(10*(math.Sqrt(float64(primes[j]))) == 10*math.Ceil(math.Sqrt(float64(primes[j]))))) {
				primesC = []int{}
				primesC = append(primes[:j], primes[j+1:]...)
				limit--
				primes = primesC
				j--
			}
		}
		if primeCheck > limit {
			break
		}
		countP++
	}
	return primesC
}
