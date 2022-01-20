/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)

	for curIdx, curNum := range nums {
		if reqIdx, isPresent := idxMap[target-curNum]; isPresent {
			return []int{reqIdx, curIdx}
		}
		idxMap[curNum] = curIdx
	}

	return []int{}
}

// @lc code=end
