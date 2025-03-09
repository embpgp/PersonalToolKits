package convert

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	str := "PAYPALISHIRING"
	got := convert(str, 3)
	fmt.Printf("%s", got)

}
