package scoreofparentheses

import (
	"fmt"
	"testing"
)

func Test_scoreOfParentheses(t *testing.T) {
	//s := `()())))()()()((`
	s := `)(())(`
	got := calculateBracketScore(s)
	fmt.Printf("%+v\n", got)

}
