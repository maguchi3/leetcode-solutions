package leetcode

import (
	"reflect"
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

type params struct {
	head *ListNode
}

func TestReverseList(t *testing.T) {
	type question struct {
		params
		answer *ListNode
	}

	qs := []question{
		{
			params: params{
				head: structures.IntsToList([]int{1, 2, 3, 4, 5}),
			},
			answer: structures.IntsToList([]int{5, 4, 3, 2, 1}),
		},
		{
			params: params{
				head: structures.IntsToList([]int{1, 2}),
			},
			answer: structures.IntsToList([]int{2, 1}),
		},
	}

	for _, q := range qs {
		got := reverseList(q.head)
		i1, i2 := structures.ListToInts(q.answer), structures.ListToInts(got)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("\nexpected: %v \ngot: %v \n%v", i1, i2, q)
		}
	}
}

func BenchmarkReverseList1(b *testing.B) {
	node := structures.IntsToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < b.N; i++ {
		reverseList(node)
	}
}

func BenchmarkReverseList2(b *testing.B) {
	node := structures.IntsToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < b.N; i++ {
		reverseList2(node)
	}
}
func BenchmarkReverseList3(b *testing.B) {
	node := structures.IntsToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < b.N; i++ {
		reverseList3(node)
	}
}
