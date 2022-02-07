/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
package leetcode

// @lc code=start
func nextPermutation(nums []int) {
	n := len(nums)

	if n == 1 {
		return
	}

	l, r := n-2, n-1

	for l >= 0 && nums[l] >= nums[l+1] {
		l--
	}

	if l >= 0 {
		for nums[r] <= nums[l] {
			r--
		}
		swap(&nums, l, r)
	}

	reverse(&nums, l+1)
}

func swap(nums *[]int, i int, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func reverse(nums *[]int, start int) {
	for i, j := start, len(*nums)-1; i < j; i, j = i+1, j-1 {
		swap(nums, i, j)
	}
}

// @lc code=end
