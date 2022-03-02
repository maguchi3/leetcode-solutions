/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

package firstUniqueChar

// @lc code=start

// Slice of rune (much faster than map)
func firstUniqChar(s string) int {
	counts := make([]rune, 26)

	for _, r := range s {
		counts[r-'a']++
	}

	for i, r := range s {
		if counts[r-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

// Map of rune
func firstUniqChar2(s string) int {
	counts := make(map[rune]int)

	for _, r := range s {
		counts[r]++
	}

	for i, r := range s {
		if counts[r] == 1 {
			return i
		}
	}
	return -1
}
