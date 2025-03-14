package sort

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	target := []int{3, 230, 33, 23, 11, 34, 1, 0, 55, 5}
	fmt.Printf("%d\n", findTopKbyHeap(target, 3))
	fmt.Printf("%+v\n", target)
	heapSort(target)
	fmt.Printf("%+v", target)

}
