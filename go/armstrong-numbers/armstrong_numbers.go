package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	sum := 0.0
	num := strconv.Itoa(n)
	for i := 0; i < len(num); i++ {
		numPiece, _ := strconv.Atoi(string(num[i]))
		sum += math.Pow(float64(numPiece), float64(len(num)))
	}
	if sum == float64(n) {
		return true
	} else {
		return false
	}
}
