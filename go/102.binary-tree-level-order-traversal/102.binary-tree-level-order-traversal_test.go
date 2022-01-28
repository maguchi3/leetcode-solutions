package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestLevelOrder(t *testing.T) {
	type question struct {
		param  *structures.TreeNode
		answer [][]int
	}

	qs := []question{
		{
			param:  structures.IntsToTree([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}),
			answer: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			param:  structures.IntsToTree([]int{1}),
			answer: [][]int{{1}},
		},
		{
			param:  structures.IntsToTree([]int{}),
			answer: nil,
		},
	}

	for _, q := range qs {
		got := levelOrder(q.param)
		fmt.Println(q.param, got)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("\nexpected: %T \ngot: %T \n%v", q.answer, got, q.param)
		}
	}
}
