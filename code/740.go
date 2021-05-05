package code

func deleteAndEarn(nums []int) int {
	maxVal := 0
	for i := range nums {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	values := make([]int, maxVal+1)

	for _, num := range nums {
		values[num] += num
	}

	do_not := values[0]
	do := values[0]
	if values[1] > do_not {
		do = values[1]
	}
	for i := 2; i < len(values); i++ {
		do_not, do = do, max(do_not+values[i], do)
	}
	return do
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
