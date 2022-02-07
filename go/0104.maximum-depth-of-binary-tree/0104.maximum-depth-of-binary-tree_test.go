package leetcode

import (
	"testing"

	"github.com/maguroid/leetcode/go/structures"
)

func TestMaxDepth(t *testing.T) {
	type question struct {
		param  *structures.TreeNode
		answer int
	}

	qs := []question{
		{
			param:  structures.IntsToTree([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}),
			answer: 3,
		},
		{
			param:  structures.IntsToTree([]int{1, structures.NULL, 2}),
			answer: 2,
		},
	}

	for _, q := range qs {
		got1, got2, got3 := maxDepth(q.param), maxDepth2(q.param), maxDepth3(q.param)

		if got1 != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got1, q)
		}
		if got2 != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got2, q)
		}
		if got3 != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got3, q)
		}
	}
}

func makeRange(min, max int) []int {
	arr := make([]int, max)
	for i := range arr {
		arr[i] = min + i
	}
	return arr
}

func BenchmarkMaxDepth(b *testing.B) {
	param := structures.IntsToTree(makeRange(1, 10000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxDepth(param)
	}
}

func BenchmarkMaxDepth2(b *testing.B) {
	param := structures.IntsToTree(makeRange(1, 10000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxDepth2(param)
	}
}

func BenchmarkMaxDepth3(b *testing.B) {
	param := structures.IntsToTree(makeRange(1, 10000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxDepth3(param)
	}
}
