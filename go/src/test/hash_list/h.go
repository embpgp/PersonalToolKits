package hashlist

type hlistNode struct {
	next  *hlistNode
	pprev **hlistNode //&prev.next
}

type hlistHead struct {
	first *hlistNode
}

func hlistAddHead(n *hlistNode, h *hlistHead) {
	first := h.first
	n.next = first
	if first != nil {
		first.pprev = &n.next
	}
	n.pprev = &h.first
	h.first = n
}

func hlistInsertAfter(list, new *hlistNode) {

	next := list.next
	new.next = next
	list.next = new
	new.pprev = &list.next
	if next != nil {
		next.pprev = &new.next
	}

}

func hlistDel(n *hlistNode) {
	next := n.next
	pprev := n.pprev
	if pprev != nil {
		*pprev = next
	}
	if next != nil {
		next.pprev = pprev
	}
	n.next = nil
	n.pprev = nil
}
