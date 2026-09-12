package accounting

import "math"

func checkedMulToInt64(a, b uint64) (int64, bool) {
	if a != 0 && b > uint64(math.MaxInt64)/a {
		return 0, false
	}
	n := a * b
	if n > uint64(math.MaxInt64) {
		return 0, false
	}
	return int64(n), true
}
