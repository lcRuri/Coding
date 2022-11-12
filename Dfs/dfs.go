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

//buildTree 重建二叉树
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

//FloodFill 733. 图像渲染 https://leetcode.cn/problems/flood-fill/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zsc
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	level := image[sr][sc]
	image[sr][sc] = color
	if sr-1 >= 0 && image[sr-1][sc] == level && image[sr-1][sc] != color {
		FloodFill(image, sr-1, sc, color)
	}
	if sr+1 < len(image) && image[sr+1][sc] == level && image[sr+1][sc] != color {
		FloodFill(image, sr+1, sc, color)
	}

	if sc-1 >= 0 && image[sr][sc-1] == level && image[sr][sc-1] != color {
		FloodFill(image, sr, sc-1, color)
	}

	if sc+1 < len(image[0]) && image[sr][sc+1] == level && image[sr][sc+1] != color {
		FloodFill(image, sr, sc+1, color)
	}

	return image

}

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

//MaxAreaOfIsland 695. 岛屿的最大面积 https://leetcode.cn/problems/max-area-of-island/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs
func MaxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0

	flags := make([][]bool, m)
	for i := 0; i < m; i++ {
		flags[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area = 0
			if grid[i][j] == 1 && flags[i][j] != true {
				area = dfs(grid, i, j, flags)

			}
			if area > res {
				res = area
			}

		}
	}

	return res
}

var area int

func dfs(grid [][]int, x int, y int, flags [][]bool) int {
	if grid[x][y] == 1 && flags[x][y] != true {
		area++
		flags[x][y] = true
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx >= 0 && mx < len(grid) && my >= 0 && my < len(grid[0]) {
				dfs(grid, mx, my, flags)
			}
		}
	}

	return area
}

//mergeTrees 617. 合并二叉树 https://leetcode.cn/problems/merge-two-binary-trees/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	//if root1 == nil && root2 == nil {
	//	return nil
	//}
	//root := &TreeNode{
	//	Val:   0,
	//	Left:  nil,
	//	Right: nil,
	//}
	//if root1 == nil && root2 != nil {
	//	root.Val = root2.Val
	//	root.Left = mergeTrees(nil, root2.Left)
	//	root.Right = mergeTrees(nil, root2.Right)
	//	return root
	//}
	//if root1 != nil && root2 == nil {
	//	root.Val = root1.Val
	//	root.Left = mergeTrees(root1.Left, nil)
	//	root.Right = mergeTrees(root1.Right, nil)
	//	return root
	//}
	//if root1 != nil && root2 != nil {
	//	root.Val = root1.Val + root2.Val
	//}
	//
	//root.Left = mergeTrees(root1.Left, root2.Left)
	//root.Right = mergeTrees(root1.Right, root2.Right)
	//
	//return root
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//connect 116. 填充每个节点的下一个右侧节点指针 https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}

	connect(root.Left)
	connect(root.Right)

	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// if head == nil {
	// 	return nil
	// }

	// if head.Next==nil{
	// 	return head
	// }

	// p:=head.Next
	// head.Next=nil
	// q:=head
	// for p!=nil {
	// 	tmp:=p.Next
	// 	p.Next=q
	//     if tmp==nil{
	//         return p
	//     }
	// 	q=p
	// 	p=tmp
	// }

	// return p

	var prev *ListNode
	p := head
	for p != nil {
		q := p.Next
		p.Next = prev
		prev = p
		p = q
	}
	return prev
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {

			p.Next = list1
			p = p.Next
			list1 = list1.Next
		} else {

			p.Next = list2
			p = p.Next
			list2 = list2.Next
		}

		if list1 == nil {
			p.Next = list2
		}
		if list2 == nil {
			p.Next = list1
		}
	}

	return head.Next
}
