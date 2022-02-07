package leetcode

import (
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

type params struct {
	preorder []int
	inorder  []int
}

func TestBuildTree(t *testing.T) {
	type question struct {
		params
		answer *TreeNode
	}

	qs := []question{
		{
			params: params{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			answer: structures.IntsToTree([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}),
		},
		{
			params: params{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			answer: structures.IntsToTree([]int{-1}),
		},
	}

	for _, q := range qs {
		got := buildTree(q.preorder, q.inorder)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("\nexpected: %v \ngot:%v \n%v", q.answer, got, q)
		}
	}
}
