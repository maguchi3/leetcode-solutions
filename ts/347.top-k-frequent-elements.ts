/*
 * @lc app=leetcode id=347 lang=typescript
 *
 * [347] Top K Frequent Elements
 */

import { MinPriorityQueue } from "@datastructures-js/priority-queue";

// @lc code=start

// using bucket (O(N))

function topKFrequent(nums: number[], k: number) {
  const freqMap = new Map<number, number>();
  const bucket: Set<number>[] = [];
  const res: number[] = [];

  for (let num of nums) freqMap.set(num, (freqMap.get(num) || 0) + 1);

  for (let [num, freq] of freqMap) {
    bucket[freq] = (bucket[freq] || new Set()).add(num);
  }

  for (let i = bucket.length - 1; i >= 0; i--) {
    if (bucket[i]) res.push(...bucket[i]);
    if (res.length === k) break;
  }

  return res;
}

/* using PriorityQue (O(N log k))

function topKFrequent(nums: number[], k: number): number[] {
  if (k == nums.length) return nums;

  const freq = new Map<number, number>();

  for (let num of nums) {
    if (freq.has(num)) freq.set(num, freq.get(num) + 1);
    else freq.set(num, 1);
  }

  const heap = new MinPriorityQueue({ priority: (n: number) => freq.get(n) });

  for (let key of freq.keys()) {
    heap.enqueue(key);
    if (heap.size() > k) heap.dequeue();
  }

  const res = new Array<number>(k);

  for (let i = k - 1; i >= 0; i--) {
    let min = heap.dequeue() as { priority: number; element: number };
    res[i] = min.element;
  }

  return res;
}
*/
// @lc code=end
