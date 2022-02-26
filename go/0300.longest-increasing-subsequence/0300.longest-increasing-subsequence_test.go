package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type params struct {
	nums []int
}

func TestLengthOfLIS(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			answer: 4,
		},
		{
			params: params{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			answer: 4,
		},
		{
			params: params{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			answer: 1,
		},
	}

	for _, q := range qs {
		got1 := lengthOfLIS(q.nums)
		got2 := lengthOfLIS2(q.nums)

		assert.Equal(t, q.answer, got1, q.params)
		assert.Equal(t, q.answer, got2, q.params)
	}
}
