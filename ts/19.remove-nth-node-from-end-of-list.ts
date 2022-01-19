/*
 * @lc app=leetcode id=19 lang=typescript
 *
 * [19] Remove Nth Node From End of List
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

function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
  let dummy = new ListNode(0);
  dummy.next = head;
  let len = 0;

  let first = head;

  while (first !== null) {
    first = first.next;
    len++;
  }

  len -= n;

  first = dummy;

  while (len > 0) {
    first = first.next;
    len--;
  }

  first.next = first.next.next;
  return dummy.next;
}
// @lc code=end
