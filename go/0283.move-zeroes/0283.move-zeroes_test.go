package leetcode

import (
	"reflect"
	"testing"
)

type params struct {
	nums []int
}

func TestMoveZeroes(t *testing.T) {
	type question struct {
		params
		answer []int
	}

	qs := []question{
		{
			params: params{
				nums: []int{0, 1, 0, 3, 12},
			},
			answer: []int{1, 3, 12, 0, 0},
		},
		{
			params: params{
				nums: []int{0},
			},
			answer: []int{0},
		},
		{
			params: params{
				nums: []int{0, 0, 1},
			},
			answer: []int{1, 0, 0},
		},
	}

	for _, q := range qs {
		nums := q.nums

		moveZeroes(nums)

		if !reflect.DeepEqual(nums, q.answer) {
			t.Errorf("\nexpected: %v \ngot: %v \n%v", q.answer, nums, q)
		}
	}
}
