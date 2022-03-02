package kSmallestPairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	nums1  []int
	nums2  []int
	k      int
	answer [][]int
}

func TestKSmallestPairs(t *testing.T) {
	qs := []question{
		{
			nums1:  []int{1, 7, 11},
			nums2:  []int{2, 4, 6},
			k:      3,
			answer: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			nums1:  []int{1, 1, 2},
			nums2:  []int{1, 2, 3},
			k:      2,
			answer: [][]int{{1, 1}, {1, 1}},
		},
		{
			nums1:  []int{1, 2},
			nums2:  []int{3},
			k:      3,
			answer: [][]int{{1, 3}, {2, 3}},
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.answer, kSmallestPairs(q.nums1, q.nums2, q.k))
	}

}
