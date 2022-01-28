package leetcode

import (
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestIsValidBST(t *testing.T) {
	type question struct {
		param  *structures.TreeNode
		answer bool
	}

	qs := []question{
		{
			param:  structures.IntsToTree([]int{2, 1, 3}),
			answer: true,
		},
		{
			param:  structures.IntsToTree([]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6}),
			answer: false,
		},
	}

	for _, q := range qs {
		got := isValidBST(q.param)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("\nexpected: %v \ngot: %v \n%v", q.answer, got, q)
		}
	}
}
