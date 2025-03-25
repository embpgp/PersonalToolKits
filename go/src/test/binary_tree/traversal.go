package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序 left print right
func inorderTraversal(root *TreeNode) []int {
	//递归
	var res []int
	var inorderFunc func(node *TreeNode)
	inorderFunc = func(node *TreeNode) {

		if node == nil {
			return
		}
		inorderFunc(node.Left)
		res = append(res, node.Val)
		inorderFunc(node.Right)
	}
	inorderFunc(root)
	return res
}

// 前序 print left  right
func beforeorderTraversal(root *TreeNode) []int {
	//递归
	var res []int
	var inorderFunc func(node *TreeNode)
	inorderFunc = func(node *TreeNode) {

		if node == nil {
			return
		}
		res = append(res, node.Val)
		inorderFunc(node.Left)
		inorderFunc(node.Right)
	}
	inorderFunc(root)
	return res
}

// 后序  left  right print
func afterorderTraversal(root *TreeNode) []int {
	//递归
	var res []int
	var inorderFunc func(node *TreeNode)
	inorderFunc = func(node *TreeNode) {

		if node == nil {
			return
		}

		inorderFunc(node.Left)
		inorderFunc(node.Right)
		res = append(res, node.Val)
	}
	inorderFunc(root)
	return res
}

// 中序 left print right
func inorderTraversal1(root *TreeNode) []int {
	//迭代
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 压栈
			stack = append(stack, root)
			root = root.Left
		}

		//出栈
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, tmp.Val)
		root = tmp.Right

	}

	return res
}

// 前序 print left  right
func beforeorderTraversal1(root *TreeNode) []int {
	//迭代
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			//压栈
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		//出栈
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = tmp.Right

	}

	return res
}

// 后序 print left  right
func afterorderTraversal1(root *TreeNode) []int {
	// 迭代
	var res []int
	var stack []*TreeNode
	last := root
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//出栈
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp.Right == nil || tmp.Right == last {
			res = append(res, tmp.Val)
			last = tmp
			tmp = nil
		} else {
			stack = append(stack, tmp)
			root = tmp.Right
		}

	}

	return res
}
