/*
 * @lc app=leetcode id=703 lang=typescript
 *
 * [703] Kth Largest Element in a Stream
 */

import { MinPriorityQueue } from "@datastructures-js/priority-queue";

// @lc code=start
/* with heap */
class KthLargest {
  k: number;
  heap: InstanceType<typeof MinPriorityQueue>;
  constructor(k: number, nums: number[]) {
    this.k = k;
    this.heap = new MinPriorityQueue({ priority: (el: number) => el });
    nums.forEach((n) => this.add(n));
  }

  add(val: number): number {
    if (this.heap.size() < this.k) this.heap.enqueue(val);
    else if (val > this.minVal) {
      this.heap.enqueue(val);
      this.heap.dequeue();
    }

    return this.minVal;
  }

  get minVal(): number {
    let minNode = this.heap.front() as { priority: number; element: number };

    return minNode.element;
  }
}

/* without heap

class KthLargest {
  k: number;
  stream: number[];
  constructor(k: number, nums: number[]) {
    this.k = k;
    this.stream = nums.sort((a, b) => a - b);
  }

  add(val: number): number {
    this.stream.splice(this.posOf(val), 0, val);

    return this.stream[this.size - this.k];
  }

  get size() {
    return this.stream.length;
  }

  private posOf(val: number): number {
    let left = 0;
    let right = this.size - 1;

    while (left <= right) {
      let mid = ~~((left + right) / 2);
      if (val === this.stream[mid]) return mid;
      else if (val > this.stream[mid]) left = mid + 1;
      else right = mid - 1;
    }

    return left;
  }
}
*/

/**
 * Your KthLargest object will be instantiated and called as such:
 * var obj = new KthLargest(k, nums)
 * var param_1 = obj.add(val)
 */
// @lc code=end
