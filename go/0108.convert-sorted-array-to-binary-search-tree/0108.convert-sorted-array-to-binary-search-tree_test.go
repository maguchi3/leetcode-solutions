package leetcode

import (
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestSortedArrayToBST(t *testing.T) {
	type question struct {
		param  []int
		answer *TreeNode
	}

	qs := []question{
		{
			param:  []int{-10, -3, 0, 5, 9},
			answer: structures.IntsToTree([]int{0, -3, 9, -10, structures.NULL, 5}),
		},
		{
			param:  []int{1, 3},
			answer: structures.IntsToTree([]int{3, 1}),
		},
	}

	for _, q := range qs {
		got := sortedArrayToBST(q.param)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("\nexpected: %v \ngot: %v \n%v", q.answer, got, q)
		}
	}
}
