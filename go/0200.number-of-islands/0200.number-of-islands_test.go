package leetcode

import (
	"math/rand"
	"testing"
)

type params struct {
	grid [][]byte
}

func TestNumIslands(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				grid: [][]byte{
					{'1', '1', '1', '1', '1'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			answer: 1,
		},
		{
			params: params{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			answer: 3,
		},
	}

	for _, q := range qs {
		grid := make([][]byte, len(q.grid))
		for i := range q.grid {
			grid[i] = make([]byte, len(q.grid[i]))
			copy(grid[i], q.grid[i])
		}
		got := numIslands(q.grid)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n%v", q.answer, got, grid)
		}
	}
}

func genGrid(nr int, nc int) [][]byte {
	grid := make([][]byte, nr)

	for i := range grid {
		grid[i] = make([]byte, nc)
	}

	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			grid[i][j] = byte(48 + rand.Intn(1))
		}
	}

	return grid
}

func BenchmarkNumIslands(b *testing.B) {
	grid := genGrid(10, 10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		numIslands(grid)
	}
}
func BenchmarkNumIslands2(b *testing.B) {
	grid := genGrid(10, 10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		numIslands2(grid)
	}
}
