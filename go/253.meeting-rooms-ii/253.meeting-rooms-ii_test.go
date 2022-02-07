package leetcode

import "testing"

type params struct {
	intervals [][]int
}

func TestMinMeetingIntervals(t *testing.T) {
	type question struct {
		params
		answer int
	}

	qs := []question{
		{
			params: params{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			answer: 2,
		},
		{
			params: params{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			answer: 1,
		},
	}

	for _, q := range qs {
		got := minMeetingRooms(q.intervals)

		if got != q.answer {
			t.Errorf("\nexpected: %d \ngot: %d \n %v", q.answer, got, q)
		}
	}
}
