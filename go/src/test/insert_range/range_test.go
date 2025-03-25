package insertrange

import (
	"fmt"
	"testing"
)

func TestInsertRange(t *testing.T) {
	arr := [][]int{[]int{1, 5}, []int{10, 14}, []int{30, 50}}
	insert := []int{5, 29}

	got := InsertRange(arr, insert)
	fmt.Printf("%+v\n", got)
}
