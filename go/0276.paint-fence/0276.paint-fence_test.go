package leetcode

import (
	"math/rand"
	"testing"
)

type params struct {
	n int
	k int
}

func TestNumWays(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				n: 3,
				k: 2,
			},
			answer: 6,
		},
		{
			params: params{
				n: 1,
				k: 1,
			},
			answer: 1,
		},
	}

	for _, q := range qs {
		got := numWays(q.n, q.k)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, q.params)
		}
	}
}

/*
 BenchmarkNumWays-8    	52083984	       122.7 ns/op	     288 B/op	       1 allocs/op
 BenchmarkNumWays2-8   	36117510	      1111 ns/op	     919 B/op	       3 allocs/op
*/

func BenchmarkNumWays(b *testing.B) {
	n, k := rand.Intn(49)+1, rand.Intn(999)+1
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		numWays(n, k)
	}
}

func BenchmarkNumWays2(b *testing.B) {
	n, k := rand.Intn(49)+1, rand.Intn(999)+1
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		numWays2(n, k)
	}
}
