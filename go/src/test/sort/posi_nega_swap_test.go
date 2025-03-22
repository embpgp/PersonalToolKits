package sort

import (
	"fmt"
	"testing"
)

func TestPosiNegaSwap(t *testing.T) {
	arr := []int{1, -2, 3, 4, -3, -1, -20, 30, 50, 43, 10, -30, -100}
	fmt.Printf("%+v\n", arr)
	PosiNegaSwap(arr)

}
