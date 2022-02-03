package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	type question struct {
		param  []int
		answer int
	}

	qs := []question{
		{
			param:  []int{7, 1, 5, 3, 6, 4},
			answer: 5,
		},
		{
			param:  []int{7, 6, 4, 3, 1},
			answer: 0,
		},
	}

	for _, q := range qs {
		got := maxProfit(q.param)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
