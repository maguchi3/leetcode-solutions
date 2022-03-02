package intersection

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	nums1  []int
	nums2  []int
	answer []int
}

func TestIntersection(t *testing.T) {
	qs := []question{
		{
			nums1:  []int{1, 2, 2, 1},
			nums2:  []int{2, 2},
			answer: []int{2},
		},
		{
			nums1:  []int{4, 9, 5},
			nums2:  []int{9, 4, 9, 8, 4},
			answer: []int{9, 4},
		},
	}

	for _, q := range qs {
		got := intersection(q.nums1, q.nums2)
		sort.Slice(got, func(i, j int) bool { return i > j })

		assert.Equal(t, q.answer, got)
	}
}
