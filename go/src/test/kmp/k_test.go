package k

import (
	"fmt"
	"testing"
)

func TestStrstr(t *testing.T) {
	s := "sadbutsad"
	p := "sad"

	got := Strstr(s, p)
	fmt.Printf("idx:%d", got)

}
