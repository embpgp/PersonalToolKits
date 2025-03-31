package sort

import (
	"fmt"
	"testing"
)

func Test_count(t *testing.T) {
	arr := []int{3, 40, 3, 38, 23, 44, 50, 11}
	//got := count(arr)
	got := selectSort(arr)
	fmt.Printf("%+v\n", got)
}
