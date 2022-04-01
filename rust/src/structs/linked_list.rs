#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub fn vec_to_list(vec: Vec<i32>) -> Option<Box<ListNode>> {
    let mut cur = None;

    for &v in vec.iter().rev() {
        let mut node = ListNode::new(v);
        node.next = cur;
        cur = Some(Box::new(node));
    }
    cur
}

#[cfg(test)]

mod tests {
    use super::*;
    #[test]
    fn test_vec_to_list() {
        let vec = vec![1, 2, 3];
        let node1 = ListNode::new(3);
        let node2 = ListNode {
            val: 2,
            next: Some(Box::new(node1)),
        };
        let node3 = ListNode {
            val: 1,
            next: Some(Box::new(node2)),
        };

        assert_eq!(vec_to_list(vec).unwrap(), Box::new(node3));
    }
}
