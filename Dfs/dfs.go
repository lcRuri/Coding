package Dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//InorderTraversal 94. 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal/ 9/30
func InorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	return res
}

//IsSameTree 100. 相同的树 https://leetcode.cn/problems/same-tree/ 9/30
func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

//isSymmetric 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/ 9/30
func isSymmetric(root *TreeNode) bool {

	var flag = 1
	////如果左子树和右子树同时为空 满足条件
	//if root.Left==nil&&root.Right==nil{
	//	return true
	//}
	//
	////只有一个为空不满足条件
	//if root.Left==nil||root.Right==nil{
	//	return false
	//}
	//
	//if root.Left.Val!=root.Right.Val{
	//	return false
	//}

	var same func(n1 *TreeNode, n2 *TreeNode, x *int)

	same = func(n1 *TreeNode, n2 *TreeNode, x *int) {
		//if *x == 0 {
		//	return
		//}
		if n1 == nil && n2 == nil {
			return
		}

		if n1 == nil || n2 == nil {
			*x = 0
			return
		}

		if n1.Val != n2.Val {
			*x = 0
			return
		}

		same(n1.Left, n2.Right, x)
		same(n1.Right, n2.Left, x)
	}

	same(root.Left, root.Right, &flag)

	if flag == 0 {
		return false
	}

	return true
}

//buildTree 105. 从前序与中序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root

}
