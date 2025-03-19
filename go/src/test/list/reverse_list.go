package list

func reverseList(head *node) *node {
	curr := head
	var prev *node
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
