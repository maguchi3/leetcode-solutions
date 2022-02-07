package leetcode

import (
	"testing"
)

type params struct {
	nums []int
}

func TestRob(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				nums: []int{2, 3, 2},
			},
			answer: 3,
		},
		{
			params: params{
				nums: []int{1, 2, 3, 1},
			},
			answer: 4,
		},
		{
			params: params{
				nums: []int{1, 2, 3},
			},
			answer: 3,
		},
	}

	for _, q := range qs {
		got := rob(q.nums)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
