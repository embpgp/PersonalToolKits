package sort

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	target := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Printf("%d\n", findTopKbyHeap(target, 4))
	fmt.Printf("%+v\n", target)
	heapSort(target)
	fmt.Printf("%+v", target)

}

func Test_streamMinHeap_Add(t *testing.T) {
	k := 4
	h := &streamMinHeap{
		k:    k,
		data: make([]int, 0, k+1),
	}

	target := []int{3, 230, 33, 23, 11, 34, 1, 0, 55, 5, 53, 304, 1000, 341, 231, 344}
	for _, num := range target {
		h.Add(num)
		fmt.Printf("添加 %2d -> 当前Top%d:%v\n", num, h.k, h.data)
	}

}
