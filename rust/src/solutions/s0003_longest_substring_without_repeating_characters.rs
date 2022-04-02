/*
 * @lc app=leetcode id=3 lang=rust
 *
 * [3] Longest Substring Without Repeating Characters
 */

#[allow(dead_code)]
pub struct Solution {}

// @lc code=start
use std::cmp::max;
use std::collections::HashMap;

impl Solution {
    #[allow(dead_code)]
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut hm = HashMap::new();
        let mut ans = 0;
        let mut before = -1;
        let mut cur = 0;
        for c in s.chars() {
            if let Some(last) = hm.insert(c, cur) {
                before = max(before, last);
            }
            ans = max(ans, cur - before);
            cur += 1;
        }
        ans
    }
}
// @lc code=end

#[cfg(test)]

mod tests {
    use super::*;
    #[test]
    fn test_length_of_longest_substring() {
        assert_eq!(
            Solution::length_of_longest_substring(String::from("abcabcbb")),
            3
        )
    }
}
