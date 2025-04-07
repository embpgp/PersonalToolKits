package plusone

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 { //没有9，否则继续++
			return digits
		}
	}
	//到这里说明有进位，上面全是0了
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
