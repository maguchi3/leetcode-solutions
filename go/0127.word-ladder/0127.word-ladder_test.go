package leetcode

import "testing"

type params struct {
	beginWord string
	endWord   string
	wordList  []string
}

func TestLadderLength(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			answer: 5,
		},
		{
			params: params{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			answer: 0,
		},
	}

	for _, q := range qs {
		got := ladderLength(q.beginWord, q.endWord, q.wordList)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
