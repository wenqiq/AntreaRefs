package code

import (
	"math"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {

	check := func(mid int) bool {
		res := 0
		cur := 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= mid {
				cur++
			} else {
				cur = 0
			}
			if cur/k == 1 {
				cur = 0
				res++
				if res >= m {
					return true
				}
			}
		}

		return false
	}

	l, r := math.MaxInt32, 0
	for _, b := range bloomDay {
		if b < l {
			l = b
		}
		if b > r {
			r = b
		}
	}
	maxDay := r

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l <= maxDay {
		return l
	}
	return -1
}

func minDays1(bloomDay []int, m int, k int) int {
	if k*m > len(bloomDay) {
		return -1
	}
	check := func(mid int) bool {
		res := 0
		cur := 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= mid {
				cur++
			} else {
				cur = 0
			}
			if cur/k == 1 {
				cur = 0
				res++
				if res >= m {
					return true
				}
			}
		}

		return false
	}

	r := 0
	for _, b := range bloomDay {
		if b > r {
			r = b
		}
	}

	return sort.Search(
		r, check,
	)
}
