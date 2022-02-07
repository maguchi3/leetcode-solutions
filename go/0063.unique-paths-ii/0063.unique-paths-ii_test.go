package leetcode

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	type question struct {
		param  [][]int
		answer int
	}

	qs := []question{
		{
			param:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			answer: 2,
		},
		{
			param:  [][]int{{0, 0}},
			answer: 1,
		},
		{
			param:  [][]int{{1, 0, 0}, {0, 0, 0}, {0, 1, 0}},
			answer: 0,
		},
	}

	for _, q := range qs {
		got := uniquePathsWithObstacles(q.param)

		if got != q.answer {
			t.Errorf("expected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
