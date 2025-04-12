package k

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrstr(t *testing.T) {
	s := "sadbutsad"
	p := "dbu"

	got := Strstr(s, p)
	idx := strings.Index(s, p)
	fmt.Printf("idx:%d,%d", got, idx)

}
