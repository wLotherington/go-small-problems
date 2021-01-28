// Easy
// https://leetcode.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func setNode(node *ListNode) (*ListNode, *ListNode) {
	next := node.Next
	node.Next = nil
	return node, next
}

func setNodeAndNext(node *ListNode) (*ListNode, *ListNode, *ListNode) {
	tmp1, tmp2 := setNode(node)
	return node, tmp1, tmp2
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var first, last *ListNode

	// Check for empty lists
	if l1 == nil && l2 == nil {
		return nil
	} else if l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	}

	// Set First
	if l1.Val <= l2.Val {
		first, l1 = setNode(l1)
	} else {
		first, l2 = setNode(l2)
	}

	// Set Last
	fmt.Println(l1, l2, first, last)

	if l2 == nil {
		first.Next, last, l1 = setNodeAndNext(l1)
	} else if l1 == nil {
		first.Next, last, l2 = setNodeAndNext(l2)
	} else if l1.Val <= l2.Val {
		first.Next, last, l1 = setNodeAndNext(l1)
	} else {
		first.Next, last, l2 = setNodeAndNext(l2)
	}

	// Merge Lists
	for l1 != nil || l2 != nil {
		if l2 == nil {
			last.Next, last, l1 = setNodeAndNext(l1)
		} else if l1 == nil {
			last.Next, last, l2 = setNodeAndNext(l2)
		} else if l1.Val <= l2.Val {
			last.Next, last, l1 = setNodeAndNext(l1)
		} else if l2.Val <= l1.Val {
			last.Next, last, l2 = setNodeAndNext(l2)
		}
	}

	return first
}