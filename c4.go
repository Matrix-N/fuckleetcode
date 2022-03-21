package fuckleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast, isCycle := hasCycle(head)
	if !isCycle {
		return nil
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

func hasCycle(head *ListNode) (*ListNode, bool) {
	if head == nil {
		return nil, false
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return slow, true
		}
	}

	return nil, false
}
