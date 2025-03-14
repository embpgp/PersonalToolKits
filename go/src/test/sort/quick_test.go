package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	target := []int{3, 230, 33, 23, 11, 34, 1, 0, 55, 5}

	//quickSort(target, 0, len(target)-1)
	//quickSortV1(target)
	//fmt.Printf("%+v", target)
	fmt.Printf("%d\n", findTopKByQuick(target, 7))

	quickSortV1(target)
	fmt.Printf("%+v", target)

}
