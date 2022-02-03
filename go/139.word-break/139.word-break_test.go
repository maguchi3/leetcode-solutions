package leetcode

import "testing"

type params struct {
	s        string
	wordDict []string
}

func TestWordBreak(t *testing.T) {
	type question struct {
		params
		answer bool
	}

	qs := []question{
		{
			params: params{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			answer: true,
		},
		{
			params: params{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			answer: true,
		},
		{
			params: params{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			answer: false,
		},
	}

	for _, q := range qs {
		got := wordBreak(q.s, q.wordDict)

		if got != q.answer {
			t.Errorf("\nexpected: %t \ngot: %t \n%v", q.answer, got, q)
		}
	}
}
