package prime

import (
	"errors"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 && i != num {
			return false
		}
	}
	return true
}

func Nth(n int) (int, error) {
	count := 0
	if n <= 0 {
		return 0, errors.New("n cants be 0 or less")
	}
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
		}
		if count == n {
			return i, nil
		}
	}
}
