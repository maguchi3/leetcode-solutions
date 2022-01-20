/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
package leetcode

import "github.com/maguroid/leetcode/go/structures"

type ListNode = structures.ListNode

// type ListNode = structures.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy, carry := &ListNode{}, 0
	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}

	return dummy.Next
}

// @lc code=end
