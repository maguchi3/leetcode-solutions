package countComponents

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type params struct {
	n     int
	edges [][]int
}

type question struct {
	params
	answer int
}

func TestCountComponents(t *testing.T) {
	qs := []question{
		{
			params: params{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {3, 4}},
			},
			answer: 2,
		},
		{
			params: params{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			},
			answer: 1,
		},
		{
			params: params{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			},
			answer: 1,
		},
		{
			params: params{
				n:     4,
				edges: [][]int{{2, 3}, {1, 2}, {1, 3}},
			},
			answer: 2,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.answer, countComponents(q.n, q.edges))
	}

}

func BenchmarkCountComponents(b *testing.B) {
	params := params{
		n:     4,
		edges: [][]int{{2, 3}, {1, 2}, {1, 3}},
	}

	for i := 0; i < b.N; i++ {
		countComponents(params.n, params.edges)
	}
}

func BenchmarkCountComponents2(b *testing.B) {
	params := params{
		n:     4,
		edges: [][]int{{2, 3}, {1, 2}, {1, 3}},
	}

	for i := 0; i < b.N; i++ {
		countComponents2(params.n, params.edges)
	}
}
