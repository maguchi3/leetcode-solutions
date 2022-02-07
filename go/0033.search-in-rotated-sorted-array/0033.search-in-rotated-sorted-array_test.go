package leetcode

import (
	"testing"
)

type question struct {
	params
	ans int
}

type params struct {
	nums   []int
	target int
}

func TestSearch(t *testing.T) {
	qs := []question{
		{
			params{[]int{4, 5, 6, 7, 0, 1, 2}, 6},
			2,
		},
		{
			params{[]int{2, 3, 4, 5, 6, 1}, 5},
			3,
		},
		{
			params{[]int{11, 2, 3, 4, 9, 10}, 9},
			4,
		},
		{
			params{[]int{5, 1, 3}, 3},
			2,
		},
	}

	for _, q := range qs {
		got := search(q.nums, q.target)

		if got != q.ans {
			t.Errorf("search(%v, %d) = %d; want %d", q.nums, q.target, got, q.ans)
		}
	}
}
