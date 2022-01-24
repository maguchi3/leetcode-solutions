package leetcode

import "testing"

type Question struct {
	params Params
	answer Answer
}

type Params struct {
	nums   []int
	target int
}

type Answer struct {
	output int
}

func TestSearchInsert(t *testing.T) {
	qs := []Question{
		{
			Params{
				[]int{1, 3, 5, 6},
				5,
			},
			Answer{2},
		},
		{
			Params{
				[]int{1, 3, 5, 6},
				2,
			},
			Answer{1},
		},
		{
			Params{
				[]int{1, 3, 5, 6},
				7,
			},
			Answer{4},
		},
		{
			Params{
				[]int{1, 3, 5, 6},
				0,
			},
			Answer{0},
		},
	}

	for _, q := range qs {
		got := searchInsert(q.params.nums, q.params.target)

		if got != q.answer.output {
			t.Errorf("Expected %d, Got %d", q.answer.output, got)
		}
	}
}
