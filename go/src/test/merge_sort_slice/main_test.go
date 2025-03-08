package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Test_merge_sort_slice(t *testing.T) {
	s2 := generateRandomArray(10)
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	s1 := generateRandomArray(20)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})

	for i := 0; i < len(s2); i++ {
		s1 = append(s1, 0)
	}
	fmt.Printf("s1:%+v, s2:%+v\n", s1, s2)
	merge_sort_slice(s1, s2, len(s1)-len(s2), len(s2))
	fmt.Printf("%v\n", s1)

}

// 生产随机数组
func generateRandomArray(count int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
