package l

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flower", "flower", "flower"}
	got := longestCommonPrefix(strs)
	fmt.Printf("%v\n", got)
	got = longestCommonPrefixV2(strs)
	fmt.Printf("%v\n", got)
}
