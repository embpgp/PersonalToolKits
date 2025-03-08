package removeElement

import (
	"fmt"
	"testing"
)


/*go test -timeout 30s -run ^Test src/test/removedup -v*/
func TestRemoveDup(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
        fmt.Printf("before:%v\n", nums)
	fmt.Printf("%v,%d\n", nums,removeDuplicates(nums))

}


func TestRemoveDup2(t *testing.T) {
	nums := []int{0,1,1,1,2,2,2,2,3}
        fmt.Printf("before:%v\n", nums)
	fmt.Printf("%v,%d\n", nums,removeDuplicates2(nums))

}

