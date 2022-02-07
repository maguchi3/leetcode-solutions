/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

package leetcode

import (
	"math"
	"strings"
)

// @lc code=start
func myAtoi(s string) int {
	s = strings.Trim(s, " ")

	if len(s) == 0 {
		return 0
	}

	var start int
	sign := 1

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	var res int

	for i := start; i < len(s); i++ {
		char := s[i]

		if !(char >= '0' && char <= '9') {
			return res * sign
		}
		res = res*10 + int(char-'0')

		if sign*res >= math.MaxInt32 {
			return math.MaxInt32
		} else if sign*res <= math.MinInt32 {
			return math.MinInt32
		}
	}

	return sign * res
}

// @lc code=end
