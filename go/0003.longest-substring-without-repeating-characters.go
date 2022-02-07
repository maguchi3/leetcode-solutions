/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var max, start int
	charIdxMap := make(map[rune]int)

	for end, char := range s {
		if _, exist := charIdxMap[char]; exist && charIdxMap[char] >= start {
			if end-start > max {
				max = end - start
			}
			start = charIdxMap[char] + 1
		}
		charIdxMap[char] = end
	}

	if len(s)-start > max {
		max = len(s) - start
	}

	return max
}

// @lc code=end
