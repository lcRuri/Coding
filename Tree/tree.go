package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//isSubtree 572. 另一棵树的子树 https://leetcode.cn/problems/subtree-of-another-tree/?envType=study-plan&id=suan-fa-ji-chu&plan=algorithms&plan_progress=1ah3sii
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}

	return false
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	} else if root.Left != nil && root.Next != nil && root.Next.Left != nil {
		root.Left.Next = root.Next.Left
	} else if root.Left != nil && root.Next != nil && root.Next.Right != nil {
		root.Left.Next = root.Next.Right
	}

	if root.Next != nil && root.Right != nil && root.Next.Left != nil {
		root.Right.Next = root.Next.Left
	} else if root.Next != nil && root.Right != nil && root.Next.Right != nil {
		root.Right.Next = root.Next.Right
	}
	connect(root.Left)
	connect(root.Right)

	return root
}
