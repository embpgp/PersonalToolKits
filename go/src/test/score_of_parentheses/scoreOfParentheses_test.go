package scoreofparentheses

import (
	"fmt"
	"testing"
)

func Test_scoreOfParentheses(t *testing.T) {
	s := `()())))()()()((`
	got := calculateBracketScore(s)
	fmt.Printf("%+v\n", got)

}
