package utils

// AbsValue calculates the absolute value
func AbsValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
