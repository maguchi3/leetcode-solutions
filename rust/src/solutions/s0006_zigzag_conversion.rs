/*
 * @lc app=leetcode id=6 lang=rust
 *
 * [6] Zigzag Conversion
 */
#[allow(dead_code)]
pub struct Solution {}

// @lc code=start
impl Solution {
    #[allow(dead_code)]
    pub fn convert(s: String, num_rows: i32) -> String {
        let mut step = if num_rows == 1 { 0 } else { -1 };
        let mut rows: Vec<String> = vec![String::new(); num_rows as usize];
        let mut i = 0;

        for c in s.chars() {
            rows[i as usize].push(c);
            if i == 0 || i == (num_rows as i32 - 1) {
                step = -step;
            }
            i += step;
        }

        rows.join("")
    }
}
// @lc code=end

#[cfg(test)]

mod tests {
    use super::*;
    #[test]
    fn test_convert() {
        assert_eq!(
            Solution::convert(String::from("PAYPALISHIRING"), 3),
            String::from("PAHNAPLSIIGYIR")
        )
    }
}
