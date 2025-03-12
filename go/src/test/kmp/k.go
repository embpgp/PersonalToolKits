package k

func getNext(p string) []int {
	n := len(p)
	next := make([]int, n+1)
	i := 0
	for j := 2; j <= n; j++ {
		for ; i > 0 && p[j-1] != p[i]; i = next[i] {
		}
		if p[j-1] == p[i] {
			i++
		}
		next[j] = i
	}
	return next
}

func Strstr(s, p string) int {
	n := len(s)
	m := len(p)
	i := 0
	next := getNext(p)
	for j := 0; j < n; j++ {
		for ; i > 0 && s[j] != p[i]; i = next[i] {

		}
		if s[j] == p[i] {
			i++
		}
		if i == m {
			return j - i + 1
		}
	}
	return -1
}
