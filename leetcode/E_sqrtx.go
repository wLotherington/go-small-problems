// Easy
// https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	low := 0
	high := x
	mid := (low + high) / 2

	for low < mid && high > mid {
		candidate := mid * mid

		if candidate == x {
			return mid
		} else if candidate > x {
			high = mid
		} else if candidate < x {
			low = mid
		}

		mid = (low + high) / 2
	}

	return low
}