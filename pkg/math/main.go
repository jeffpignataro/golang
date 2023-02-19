package math

func IMin(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func IMax(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func FMin(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func FMax(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
