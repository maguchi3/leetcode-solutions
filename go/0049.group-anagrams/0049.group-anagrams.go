/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

package leetcode

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	anagrams := map[[26]int][]string{}
	res := [][]string{}

	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		anagrams[k] = append(anagrams[k], s)
	}

	for _, v := range anagrams {
		res = append(res, v)
	}

	return res
}

// @lc code=end
