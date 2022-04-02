/*
 * @lc app=leetcode id=20 lang=rust
 *
 * [20] Valid Parentheses
 */

#[allow(dead_code)]
pub struct Solution {}

// @lc code=start
impl Solution {
    #[allow(dead_code)]
    pub fn is_valid(s: String) -> bool {
        let mut stack = Vec::new();
        for c in s.chars() {
            match c {
                '(' => stack.push(')'),
                '{' => stack.push('}'),
                '[' => stack.push(']'),
                ')' | '}' | ']' if Some(c) != stack.pop() => return false,
                _ => (),
            }
        }
        return stack.is_empty();
    }
}
// @lc code=end

#[cfg(test)]

mod tests {
    use super::*;
    #[test]
    fn test_is_valid() {
        assert!(Solution::is_valid("()".to_string()));
        assert!(Solution::is_valid("()[]{}".to_string()));
        assert!(!Solution::is_valid("(]".to_string()));
    }
}
