package xx2

import "gonum.org/v1/gonum/floats/scalar"

func Round(i float64, digits int) float64 {
	return scalar.Round(i, digits)
}

func Round0(i float64) float64 {
	return Round(i, 0)
}

func Round1(i float64) float64 {
	return Round(i, 1)
}

func Round2(i float64) float64 {
	return Round(i, 2)
}

func Round3(i float64) float64 {
	return Round(i, 3)
}

func Round4(i float64) float64 {
	return Round(i, 4)
}

func Round5(i float64) float64 {
	return Round(i, 5)
}

func Round6(i float64) float64 {
	return Round(i, 6)
}
