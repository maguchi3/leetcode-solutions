/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

package subarraySum

// @lc code=start

// Hash Map (O(n))
func subarraySum(nums []int, k int) int {
	var res, sum int
	sums := map[int]int{0: 1}

	for _, num := range nums {
		sum += num
		if count, ok := sums[sum-k]; ok {
			res += count
		}
		sums[sum]++
	}

	return res
}

// @lc code=end

// cumulative sum (O(n^2))
func subarraySum2(nums []int, k int) int {
	res := 0

	for s := 0; s < len(nums); s++ {
		sum := 0

		for e := s; e < len(nums); e++ {
			sum += nums[e]
			if sum == k {
				res++
			}
		}
	}

	return res
}
