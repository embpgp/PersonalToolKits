package merge

import (
	"fmt"
	"testing"
)

func Test_mergeRange(t *testing.T) {
	arr := [][]int{}
	arr = append(arr, []int{2, 6})
	arr = append(arr, []int{30, 44})
	arr = append(arr, []int{6, 10})
	arr = append(arr, []int{16, 38})
	arr = append(arr, []int{23, 45})
	arr = append(arr, []int{46, 66})
	got := mergeRange(arr)
	fmt.Printf("merged:%+v\n", got)

}
