package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearchInsertion(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3, 4, 4, 4, 5, 8, 10}
	target := 4
	got := BinarySearchInsertion(nums, target)
	fmt.Printf("%d\n", got)
}
