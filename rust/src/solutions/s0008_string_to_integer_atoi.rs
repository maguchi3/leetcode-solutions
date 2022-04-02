/*
 * @lc app=leetcode id=8 lang=rust
 *
 * [8] String to Integer (atoi)
 */
#[allow(dead_code)]
pub struct Solution {}

// @lc code=start
impl Solution {
    #[allow(dead_code)]
    pub fn my_atoi(s: String) -> i32 {
        use std::i32;

        let mut ans: i32 = 0;
        let mut is_negative = false;
        let mut has_started = false;

        for c in s.chars() {
            match c {
                ' ' if !has_started => continue,
                '+' if !has_started => has_started = true,
                '-' if !has_started & !is_negative => {
                    has_started = true;
                    is_negative = true;
                    ans = -ans
                }
                '0'..='9' => {
                    has_started = true;

                    ans = match ans
                        .checked_mul(10)
                        .and_then(|prod| prod.checked_add(c as i32 - '0' as i32))
                    {
                        Some(v) => v,
                        None => return if is_negative { i32::MIN } else { i32::MAX },
                    };
                }
                _ => break,
            }
        }

        if is_negative {
            -ans
        } else {
            ans
        }
    }
}
// @lc code=end
#[cfg(test)]

mod tests {
    use super::*;
    #[test]
    fn test_my_atoi() {
        assert_eq!(Solution::my_atoi("42".to_string()), 42);
        assert_eq!(Solution::my_atoi("-42".to_string()), -42);
        assert_eq!(Solution::my_atoi("4193".to_string()), 4193);
    }
}
