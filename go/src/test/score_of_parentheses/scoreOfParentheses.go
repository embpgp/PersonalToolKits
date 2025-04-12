package scoreofparentheses

import (
	"fmt"
	"strconv"
)

// 平衡的
func scoreOfParentheses(s string) int {
	st := []int{0}
	for _, c := range s {
		if c == '(' {
			st = append(st, 0)
		} else {
			v := st[len(st)-1]
			st = st[:len(st)-1]
			st[len(st)-1] += max(2*v, 1)
		}
	}
	return st[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var gDebug = false

func DbgPrintf(format string, a ...any) (n int, err error) {
	if !gDebug {
		return
	}
	return fmt.Printf(format, a...)
}

// 计算平衡子串的分数
func calculateBracketScore(s string) string {
	stack := []rune{}
	for _, c := range s {
		DbgPrintf("cccc:%v,stack:%v\n", c, stack)
		if c == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) == 0 {
				stack = append(stack, ')') //非闭合的)
				DbgPrintf("111\n")
				continue
			}

			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			DbgPrintf("vvvvv:%v, stack:%v\n", v, stack)

			if v == '(' {
				curr := 1
				for len(stack) > 0 && stack[len(stack)-1] != '(' && stack[len(stack)-1] != ')' {
					//如果还有数字
					digital := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					curr = int(digital) + int(curr)
				}
				stack = append(stack, rune(curr))

				DbgPrintf("---------stack:%v\n", stack)
			} else {
				new := v + 1
				if len(stack) > 0 {
					pre := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if pre == '(' {
						new = v * 2
					} else {
						stack = append(stack, pre)
						stack = append(stack, v)
						stack = append(stack, c)
						fmt.Printf("-----333----stack:%v\n", stack)
						continue
					}
				} else { //非闭合，直接追加数据
					stack = append(stack, v)
					stack = append(stack, c)
					DbgPrintf("-----222----stack:%v\n", stack)
					continue

				}
				DbgPrintf("new:%d\n", new)
				//stack = append(stack, rune(new))

				for len(stack) > 0 && stack[len(stack)-1] != '(' && stack[len(stack)-1] != ')' {
					//如果还有数字持续合并
					digital := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					new = digital + new
				}
				stack = append(stack, rune(new))
			}
		}

	}

	var ret string
	for _, v := range stack {
		if v == '(' || v == ')' {
			ret += string(v)
		} else {
			ret += strconv.Itoa(int(v))
		}
	}
	return ret
}
