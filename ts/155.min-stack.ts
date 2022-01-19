/* * @lc app=leetcode id=155 lang=typescript *
 * [155] Min Stack
 */

// @lc code=start
function getLast<T>(arr: T[]): T {
  return arr[arr.length - 1];
}

class MinStack {
  stack: number[][];
  constructor() {
    this.stack = [];
  }

  push(val: number): void {
    if (this.stack.length === 0) {
      this.stack.push([val, val]);
      return;
    }
    const min = getLast(this.stack)[1];
    this.stack.push([val, Math.min(val, min)]);
  }

  pop(): void {
    this.stack.pop();
  }

  top(): number {
    return getLast(this.stack)[0];
  }

  getMin(): number {
    return getLast(this.stack)[1];
  }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(val)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
// @lc code=end
