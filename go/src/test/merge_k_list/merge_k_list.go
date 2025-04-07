package mergeklist

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (r + l) / 2
	return mergeTwo(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwo(a, b *ListNode) *ListNode {
	if a == nil || b == nil {
		if a == nil {
			return b
		} else {
			return a
		}
	}
	var head ListNode
	tail := &head
	for a != nil && b != nil {
		if a.Val < b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a == nil {
		tail.Next = b
	} else {
		tail.Next = a
	}

	return head.Next
}
