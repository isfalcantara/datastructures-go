package datastructures

func mod(dividend int, divisor int) int {
	return ((dividend % divisor) + divisor) % divisor
}
