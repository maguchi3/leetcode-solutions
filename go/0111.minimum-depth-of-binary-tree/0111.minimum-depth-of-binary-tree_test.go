package leetcode

import (
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestMinDepth(t *testing.T) {
	type question struct {
		param  *TreeNode
		answer int
	}

	qs := []question{
		{
			param:  structures.IntsToTree([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}),
			answer: 2,
		},
		{
			param:  structures.IntsToTree([]int{2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6}),
			answer: 5,
		},
		{
			param:  structures.IntsToTree([]int{1, 2}),
			answer: 2,
		},
	}

	for _, q := range qs {
		got := minDepth(q.param)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
