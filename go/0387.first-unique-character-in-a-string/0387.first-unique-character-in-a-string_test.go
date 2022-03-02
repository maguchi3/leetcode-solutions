package firstUniqueChar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	s      string
	answer int
}

func TestFirstUniqueChar(t *testing.T) {
	qs := []question{
		{
			s:      "leetcode",
			answer: 0,
		},
		{
			s:      "loveleetcode",
			answer: 2,
		},
		{
			s:      "aabb",
			answer: -1,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.answer, firstUniqChar(q.s))
	}
}

func BenchmarkFirstUniqueChar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstUniqChar("benchmarktest")
	}
}

func BenchmarkFirstUniqueChar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstUniqChar2("benchmarktest")
	}
}
