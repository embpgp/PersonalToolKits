package removenthfromend

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func isValid(s string) bool {
	n := len(s)
	arr := make([]byte, 0)
	for i := 0; i < n; i++ {
		if len(arr) == 0 {
			arr = append(arr, s[i])
			continue
		}
		if (arr[len(arr)-1] == '(' && s[i] == ')') ||
			(arr[len(arr)-1] == '[' && s[i] == ']') ||
			(arr[len(arr)-1] == '{' && s[i] == '}') {
			arr = arr[:len(arr)-1]
		} else {
			arr = append(arr, s[i])
		}

	}
	return len(arr) == 0
}
