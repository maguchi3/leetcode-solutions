package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func IntsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	head := dummy

	for _, v := range nums {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}

	return dummy.Next
}

func ListToInts(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
