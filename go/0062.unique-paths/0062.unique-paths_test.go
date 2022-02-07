package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	type params struct {
		m int
		n int
	}
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				m: 3,
				n: 7,
			},
			answer: 28,
		},
	}

	for _, q := range qs {
		got := uniquePaths(q.m, q.n)

		if got != q.answer {
			t.Errorf("expected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
