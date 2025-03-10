package m

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	v1 := ListNode{Val: 1}
	v2 := ListNode{Val: 2}
	v3 := ListNode{Val: 4}
	v1.Next = &v2
	v2.Next = &v3

	v4 := ListNode{Val: 1}
	v5 := ListNode{Val: 3}
	v6 := ListNode{Val: 4}
	v4.Next = &v5
	v5.Next = &v6
	got := mergeTwoLists(&v1, &v4)
	for ; got != nil; got = got.Next {
		fmt.Printf("%d", got.Val)
	}

}
