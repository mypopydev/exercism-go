package triangle

import "math"

type Kind int

const (
	Equ = iota // equilateral
	Iso        // isosceles
	Sca        // scalene
	NaT        // not a triangle
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a <= float64(0) || b <= float64(0) || c <= float64(0) ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
		math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	} else if a+b < c || a+c < b || b+c < a {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}
