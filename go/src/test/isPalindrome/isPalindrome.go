package ispalindrome

import "fmt"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isNumLetter(s[j]) {
			fmt.Printf("%c\n", s[j])
			j--
		}
		for i < j && !isNumLetter(s[i]) {
			fmt.Printf("%c\n", s[i])
			i++
		}
		//fmt.Printf("i:%d,s[i]:%c,j:%d,s[j]:%c\n", i, s[i], j, s[j])
		l := trans2Lower(s[i])
		r := trans2Lower(s[j])
		if l != r {
			return false
		}
		i++
		j--

	}
	return true

}

func isNumLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func trans2Lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
