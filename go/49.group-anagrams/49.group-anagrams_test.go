package leetcode

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type question struct {
		param  []string
		answer [][]string
	}

	qs := []question{
		{
			param:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			answer: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			param:  []string{""},
			answer: [][]string{{""}},
		},
	}

	for _, q := range qs {
		got := groupAnagrams(q.param)

		if !reflect.DeepEqual(got, q.answer) {
			t.Errorf("expected: %v \n got: %v \n %v", q.answer, got, q)
		}
	}
}
