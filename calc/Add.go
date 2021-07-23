package calc

func Add(args ...int) int {
	res := 0
	for _, val := range args {
		res += val
	}
	return res
}
