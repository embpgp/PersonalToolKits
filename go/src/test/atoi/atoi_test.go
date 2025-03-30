package atoi

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	s := `  -001000`
	got := MyAtoi(s)
	new, err := strconv.Atoi(s)
	fmt.Printf("got:%d, atoi:%d,err:%v\n", got, new, err)
}
