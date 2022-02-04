package leetcode

import "testing"

type params struct {
	n uint32
}

func TestReverseBits(t *testing.T) {
	type question struct {
		params
		answer uint32
	}

	qs := []question{
		{
			params: params{
				n: 43261596,
			},
			answer: 964176192,
		},
		{
			params: params{
				n: 4294967293,
			},
			answer: 3221225471,
		},
	}

	for _, q := range qs {
		got := reverseBits(q.n)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, q)
		}
	}
}
