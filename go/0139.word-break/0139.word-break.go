/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */
package leetcode

// @lc code=start

// Dynamic Programming
func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && includes(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func includes(s string, dict []string) bool {
	for _, v := range dict {
		if v == s {
			return true
		}
	}

	return false
}

// @lc code=end

// Recursion with Memorization
func _(s string, wordDict []string) bool {
	memo := make(map[int]bool)
	wordMap := make(map[string]int)

	for i, word := range wordDict {
		wordMap[word] = i
	}

	var backTrack func(start int) bool
	backTrack = func(start int) bool {
		if start == len(s) {
			return true
		}
		if _, ok := memo[start]; ok {
			return memo[start]
		}
		for end := start + 1; end <= len(s); end++ {
			if _, ok := wordMap[s[start:end]]; ok && backTrack(end) {
				memo[end] = true
				return true
			}
		}
		memo[start] = false
		return false
	}

	return backTrack(0)
}
