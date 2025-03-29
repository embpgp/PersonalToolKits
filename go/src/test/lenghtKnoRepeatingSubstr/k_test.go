package lenghtknorepeatingsubstr

import (
	"fmt"
	"testing"
)

func Test_numKLenSubstrNoRepeats(t *testing.T) {
	s := `havefunonleetcode`
	k := 5
	got := numKLenSubstrNoRepeats(s, k)
	fmt.Printf("%v\n", got)
}
