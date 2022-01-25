/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
package leetcode

// @lc code=start
func permute(nums []int) [][]int {
	ans := make([][]int, 0)

	backTrack(nums, &ans, 0)

	return ans
}

func backTrack(nums []int, ans *[][]int, left int) {
	if left == len(nums) {
		cpy := make([]int, len(nums))
		copy(cpy, nums)

		*ans = append(*ans, cpy)
		return
	}

	for right := left; right < len(nums); right++ {
		swap(&nums, left, right)
		backTrack(nums, ans, left+1)
		swap(&nums, right, left)
	}
}

func swap(nums *[]int, x int, y int) {
	(*nums)[x], (*nums)[y] = (*nums)[y], (*nums)[x]
}

// @lc code=end
