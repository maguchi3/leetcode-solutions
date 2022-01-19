/*
 * @lc app=leetcode id=206 lang=typescript
 *
 * [206] Reverse Linked List
 */

import ListNode from "./types/list-node";

// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function reverseList(head: ListNode | null): ListNode | null {
  let prev = null;
  let curr = head;

  while (curr !== null) {
    const tmp = curr.next;
    curr.next = prev;
    prev = curr;
    curr = tmp;
  }

  return prev;
}

// function reverseList(head: ListNode | null): ListNode | null {
//   if (head === null || head.next === null) return head;

//   const prev = reverseList(head.next);

//   head.next.next = head;

//   head.next = null;

//   return prev;
// }

// @lc code=end
