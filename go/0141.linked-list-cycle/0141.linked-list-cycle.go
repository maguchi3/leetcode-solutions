/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
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

/* using Floyd's Circle Finding Algorithm - Space Complexity O(1) */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// @lc code=end

/* using Hash Table - Space Complexity O(n) */
func _(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool)

	for head != nil {
		if nodeMap[head] {
			return true
		}
		nodeMap[head] = true
		head = head.Next
	}

	return false
}
