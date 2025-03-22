package longestConsecutive

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	got := longestConsecutive(nums)
	fmt.Printf("%d\n", got)
}
