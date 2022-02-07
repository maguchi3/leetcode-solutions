package leetcode

import (
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestDeleteDuplicate(t *testing.T) {
	type question struct {
		param  *structures.ListNode
		answer *structures.ListNode
	}

	qs := []question{
		{
			param:  structures.IntsToList([]int{1, 1, 2}),
			answer: structures.IntsToList([]int{1, 2}),
		},
		{
			param:  structures.IntsToList([]int{1, 1, 2, 3, 3}),
			answer: structures.IntsToList([]int{1, 2, 3}),
		},
	}

	for _, q := range qs {
		got := deleteDuplicates(q.param)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("\nexpected: %v \ngot: %v \n%v", structures.ListToInts(q.answer), structures.ListToInts(got), structures.ListToInts(q.param))
		}
	}
}
