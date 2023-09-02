package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if !(a > 0) || !(b > 0) || !(c > 0) || (a+b < c || a+c < b || b+c < a) {
		return NaT
	} else if a == b && a == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else {
		return Sca
	}
}
