/*
 * @lc app=leetcode id=373 lang=typescript
 *
 * [373] Find K Pairs with Smallest Sums
 */

import { MinPriorityQueue } from "@datastructures-js/priority-queue";

// @lc code=start
function kSmallestPairs(
  nums1: number[],
  nums2: number[],
  k: number
): number[][] {
  const minHeap = new MinPriorityQueue({
    priority: (arr: number[]) => arr[0],
  });
  const res: number[][] = [];

  for (let i = 0; i < nums1.length; i++) {
    minHeap.enqueue([nums1[i] + nums2[0], i, 0]);
  }

  while (k > 0 && !minHeap.isEmpty()) {
    const root = minHeap.dequeue() as { priority: number; element: number[] };
    const [_, i1, i2] = root.element;

    res.push([nums1[i1], nums2[i2]]);

    if (res.length === k) return res;

    if (i2 < nums2.length - 1)
      minHeap.enqueue([nums1[i1] + nums2[i2 + 1], i1, i2 + 1]);
  }

  return res;
}

// @lc code=end
