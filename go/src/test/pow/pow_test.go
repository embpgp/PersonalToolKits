package pow

import (
	"fmt"
	"testing"
)

func Test_myPow(t *testing.T) {
	x := 2.0
	n := 10
	got := myPow(x, n)
	got1 := myPow1(x, n)
	fmt.Printf("%f\n%f\n", got, got1)
}
