package leetcode

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	type question struct {
		param  []int
		answer [][]int
	}

	qs := []question{
		{
			param:  []int{1, 2, 3},
			answer: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
		},
	}

	for _, q := range qs {
		got := permute(q.param)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("expected: %v \ngot: %v \n %v", q.answer, got, q)
		}
	}
}
