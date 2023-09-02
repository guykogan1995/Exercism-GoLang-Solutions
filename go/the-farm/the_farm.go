package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	w, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 && w > 0 {
		var SillyNephewError = errors.New(fmt.Sprintf("silly nephew, there cannot be %d cows", cows))
		return 0, SillyNephewError
	}
	if err == nil && w > 0 {
		return w / float64(cows), nil
	} else if w > 0 && err == ErrScaleMalfunction {
		return (w * 2) / float64(cows), nil
	} else if w < 0 && err != ErrScaleMalfunction && err != nil {
		return 0, errors.New("non-scale error")
	} else if w < 0 {
		return 0, errors.New("negative fodder")
	} else {
		return 0, err
	}
}
