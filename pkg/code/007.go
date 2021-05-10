package code

import "math"

func reverse(x int) int {

	res := 0
	for x != 0 {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		num := x % 10
		res = res*10 + num
		x /= 10
	}
	return res
}
