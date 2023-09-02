package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	s, num := "", 0
	m := map[string]string{
		"black": "0", "brown": "1", "red": "2", "orange": "3",
		"yellow": "4", "green": "5", "blue": "6", "violet": "7",
		"grey": "8", "white": "9",
	}
	for i := 0; i < 2; i++ {
		s += m[colors[i]]
	}
	num, _ = strconv.Atoi(s)
	return num
}
