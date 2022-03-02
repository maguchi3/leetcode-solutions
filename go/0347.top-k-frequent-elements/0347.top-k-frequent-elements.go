/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

package topKFrequent

import "container/heap"

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	freq := FreqMap{}
	for _, num := range nums {
		freq[num]++
	}

	return freq.quickSelect(0, len(freq)-1, k)
}

type FreqMap map[int]int

func (f FreqMap) keys() []int {
	keys := make([]int, len(f))

	i := 0
	for num := range f {
		keys[i] = num
		i++
	}

	return keys
}

func (f FreqMap) quickSelect(start, end, k int) []int {
	keys := f.keys()

	n := len(keys)
	target := n - k

	var pIdx int
	for {
		pIdx = f.pertition(keys, start, end)
		if pIdx == target {
			break
		}
		if pIdx < target {
			start = pIdx + 1
		} else {
			end = pIdx - 1
		}
	}

	return keys[pIdx:]
}

func (f FreqMap) pertition(nums []int, start, end int) int {
	pivot := f[nums[start]]
	left, right := start+1, end

	for left <= right {
		for left <= right && f[nums[left]] <= pivot {
			left++
		}
		for left <= right && f[nums[right]] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right], nums[start] = nums[start], nums[right]

	return right
}

// @lc code=end

// Priority Queue
func topKFrequent2(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	mHeap := make(MaxHeap, len(freq))

	i := 0
	for number, value := range freq {
		mHeap[i] = &Frequency{
			number: number,
			value:  value,
		}
		i++
	}
	heap.Init(&mHeap)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&mHeap).(*Frequency).number
	}

	return res
}

type Frequency struct {
	number int
	value  int
}

type MaxHeap []*Frequency

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Frequency))
}
