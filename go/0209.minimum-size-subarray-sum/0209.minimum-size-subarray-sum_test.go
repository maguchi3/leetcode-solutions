package leetcode

import (
	"math/rand"
	"testing"

	"github.com/maguroid/leetcode/go/utils"
)

type params struct {
	target int
	nums   []int
}

func TestMinSubArrayLen(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			answer: 2,
		},
		{
			params: params{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			answer: 1,
		},
		{
			params: params{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			answer: 0,
		},
	}

	for _, q := range qs {
		got := minSubArrayLen(q.target, q.nums)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	seq := utils.Sequence(1, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		minSubArrayLen(rand.Intn(100), seq)
	}
}

func BenchmarkMinSubArrayLen2(b *testing.B) {
	seq := utils.Sequence(1, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		minSubArrayLen2(rand.Intn(100), seq)
	}
}
