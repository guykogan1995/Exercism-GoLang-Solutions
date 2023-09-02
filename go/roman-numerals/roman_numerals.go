package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	s := ""
	if input <= 0 || input > 3999 {
		return "", errors.New("the Roman's did not create numerals for these numbers")
	}
	for input > 0 {
		if input >= 1000 {
			s += "M"
			input -= 1000
		} else if input >= 900 {
			s += "CM"
			input -= 900
		} else if input >= 500 {
			s += "D"
			input -= 500
		} else if input >= 450 {
			s += "LD"
			input -= 450
		} else if input >= 400 {
			s += "CD"
			input -= 400
		} else if input >= 100 {
			s += "C"
			input -= 100
		} else if input >= 90 {
			s += "XC"
			input -= 90
		} else if input >= 50 {
			s += "L"
			input -= 50
		} else if input >= 40 {
			s += "XL"
			input -= 40
		} else if input >= 10 {
			s += "X"
			input -= 10
		} else if input == 9 {
			s += "IX"
			input -= 9
		} else if input >= 5 {
			s += "V"
			input -= 5
		} else if input == 4 {
			s += "IV"
			input -= 4
		} else {
			s += "I"
			input -= 1
		}
	}
	return s, nil
}
