package topKFrequent

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	nums   []int
	k      int
	answer []int
}

func TestTopKFrequent(t *testing.T) {
	qs := []question{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			k:      2,
			answer: []int{1, 2},
		},
		{
			nums:   []int{1},
			k:      1,
			answer: []int{1},
		},
	}

	for _, q := range qs {
		got := topKFrequent(q.nums, q.k)

		// This problem allows slice to be any order.
		// But in this test, judgeing as true if a slice is incleasing order because of less checking cost.

		sort.Ints(got)
		assert.Equal(t, q.answer, got)
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	params := struct {
		nums []int
		k    int
	}{
		nums: []int{1, 1, 1, 2, 2, 3},
		k:    2,
	}

	for i := 0; i < b.N; i++ {
		topKFrequent(params.nums, params.k)
	}
}

func BenchmarkTopKFrequent2(b *testing.B) {
	params := struct {
		nums []int
		k    int
	}{
		nums: []int{1, 1, 1, 2, 2, 3},
		k:    2,
	}

	for i := 0; i < b.N; i++ {
		topKFrequent2(params.nums, params.k)
	}
}
