/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */
package intersection

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set1, set2 := make(map[int]struct{}), make(map[int]struct{})

	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	res := make([]int, 0)
	for key := range set1 {
		if _, ok := set2[key]; ok {
			res = append(res, key)
		}
	}
	return res
}

// @lc code=end
