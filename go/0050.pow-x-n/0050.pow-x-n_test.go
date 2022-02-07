package leetcode

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	type params struct {
		x float64
		n int
	}
	type question struct {
		params
		answer float64
	}

	qs := []question{
		{
			params: params{
				x: 2,
				n: 10,
			},
			answer: 1024,
		},
		{
			params: params{
				x: 2.1,
				n: 3,
			},
			answer: 9.261,
		},
		{
			params: params{
				x: 2,
				n: -2,
			},
			answer: 0.25,
		},
	}

	for _, q := range qs {
		got := myPow(q.x, q.n)

		if math.Abs(got-q.answer) > 0.0001 {
			t.Errorf("expected: %f \ngot: %f \n%v", q.answer, got, q)
		}
	}
}
