package removek

import (
	"fmt"
	"testing"
)

func Test_removeKdigits(t *testing.T) {
	s := "0001234567890"
	got := removeKdigits1(s, 9)
	fmt.Printf("%s", got)
}
