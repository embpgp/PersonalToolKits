package main

import (
	"fmt"
	"testing"
)

func TestLPS_dp(t *testing.T) {
	s := LPS{}
	s.S = "abababacccccccc"
	fmt.Printf("input:%s\nlps->dp:%s\nlps->center:%s\n", s.S, s.dp(), s.center())

}
