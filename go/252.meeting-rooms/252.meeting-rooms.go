/*
* @lc app=leetcode id=252 lang=golang
* [252] Meeting Rooms
 */
package leetcode

import "sort"

// @lc code=start
func canAttendMeetings(intervals [][]int) bool {
	// sort by ascending order
	sort.Slice(intervals, func(i, j int) bool {
		start1, start2 := intervals[i][0], intervals[j][0]
		end1, end2 := intervals[i][1], intervals[j][1]
		if start1 == start2 {
			return end1 < end2
		}
		return start1 < start2
	})

	// check timeslot overlap
	for i := 0; i < len(intervals)-1; i++ {
		end1, start2 := intervals[i][1], intervals[i+1][0]
		if end1 > start2 {
			return false
		}
	}

	return true
}

// @lc code=end
