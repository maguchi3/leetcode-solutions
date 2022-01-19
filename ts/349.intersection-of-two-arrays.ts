/*
 * @lc app=leetcode id=349 lang=typescript
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
function intersection(nums1: number[], nums2: number[]): number[] {
  let set1 = new Set(nums1);
  let set2 = new Set(nums2);
  const res: number[] = [];

  if (set1.size > set2.size) {
    const tmp = set1;
    set1 = set2;
    set2 = tmp;
  }

  for (let n of set1) {
    if (set2.has(n)) res.push(n);
  }

  return res;
}
// @lc code=end
