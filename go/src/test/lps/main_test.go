package main

import (
	"fmt"
	"testing"
)
/*go test -timeout 30s -run ^TestLPS_dp$ src/test/lps -v*/
func TestLPS_dp(t *testing.T) {
	s := LPS{}
	s.S = "abababacccccccc"
	fmt.Printf("input:%s\nlps->dp:%s\nlps->center:%s\n", s.S, s.dp(), s.center())

}
