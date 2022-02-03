/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
package leetcode

import "github.com/maguroid/leetcode/go/structures"

type ListNode = structures.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* using Floyd Cycle Detection Algorithm */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// intersection -> cycle start = head -> cyclestart
	intersection := getIntersection(*head)

	if intersection == nil {
		return nil
	}

	p1, p2 := head, intersection

	// If p1 equal to p2, it is the cycle-start
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

func getIntersection(head ListNode) *ListNode {
	slow, fast := &head, &head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return slow
		}
	}

	return nil
}

// @lc code=end

/* using Hash Table */
func _(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)

	for head != nil {
		if nodeMap[head] {
			return head
		}
		nodeMap[head] = true
		head = head.Next
	}

	return nil
}
