package slowfast

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	// 测试用例1：无环链表
	t.Run("No cycle", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
		if detectCycle(head) != nil {
			t.Error("Should return nil for non-cycle list")
		}
	})

	// 测试用例2：环在尾部
	t.Run("Tail cycle", func(t *testing.T) {
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2, Next: node1}
		node1.Next = node2 // 形成环

		cycleNode := detectCycle(node1)
		if cycleNode != node2 {
			t.Errorf("Expected cycle node %v, got %v", node2, cycleNode)
		}
	})

	// 测试用例3：空链表
	t.Run("Empty list", func(t *testing.T) {
		if detectCycle(nil) != nil {
			t.Error("Should return nil for empty list")
		}
	})

	// 测试用例4：单节点无环
	t.Run("Single node no cycle", func(t *testing.T) {
		node := &ListNode{Val: 1}
		if detectCycle(node) != nil {
			t.Error("Should return nil for single node without cycle")
		}
	})

	// 测试用例5：环在中间
	t.Run("Middle cycle", func(t *testing.T) {
		tail := &ListNode{Val: 3}
		middle := &ListNode{Val: 2, Next: tail}
		head := &ListNode{Val: 1, Next: middle}
		tail.Next = middle // 形成环

		cycleNode := detectCycle(head)
		if cycleNode != middle {
			t.Errorf("Expected cycle node %v, got %v", middle, cycleNode)
		}
	})
}