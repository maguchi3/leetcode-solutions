/*
 * @lc app=leetcode id=2 lang=rust
 *
 * [2] Add Two Numbers
 */

pub struct Solution {}
use crate::structs::linked_list::ListNode;

// @lc code=start
impl Solution {
    #[allow(dead_code)]
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (l1, l2) {
            (None, None) => None,
            (Some(n), None) | (None, Some(n)) => Some(n),
            (Some(n1), Some(n2)) => {
                let sum = n1.val + n2.val;
                if sum < 10 {
                    Some(Box::new(ListNode {
                        val: sum,
                        next: Solution::add_two_numbers(n1.next, n2.next),
                    }))
                } else {
                    let carry = Some(Box::new(ListNode::new(1)));

                    Some(Box::new(ListNode {
                        val: sum - 10,
                        next: Solution::add_two_numbers(
                            Solution::add_two_numbers(carry, n1.next),
                            n2.next,
                        ),
                    }))
                }
            }
        }
    }
}
// @lc code=end

#[cfg(test)]

mod tests {
    use crate::structs::linked_list::vec_to_list;

    use super::*;
    #[test]
    fn test_add_two_numbers() {
        let got = Solution::add_two_numbers(vec_to_list(vec![2, 4, 3]), vec_to_list(vec![5, 6, 4]));

        assert_eq!(got, vec_to_list(vec![7, 0, 8]));
    }
}
