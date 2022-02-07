package leetcode

import (
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

type params struct {
	head *ListNode
}

func TestDetectCycle(t *testing.T) {
	type question struct {
		params
		pos    int
		answer *ListNode
	}

	qs := []question{
		{
			params: params{
				head: structures.IntsToList([]int{3, 2, 0, -4}),
			},
			pos: 1,
		},
		{
			params: params{
				head: structures.IntsToList([]int{1, 2}),
			},
			pos: 0,
		},
		{
			params: params{
				head: structures.IntsToList([]int{1}),
			},
			pos:    -1,
			answer: nil,
		},
	}

	for _, q := range qs {
		head := structures.CreateCycle(q.head, q.pos)
		answer := head
		if q.pos != -1 {
			for i := 0; i < q.pos; i++ {
				answer = answer.Next
			}
			q.answer = answer
		}

		got := detectCycle(head)

		if got != q.answer {
			t.Errorf("\nexpected: %v \ngot: %v \n %v", q.answer, got, q)
		}
	}
}
