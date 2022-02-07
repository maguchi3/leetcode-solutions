package leetcode

import "testing"

type params struct {
	intervals [][]int
}

func TestCanAttendMeetings(t *testing.T) {
	type question struct {
		params
		answer bool
	}

	qs := []question{
		{
			params: params{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			answer: false,
		},
		{
			params: params{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			answer: true,
		},
	}

	for _, q := range qs {
		got := canAttendMeetings(q.intervals)

		if got != q.answer {
			t.Errorf("\nexpected: %t \ngot: %t \n%v", q.answer, got, q)
		}
	}
}
