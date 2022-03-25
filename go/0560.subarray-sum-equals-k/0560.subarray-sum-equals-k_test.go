package subarraySum

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type question struct {
	nums   []int
	k      int
	answer int
}

func TestSubarraySum(t *testing.T) {
	qs := []question{
		{
			nums:   []int{1, 1, 1},
			k:      2,
			answer: 2,
		},
		{
			nums:   []int{1, 2, 3},
			k:      3,
			answer: 2,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.answer, subarraySum(q.nums, q.k))
	}
}

func randomSubarray(length int) []int {
	nums := make([]int, 1000)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		nums[i] = rand.Intn(1000) - rand.Intn(1000)
	}
	return nums
}

func BenchmarkSubarraySum(b *testing.B) {
	nums := randomSubarray(1000)
	k := 3
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		subarraySum(nums, k)
	}
}

func BenchmarkSubarraySum2(b *testing.B) {
	nums := randomSubarray(1000)
	k := 3
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		subarraySum2(nums, k)
	}
}
