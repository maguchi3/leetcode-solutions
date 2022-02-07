package leetcode

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type question struct {
		param  []int
		answer int
	}

	qs := []question{
		{
			param:  []int{-2, 1, 3, -3, 5, 2},
			answer: 8,
		},
	}

	for _, q := range qs {
		got := maxSubArray(q.param)
		if got != q.answer {
			t.Errorf("expected: %d \n got: %d \n %v", q.answer, got, q)
		}
	}
}
