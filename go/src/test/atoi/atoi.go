package atoi

const (
	Maxint = (1 << 31) - 1
	Minint = -(1 << 31)
)

func MyAtoi(s string) int {
	n := len(s)

	var ans int64
	//去掉前面的''
	var i int
	for i = 0; i < n; i++ {
		if s[i] != ' ' {
			break
		}
	}
	flag := 1 //+
	if s[i] == '+' {
		flag = 1
		i++
	} else if s[i] == '-' {
		flag = -1
		i++
	}
	// remove 前面的0
	for ; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			ans = ans*10 + int64(s[i]-'0')
		} else {
			break
		}
		if ans > Maxint {
			return Maxint
		}
		if ans < Minint {
			return Minint
		}
	}
	return flag * int(ans)
}
