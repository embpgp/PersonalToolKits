package pow

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)

}

func myPow1(x float64, n int) float64 {
	if n > 0 {
		return quickMul1(x, n)
	}
	return 1.0 / quickMul1(x, -n)

}

// 递归
func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

// 迭代
func quickMul1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	extra := x
	ans := 1.0
	for n > 0 {
		if n%2 == 1 {
			ans *= extra
		}
		extra *= extra
		n /= 2
	}
	return ans
}
