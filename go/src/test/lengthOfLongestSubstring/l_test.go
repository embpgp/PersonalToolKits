package l

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "dddabcd444jkdes"
	got := lengthOfLongestSubstring(s)
	fmt.Printf("%d\n", got)
	got = lengthOfLongestSubstringv2(s)
	fmt.Printf("%d\n", got)
}
