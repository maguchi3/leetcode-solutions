package isSubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	s      string
	t      string
	answer bool
}

func TestIsSubsequence(t *testing.T) {
	qs := []question{
		{
			s:      "abc",
			t:      "ahbgdc",
			answer: true,
		},
		{
			s:      "axc",
			t:      "ahbgdc",
			answer: false,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.answer, isSubsequence(q.s, q.t))
	}

}
