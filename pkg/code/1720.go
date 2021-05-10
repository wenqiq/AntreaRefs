package code

func decode(encoded []int, first int) []int {

	res := make([]int, len(encoded)+1)

	res[0] = first

	for i, num := range encoded {
		res[i+1] = res[i] ^ num
	}

	return res

}
