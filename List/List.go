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

//middleNode 876. 链表的中间结点 https://leetcode.cn/problems/middle-of-the-linked-list/description/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs&q=go&orderBy=most_relevant
//快慢指针
func middleNode(head *ListNode) *ListNode {

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//removeNthFromEnd 19. 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	size := 1
	for p != nil && p.Next != nil {
		p = p.Next
		size++
	}

	if size == 1 {
		return nil
	}

	if n == size {
		head = head.Next
		return head
	}

	p = head
	q := head
	for i := 0; i < size-n; i++ {
		q = p
		p = p.Next

	}

	q.Next = p.Next

	return head
}
