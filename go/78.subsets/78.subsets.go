/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

package leetcode

// @lc code=start
func subsets(nums []int) [][]int {
	subs := [][]int{{}}
	sub := []int{}

	add(0, &subs, sub, nums)

	return subs
}

func add(s int, subsets *[][]int, subset []int, nums []int) {
	for i := s; i < len(nums); i++ {
		subset = append(subset, nums[i])

		cp := make([]int, len(subset))
		copy(cp, subset)
		*subsets = append(*subsets, cp)

		add(i+1, subsets, subset, nums)
		// back track
		subset = subset[:len(subset)-1]
	}
}

// @lc code=end
