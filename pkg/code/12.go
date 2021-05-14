package code

func intToRoman(num int) string {
	//         hash_map = (
	//     ("M", 1000),("CM", 900),  ("D", 500),("CD", 400), ("C", 100), ("XC", 90), ("L", 50), ("XL", 40),
	//     ("X", 10), ("IX", 9), ("V", 5),("IV", 4), ("I", 1)
	// )

	vals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	res := []byte{}
	for i, key := range keys {
		for num >= key {
			num -= key
			res = append(res, vals[i]...)
		}
		if num == 0 {
			break
		}
	}
	return string(res)

}
