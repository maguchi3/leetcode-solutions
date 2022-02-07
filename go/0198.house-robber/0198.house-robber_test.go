package leetcode

import "testing"

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
				nums: []int{1, 2, 3, 1},
			},
			answer: 4,
		},
		{
			params: params{
				nums: []int{2, 7, 9, 3, 1},
			},
			answer: 12,
		},
	}

	for _, q := range qs {
		got1, got2 := rob(q.nums), rob2(q.nums)

		if got1 != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n %v", q.answer, got1, q)
		}
		if got2 != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n %v", q.answer, got2, q)
		}
	}
}
