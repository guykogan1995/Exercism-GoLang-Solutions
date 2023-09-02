package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	maximum, oldMax := 1, 0
	if len(digits) < span || (digits == "" && span != 0) || span < 0 {
		return 0, errors.New("span must be smaller than string length")
	}
	for i := 0; i <= len(digits)-span; i++ {
		maximum = 1
		series := digits[i : i+span]
		for _, value := range series {
			num, err := strconv.Atoi(string(value))
			if err != nil {
				return 0, errors.New("could not convert to number")
			} else {
				maximum *= num
			}
		}
		if maximum > oldMax {
			oldMax = maximum
		}
	}
	return int64(oldMax), nil
}
