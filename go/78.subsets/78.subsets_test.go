package leetcode

import (
	"reflect"
	"testing"
)

func TestSubset(t *testing.T) {
	type question struct {
		params []int
		answer [][]int
	}

	qs := []question{
		{
			params: []int{1, 2, 3},
			answer: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, q := range qs {
		got := subsets(q.params)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("expected: %v \ngot: %v \n%v", q.answer, got, q)
		}
	}
}
