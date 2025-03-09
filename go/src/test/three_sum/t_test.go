package t

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T) {
	nums := []int{3, 0, -2, -1, 1, 2}
	got := threeSum(nums)
	fmt.Printf("%+v\n", got)
	got = threeSum1(nums)
	fmt.Printf("%+v", got)

}
