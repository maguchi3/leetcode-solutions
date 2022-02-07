package leetcode

import (
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

type params struct {
	root      *TreeNode
	targetSum int
}

func TestHasPathSum(t *testing.T) {
	type question struct {
		params
		answer bool
	}

	qs := []question{
		{
			params: params{
				root:      structures.IntsToTree([]int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, structures.NULL, 1}),
				targetSum: 22,
			},
			answer: true,
		},
	}

	for _, q := range qs {
		got := hasPathSum(q.root, q.targetSum)

		if got != q.answer {
			t.Errorf("\nexpected: %t \ngot: %t \n%v", q.answer, got, q)
		}

	}
}
