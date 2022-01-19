/*
 * @lc app=leetcode id=141 lang=typescript
 *
 * [141] Linked List Cycle
 */

class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

// @lc code=start

function hasCycle(head: ListNode | null): boolean {
  const nodesSeen = new Set<ListNode>();

  while (head !== null) {
    if (nodesSeen.has(head)) return true;
    nodesSeen.add(head);

    head = head.next;
  }

  return false;
}
// @lc code=end
