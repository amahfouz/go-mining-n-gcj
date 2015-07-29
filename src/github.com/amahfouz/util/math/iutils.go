package math

// Integer math utils

func Abs(x uint8) uint8 {
	if x > 0 {
		return x
	}
	return -x
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
