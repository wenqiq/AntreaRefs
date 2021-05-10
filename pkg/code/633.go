package code

import "math"

func JudgeSquareSum(c int) bool {
	sqrt := int(math.Sqrt(float64(c)))
	for i := 0; i < sqrt; i ++{
		b := float64(c - i * i)
		a := math.Sqrt(b)
		if a * a == b{
			return true
		}
	}
	return false
}



