package leetcode

import "testing"

type params struct {
	nums []int
}

func TestFindMin(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				nums: []int{3, 4, 5, 1, 2},
			},
			answer: 1,
		},
		{
			params: params{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			answer: 0,
		},
		{
			params: params{
				nums: []int{11, 13, 15, 17},
			},
			answer: 11,
		},
	}

	for _, q := range qs {
		got := findMin(q.nums)

		if got != q.answer {
			t.Errorf("\nezpected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
