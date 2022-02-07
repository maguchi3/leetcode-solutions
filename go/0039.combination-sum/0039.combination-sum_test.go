package leetcode

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	type params struct {
		candidates []int
		target     int
	}

	type question struct {
		params
		answer [][]int
	}

	qs := []question{
		{
			params: params{
				[]int{2, 3, 6, 7},
				7,
			},
			answer: [][]int{{2, 2, 3}, {7}},
		},
		{
			params: params{
				[]int{3, 5, 8},
				11,
			},
			answer: [][]int{{3, 3, 5}, {3, 8}},
		},
		{
			params: params{
				[]int{2},
				1,
			},
			answer: [][]int{},
		},
	}

	for _, q := range qs {
		got := combinationSum(q.candidates, q.target)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("expected: %v, got: %v \n %v", q.answer, got, q)
		}
	}
}
