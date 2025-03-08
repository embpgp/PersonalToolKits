package removeElement

import (
	"fmt"
	"testing"
)


/*go test -timeout 30s -run ^Test src/test/removeElement -v*/
func TestRemoveElement(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
        fmt.Printf("before:%v\n", nums)
	fmt.Printf("%v,%d\n", nums,removeElement(nums,2))

}



