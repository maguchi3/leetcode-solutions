/*
 * @lc app=leetcode id=142 lang=typescript
 *
 * [142] Linked List Cycle II
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

function detectCycle(head: ListNode | null): ListNode | null {
  const visitedNodes = new Set<ListNode>();

  while (head !== null) {
    if (visitedNodes.has(head)) return head;
    visitedNodes.add(head);
    head = head.next;
  }

  return null;
}
// @lc code=end
