package plusone

import (
	"fmt"
	"testing"
)

func Test_plusOne(t *testing.T) {
	digits := []int{9, 9, 9, 1}
	got := plusOne(digits)
	fmt.Printf("%+v\n", got)
}
