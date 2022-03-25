/*
 * @lc app=leetcode id=1 lang=rust
 *
 * [1] Two Sum
 */

 #[allow(dead_code)]
pub struct Solution{}


// @lc code=start
use std::collections::HashMap;

impl Solution {
    #[allow(dead_code)]
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::with_capacity(nums.len());

        let mut res = vec![-1, -1];

        for (i, num) in nums.iter().enumerate() {
            match map.get(&(target-num)) {
                None => {
                    map.insert(*num, i as i32);
                }
                Some(j) => {
                    res = vec![*j, i as i32];
                    break
                }
            }
        }
        res
   }
}
// @lc code=end


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum() {
        assert_eq!(vec![0,1], Solution::two_sum(vec![2,7,11,15], 9));
        assert_eq!(vec![1,2], Solution::two_sum(vec![3,2,4], 6));
        assert_eq!(vec![0,1], Solution::two_sum(vec![3,3], 6));
    }
}