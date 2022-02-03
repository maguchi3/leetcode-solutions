/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

package leetcode

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// word map to validate candidate
	wordMap := getWordMap(wordList, beginWord)

	// que for BFS
	que := []string{beginWord}
	depth := 0

	for len(que) > 0 {
		depth++

		n := len(que)
		for i := 0; i < n; i++ {
			word := que[0]
			que = que[1:]

			candidates := getCandidates(word)

			for _, c := range candidates {
				if _, ok := wordMap[c]; ok {
					if c == endWord {
						return depth + 1
					}

					delete(wordMap, c)
					que = append(que, c)
				}
			}
		}
	}

	return 0
}

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)

	for i, word := range wordList {
		if _, ok := wordMap[word]; !ok {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}

	return wordMap
}

func getCandidates(word string) []string {
	res := []string{}
	for i := 0; i < 26; i++ {
		for j := range word {
			new := replaceAt(word, j, rune('a'+i))
			if new != word {
				res = append(res, new)
			}
		}
	}
	return res
}
func replaceAt(target string, index int, new rune) string {
	runes := []rune(target)

	runes[index] = new

	return string(runes)
}

// @lc code=end
