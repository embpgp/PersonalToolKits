package reversestr

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	s := `你好啊，hello world!`
	got := reverse([]rune(s))
	fmt.Printf("%s\n", string(got))

}
