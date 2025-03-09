package removek

import (
	"fmt"
	"testing"
)

func Test_removeKdigits(t *testing.T) {
	s := "1234567890"
	got := removeKdigits(s, 9)
	fmt.Printf("%s", got)
}
