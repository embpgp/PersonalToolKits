package binarytreetolinkedlist

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

// O(1)空间是把左子树放在右子树，右子树放在左子树的最右边
func flattenO1(root *TreeNode) {
	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			//遍历左子树的右子树
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			root = root.Right

		}
	}

}
