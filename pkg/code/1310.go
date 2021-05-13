package code

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)

	for i, a := range arr {
		xors[i+1] = xors[i] ^ a
	}

	res := make([]int, len(queries))

	for i, vals := range queries {
		res[i] = xors[vals[0]] ^ xors[vals[1]+1]
	}

	return res

}
