package leetcode

import (
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

type params struct {
	head *ListNode
}

func TestHasCycle(t *testing.T) {
	type question struct {
		params
		pos    int
		answer bool
	}

	qs := []question{
		{
			params: params{
				head: structures.IntsToList([]int{3, 2, 0, -4}),
			},
			pos:    1,
			answer: true,
		},
		{
			params: params{
				head: structures.IntsToList([]int{1, 2}),
			},
			pos:    0,
			answer: true,
		},
		{
			params: params{
				head: structures.IntsToList([]int{1}),
			},
			pos:    -1,
			answer: false,
		},
	}

	for _, q := range qs {
		// create cycle in head
		head := structures.CreateCycle(q.head, q.pos)
		got := hasCycle(head)

		if got != q.answer {
			t.Errorf("\nexpected: %t \ngot: %t \n%v", q.answer, got, q)
		}
	}

}
