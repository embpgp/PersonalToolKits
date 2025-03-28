package ispalindrome

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	s := `race a car`
	got := isPalindrome(s)
	fmt.Printf("%+v\n", got)

}
