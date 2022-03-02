/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */
package kSmallestPairs

import (
	"container/heap"
)

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pairs := make(MinHeap, 0, len(nums1))

	for i := 0; i < len(nums1) && i < k; i++ {
		pairs = append(pairs, &Pair{0, nums1[i], nums2[0]})
	}
	heap.Init(&pairs)

	result := make([][]int, 0, k)
	pos := 0
	for pos < k && len(pairs) > 0 {
		p := pairs[0]
		result = append(result, []int{p.x, p.y})
		if p.index+1 < len(nums2) {
			pairs[0] = &Pair{p.index + 1, p.x, nums2[p.index+1]}
			heap.Fix(&pairs, 0)
		} else {
			heap.Pop(&pairs)
		}
		pos++
	}

	return result
}

type Pair struct {
	index int
	x, y  int
}

type MinHeap []*Pair

func (m MinHeap) Len() int { return len(m) }
func (m MinHeap) Less(i, j int) bool {
	if m[i] == nil || m[j] == nil {
		return false
	}
	return m[i].x+m[i].y < m[j].x+m[j].y
}
func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *MinHeap) Push(x interface{}) {
	n := len(*m)
	pair := x.(*Pair)
	pair.index = n
	*m = append(*m, pair)

}
func (m *MinHeap) Pop() interface{} {
	tmp := *m
	n := len(tmp)
	res := tmp[n-1]
	tmp[n-1] = nil
	*m = tmp[:n-1]

	return res
}

// @lc code=end
