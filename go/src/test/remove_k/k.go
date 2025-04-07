package removek

import (
	"fmt"
	"strings"
)

// 贪心算法转换成移除元素满足第一个 n[i-1] > n[i]的时候，移除掉n[i-1]，
func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	for i := 0; i < len(num); i++ {
		digital := num[i]
		//fmt.Printf("digit:%d\n", digital)
		for k > 0 && len(stack) > 0 && digital < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		fmt.Printf("%s,k:%d\n", stack, k)
		stack = append(stack, digital)
	}
	stack = stack[:len(stack)-k]
	s := string(stack)
	s = strings.TrimLeft(s, "0")
	if len(s) == 0 {
		return "0"
	}
	return s
}

func removeKdigits1(num string, k int) string {
	//这个题转换为，当n[k-1]>n[k]的时候，需要移除掉n[k-1]
	n := len(num)
	if k >= n {
		return "0"
	}
	var stack []byte
	for i := 0; i < n; i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			//remove
			stack = stack[:len(stack)-1]
			k--
		}
		fmt.Printf("%s,k:%d\n", stack, k)
		stack = append(stack, num[i])

	}
	//若k>0
	stack = stack[:len(stack)-k]
	if len(stack) == 0 {
		return "0"
	}
	for i, v := range stack {
		if v == '0' {
			continue
		} else {
			//直接返回所有的
			return string(stack[i:])
		}
	}
	return "0"
}
