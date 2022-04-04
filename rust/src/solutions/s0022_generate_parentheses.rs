/*
 * @lc app=leetcode id=22 lang=rust
 *
 * [22] Generate Parentheses
 */
#[allow(dead_code)]
pub struct Solution {}

// @lc code=start
impl Solution {
    #[allow(dead_code)]
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        fn back_track(s: String, open: i32, close: i32) -> Vec<String> {
            let mut res = vec![];

            if open == 0 && close == 0 {
                return vec![s];
            }

            if open > 0 {
                res.append(&mut back_track(s.clone() + "(", open - 1, close + 1));
            }

            if close > 0 {
                res.append(&mut back_track(s.clone() + ")", open, close - 1));
            }
            res
        }
        back_track("".to_string(), n, 0)
    }
}
// @lc code=end

#[cfg(test)]

mod tests {
    use super::*;
    #[test]

    fn test_generate_parenthesis() {
        assert_eq!(
            Solution::generate_parenthesis(3),
            vec!["((()))", "(()())", "(())()", "()(())", "()()()"]
        )
    }
}
