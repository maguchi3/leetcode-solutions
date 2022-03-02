package coinChange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type params struct {
	coins  []int
	amount int
}

type question struct {
	params
	answer int
}

func TestCoinChange(t *testing.T) {
	qs := []question{
		{
			params: params{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			answer: 3,
		},
		{
			params: params{
				coins:  []int{2},
				amount: 3,
			},
			answer: -1,
		},
		{
			params: params{
				coins:  []int{1},
				amount: 0,
			},
			answer: 0,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.answer, coinChange(q.coins, q.amount))
	}

}

// top down(iterative) apploach scores high performance
func BenchmarkCoinChange1(b *testing.B) {
	coins := make([]int, 100)
	for i := range coins {
		coins[i] = i
	}

	params := params{
		coins:  coins,
		amount: 100,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChange(params.coins, params.amount)
	}
}

func BenchmarkCoinChange2(b *testing.B) {
	coins := make([]int, 100)
	for i := range coins {
		coins[i] = i + 1
	}

	params := params{
		coins:  coins,
		amount: 100,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChange2(params.coins, params.amount)
	}
}
