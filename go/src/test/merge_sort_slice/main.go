package main

func main() {

}

func merge_sort_slice(s1, s2 []int, c1, c2 int) {
	for p1, p2, tail := c1-1, c2-1, c1+c2-1; p1 >= 0 || p2 >= 0; tail-- {
		if p1 < 0 {
			s1[tail] = s2[p2]
			p2--
		} else if p2 < 0 {
			s1[tail] = s1[p1]
			p1--
		} else if s1[p1] > s2[p2] {
			s1[tail] = s1[p1]
			p1--
		} else {
			s1[tail] = s2[p2]
			p2--
		}
	}
}
