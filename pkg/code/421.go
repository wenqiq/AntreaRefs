package code

func findMaximumXOR(nums []int) int {

	var res int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := nums[i] ^ nums[j]
			if tmp > res {
				res = nums[i] ^ nums[j]
			}
		}
	}

	return res

}
