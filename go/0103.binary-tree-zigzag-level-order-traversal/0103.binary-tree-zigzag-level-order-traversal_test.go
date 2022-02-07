package leetcode

import (
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestZigzagLevelOrder(t *testing.T) {
	type question struct {
		param  *structures.TreeNode
		answer [][]int
	}

	qs := []question{
		{
			param:  structures.IntsToTree([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}),
			answer: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			param:  structures.IntsToTree([]int{1}),
			answer: [][]int{{1}},
		},
		{
			param:  structures.IntsToTree([]int{}),
			answer: [][]int{},
		},
		{
			param:  structures.IntsToTree([]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}),
			answer: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	for _, q := range qs {
		got := zigzagLevelOrder(q.param)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("\nexpected: %v \ngot: %v \n%v", q.answer, got, q.param)
		}
	}

}
