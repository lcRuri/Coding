package List

type ListNode struct {
	Val  int
	Next *ListNode
}

//numComponents 817. 链表组件 https://leetcode.cn/problems/linked-list-components/
func numComponents(head *ListNode, nums []int) (ans int) {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}
	for inSet := false; head != nil; head = head.Next {
		if _, ok := set[head.Val]; !ok {
			inSet = false
		} else if !inSet {
			inSet = true
			ans++
		}
	}
	return
}

//reversePrint 剑指 Offer 06. 从尾到头打印链表 https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/description/?favorite=xb9nqhhg
func reversePrint(head *ListNode) []int {
	if head == nil {
		res := make([]int, 0)
		return res
	}

	return append(reversePrint(head.Next), head.Val)
}
