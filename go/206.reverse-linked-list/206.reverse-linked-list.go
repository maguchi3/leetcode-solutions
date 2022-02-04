/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
package leetcode

import (
	"github.com/maguroid/leetcode/go/structures"
)

type ListNode = structures.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* iteration (faster than below) */
func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}

	return prev
}

// @lc code=end

/* iteration (standard) */
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

/* recursion */
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rev := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil

	return rev
}
