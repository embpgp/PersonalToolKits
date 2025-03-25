package binarytree

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	root := TreeNode{Val: 0}
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	root.Left = &node1
	root.Right = &node2
	node1.Left = &node3
	node2.Right = &node4
	in := inorderTraversal(&root)
	in1 := inorderTraversal1(&root)
	fmt.Printf("%+v\n%+v\n", in, in1)

	before := beforeorderTraversal(&root)
	before1 := beforeorderTraversal1(&root)
	fmt.Printf("%+v\n%+v\n", before, before1)

	after := afterorderTraversal(&root)
	after1 := afterorderTraversal1(&root)
	fmt.Printf("%+v\n%+v\n", after, after1)

}
