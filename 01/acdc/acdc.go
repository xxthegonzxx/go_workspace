package acdc

// MySum takes a variadic parameter of ints and returns it's sum.
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
