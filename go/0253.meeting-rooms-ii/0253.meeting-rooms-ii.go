package leetcode

import (
	"container/heap"
	"sort"
)

// define IntHeap
type Occupations []int

func (oc Occupations) Len() int           { return len(oc) }
func (oc Occupations) Less(i, j int) bool { return oc[i] < oc[j] }
func (oc Occupations) Swap(i, j int)      { oc[i], oc[j] = oc[j], oc[i] }
func (oc Occupations) Peek() int          { return oc[0] }

func (oc *Occupations) Push(x interface{}) {
	*oc = append(*oc, x.(int))
}
func (oc *Occupations) Pop() interface{} {
	old := *oc
	n := len(old)
	x := old[n-1]
	*oc = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	// sort intervals by ascending order
	sort.Slice(intervals, func(i, j int) bool {
		s1, s2 := intervals[i][0], intervals[j][0]
		e1, e2 := intervals[i][1], intervals[j][1]
		if s1 == s2 {
			return e1 < e2
		}
		return s1 < s2
	})

	// initialize IntHeap
	oc := new(Occupations)
	heap.Init(oc)

	rooms := 1

	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if oc.Len() > 0 {
			if start < oc.Peek() {
				rooms++
			} else {
				heap.Pop(oc)
			}
		}
		heap.Push(oc, end)
	}
	return rooms
}
