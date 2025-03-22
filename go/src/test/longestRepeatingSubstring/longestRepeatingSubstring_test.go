package longestrepeatingsubstring

import (
	"fmt"
	"testing"
)

func Test_longestRepeatingSubstring(t *testing.T) {
	s := "aaaaa"
	got := longestRepeatingSubstring(s)
	fmt.Printf("%s=%d\n", s, got)
}
