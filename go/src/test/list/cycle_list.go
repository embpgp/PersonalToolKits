package list

type node struct {
	value int
	next  *node
}

func findCycles(head *node) *node {
	if head == nil || head.next == nil {
		return nil
	}
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		if fast.next == nil {
			return nil
		}
		fast = fast.next.next
		if slow == fast { //相遇，往后偏移c段长度
			p := head
			for p != slow {
				p = p.next
				slow = slow.next
			}
			return p

		}

	}
	return nil
}
