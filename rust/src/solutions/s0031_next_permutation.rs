/*
 * @lc app=leetcode id=31 lang=rust
 *
 * [31] Next Permutation
 */
#[allow(dead_code)]
pub struct Solution {}

// @lc code=start
impl Solution {
    #[allow(dead_code)]
    pub fn next_permutation(nums: &mut Vec<i32>) {
        let len = nums.len();
        let (mut i, mut j) = (len - 1, len - 1);

        while i > 0 && nums[i - 1] >= nums[i] {
            i -= 1
        }

        if i > 0 {
            while j >= i && nums[j] <= nums[i - 1] {
                j -= 1
            }
            nums.swap(i - 1, j);
        }
        nums[i..len].reverse();
    }
}
// @lc code=end

#[cfg(test)]

mod tests {
    use super::*;

    #[test]
    fn test_next_permutation() {
        let mut v1 = vec![1, 2, 3];
        Solution::next_permutation(&mut v1);
        assert_eq!(v1, vec![1, 3, 2]);

        let mut v2 = vec![1, 1, 5];
        Solution::next_permutation(&mut v2);
        assert_eq!(v2, vec![1, 5, 1]);

        let mut v3 = vec![3, 2, 1];
        Solution::next_permutation(&mut v3);
        assert_eq!(v3, vec![1, 2, 3]);
    }
}
