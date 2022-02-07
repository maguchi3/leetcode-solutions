/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}

	prev := dummy

	for head != nil {
		if isDuplicated(head) {
			for isDuplicated(head) {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}

		head = head.Next
	}

	return dummy.Next
}

func isDuplicated(head *ListNode) bool {
	return head.Next != nil && head.Val == head.Next.Val
}

// @lc code=end
