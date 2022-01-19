/*
 * @lc app=leetcode id=560 lang=typescript
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start

// with Map - O(n)
function subarraySum(nums: number[], k: number): number {
  let [count, sum] = [0, 0];
  const sumOccurrences = new Map<number, number>();

  sumOccurrences.set(0, 1);

  for (let num of nums) {
    sum += num;
    if (sumOccurrences.has(sum - k)) count += sumOccurrences.get(sum - k);
    sumOccurrences.set(sum, (sumOccurrences.get(sum) || 0) + 1);
  }
  return count;
}

/* with iteration - O(n^2)
function subarraySum(nums: number[], k: number): number {
  let count = 0;
  for (let s = 0; s < nums.length; s++) {
    let sum = 0;
    for (let e = s; e < nums.length; e++) {
      sum += nums[e];
      if (sum === k) count++;
    }
  }
  return count;
}
*/
// @lc code=end
