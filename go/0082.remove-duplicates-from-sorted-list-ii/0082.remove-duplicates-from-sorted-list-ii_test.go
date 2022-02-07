package leetcode

import (
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestDeletDuplicates(t *testing.T) {
	type question struct {
		param  *structures.ListNode
		answer *structures.ListNode
	}

	qs := []question{
		{
			param:  structures.IntsToList([]int{1, 2, 3, 3, 4, 4, 5}),
			answer: structures.IntsToList([]int{1, 2, 5}),
		},
	}

	for _, q := range qs {
		got := deleteDuplicates(q.param)

		if !reflect.DeepEqual(*got, *q.answer) {
			t.Errorf("\nexpected: %v \ngot: %v \n%v", structures.ListToInts(q.answer), structures.ListToInts(got), structures.ListToInts(q.param))
		}
	}
}
