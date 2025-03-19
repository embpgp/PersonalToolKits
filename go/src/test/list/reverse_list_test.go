package list

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node1.next = node2
	node3 := &node{value: 3}
	node2.next = node3
	node4 := &node{value: 4}
	node3.next = node4
	node5 := &node{value: 5}
	node4.next = node5
	node6 := &node{value: 6}
	node5.next = node6
	head := &node{value: 0}
	head.next = node1
	got := reverseList(head)
	for next := got; next != nil; next = next.next {
		fmt.Printf("%d ", next.value)
	}
}
