/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
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
// with iteration
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := ListNode{Val: 0, Next: head}

	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}

	return dummy.Next
}

// @lc code=end

// with recursion
func _(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}

	return head
}
