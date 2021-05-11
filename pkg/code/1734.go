package code

func Decode(encoded []int) []int {
	n := len(encoded)
	total := 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}

	odd := 0
	for i := 1; i < n; i += 2 {
		odd ^= encoded[i]
	}

	res := []int{total ^ odd}
	for i := 0; i < n; i++ {
		tmp := res[len(res)-1]
		res = append(res, tmp^encoded[i])
	}

	return res
}
