package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	sum := uint64(0)
	if number <= 0 || number > 64 {
		return 0, errors.New("outside of chessboard")
	}
	sum = uint64(math.Pow(2, float64(number-1)))
	return sum, nil
}

func Total() uint64 {
	return uint64((1 << 64) - 1)
}
