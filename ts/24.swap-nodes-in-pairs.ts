/*
 * @lc app=leetcode id=24 lang=typescript
 *
 * [24] Swap Nodes in Pairs
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

function swapPairs(head: ListNode | null): ListNode | null {
  if (head === null || head.next === null) return head;

  let firstNode = head;
  let secondNode = head.next;

  firstNode.next = swapPairs(secondNode.next);
  secondNode.next = firstNode;

  return secondNode;
}
// @lc code=end
