package util

//Max 求两个整数的较大值
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//Min 求两个整数的较小值
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
