package Offer

import (
	"math"
	"sort"
	"strconv"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//NumWays 剑指 Offer 10- II. 青蛙跳台阶问题 https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/description/?favorite=xb9nqhhg
func NumWays(n int) int {
	res := 0
	x, y := 1, 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return y
	}

	for i := 2; i <= n; i++ {
		if int64(x+y) > int64(math.Pow(float64(10), float64(9)))+7 {
			tmp := int64(x+y) % int64(math.Pow(float64(10), float64(9))+7)
			res = int(tmp)
		} else {
			res = x + y
		}
		x = y
		y = res

	}

	return res

}

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	pre1, pre2 := 1, 1
	for i := 2; i <= n; i++ {
		pre1, pre2 = pre2, (pre1+pre2)%1000000007
	}
	return pre2
}

//Fib 剑指 Offer 10- I. 斐波那契数列 https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/?favorite=xb9nqhhg
func Fib(n int) int {
	p1, p2 := 0, 1
	if n == 0 {
		return p1
	}
	if n == 1 {
		return p2
	}

	for i := 2; i <= n; i++ {
		p1, p2 = p2, (p1+p2)%1000000007
	}

	return p2
}

//MinArray 剑指 Offer 11. 旋转数组的最小数字 https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/description/?favorite=xb9nqhhg
func MinArray(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

//Exist 剑指 Offer 12. 矩阵中的路径 https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/description/?favorite=xb9nqhhg
//默认情况下，Go语言使用的是值传递(则先拷贝参数的副本，再将副本传递给函数)，即在调用过程中不会影响到实际参数
func Exist(board [][]byte, word string) bool {
	//DFS
	m := len(board)
	n := len(board[0])
	f := 0
	visit := make([][]byte, m)
	for i := range visit {
		visit[i] = make([]byte, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[f] {
				//如果第一个字符满足，则进行dfs
				flag := 0
				visit[i][j] = 1
				existDfs(board, i, j, word, f+1, &flag, visit)
				//清空
				visit[i][j] = 0
				if flag == len(word) {
					return true
				}
			}
		}
	}

	return false
}

func existDfs(board [][]byte, i int, j int, s string, loc int, flag *int, visit [][]byte) {
	*flag = *flag + 1
	if loc == len(s) {
		return
	}

	for k := 0; k < 4; k++ {

		x := i + dx[k]
		y := j + dy[k]

		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] == s[loc] && visit[x][y] != 1 {
			visit[x][y] = 1
			existDfs(board, x, y, s, loc+1, flag, visit)
			visit[x][y] = 0
		}

		if *flag == len(s) {
			return
		}
	}

	//如果上下左右都没有，要还原flag和visit的值
	*flag = *flag - 1

}

//CuttingRope 剑指 Offer 14- I. 剪绳子 https://leetcode.cn/problems/jian-sheng-zi-lcof/?favorite=xb9nqhhg
//动态规划 dp[n]=max(1*dp[n-1],2*dp[n-2],...,n-1*dp[1]
func CuttingRope(n int) int {
	dp := make([]int, n+1)

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}

	return dp[n]
}

//HammingWeight 剑指 Offer 15. 二进制中1的个数 https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/description/?favorite=xb9nqhhg
func HammingWeight(num uint32) int {
	res := 0
	for {
		bit := num % 2
		if bit == 1 {
			res++
		}
		num = num / 2
		if num == 0 {

			return res
		}

	}

}

func hammingWeight(num uint32) (ones int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ones++
		}
	}
	return
}

//PrintNumbers 剑指 Offer 17. 打印从1到最大的n位数 https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/description/?favorite=xb9nqhhg
func PrintNumbers(n int) []int {
	res := []int{}
	pow := math.Pow(float64(10), float64(n))
	for i := 1; i < int(pow); i++ {
		res = append(res, i)
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//DeleteNode 剑指 Offer 18. 删除链表的节点 https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/description/?favorite=xb9nqhhg
func DeleteNode(head *ListNode, val int) *ListNode {
	p, q := head, head
	for p != nil {
		if p.Val == val && q != p {
			q.Next = p.Next
			break
		} else if p.Val == val && q == p {
			return head.Next
		}
		q = p
		p = p.Next
	}
	return head
}

//exchange 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面 https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/description/?favorite=xb9nqhhg
func exchange(nums []int) []int {
	//l,r:=-1,-1//l指向偶数位置，r指向奇数位置
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			res = append(res, nums[i])
			//nums = append(nums[:i], nums[i+1:]...)
			//i--
		}

	}

	for _, num := range nums {
		if num%2 == 0 {
			res = append(res, num)
		}

	}

	return res
}

//getKthFromEnd 剑指 Offer 22. 链表中倒数第k个节点 https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/description/?favorite=xb9nqhhg
func getKthFromEnd(head *ListNode, k int) *ListNode {
	num := 0
	p := head
	for p != nil {
		num++
		p = p.Next
	}

	step := num - k

	for i := 0; i < step; i++ {
		head = head.Next
	}

	return head
}

//reverseList 剑指 Offer 24. 反转链表 https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/description/?favorite=xb9nqhhg
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head.Next
	q := head
	k := head
	for p != nil {
		q = p
		p = q.Next
		q.Next = k
		k = q

	}

	head.Next = nil

	return q
}
func reverseListI(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var flag = 0

func isSubStructure1(A *TreeNode, B *TreeNode) bool {
	//if A == nil && B == nil {
	//	return true
	//}
	//
	//if A == nil || B == nil {
	//	if flag == 0 {
	//		return false
	//	} else {
	//		return true
	//	}
	//}
	//
	//if A.Val != B.Val {
	//	flag = 0
	//	if A.Left == nil && A.Right == nil {
	//		return false
	//	}
	//	isSubStructure(A.Left, B)
	//	isSubStructure(A.Right, B)
	//	return false
	//}
	//
	//if A.Val == B.Val {
	//	flag = 1
	//	isSubStructure(A.Left, B.Left)
	//	isSubStructure(A.Right, B.Right)
	//}
	//if flag == 1 {
	//	return true
	//}
	//return false
	if A == nil || B == nil {
		return false
	}
	flag := 0
	node1 := []int{}
	node2 := []int{}
	node1 = appendNode(A, &node1)
	node2 = appendNode(B, &node2)
	j := 0
	for i := 0; i < len(node1); i++ {
		if flag == len(node2) {
			return true
		}

		if node1[i] == node2[j] {
			flag++
			j++
		} else {
			flag = 0
		}
	}

	return false
}

//先序遍历不行，有bug，需要用层次遍历
//层次遍历写法，利用队列，先进先出
func appendNode(node *TreeNode, t *[]int) []int {
	queue, root := []*TreeNode{node}, node

	for len(queue) > 0 {
		*t = append(*t, root.Val)
		root = queue[0]
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}

	}

	//appendNode(node.Left, t)
	//appendNode(node.Right, t)
	return *t
}

//isSubStructure 剑指 Offer 26. 树的子结构 https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/?favorite=xb9nqhhg
func isSubStructure(A *TreeNode, B *TreeNode) bool {

	// 根据题目判断条件，初始有空节点返回false
	if A == nil || B == nil {
		return false
	}

	queue, root := []*TreeNode{A}, A

	for len(queue) > 0 {

		root = queue[0]
		queue = queue[1:]

		if root.Val == B.Val && helper(root, B) {
			return true
		}

		// 还没遍历完，而且没有找到匹配的子树，继续遍历A子树
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}

	return false

}

func helper(A *TreeNode, B *TreeNode) bool {

	if B == nil {
		return true
	}

	if B == nil || A == nil {
		return false
	}

	if A.Val == B.Val {
		return helper(A.Left, B.Left) && helper(A.Right, B.Right) //递归
	} else {
		return false
	}
}

//mirrorTree 剑指 Offer 27. 二叉树的镜像 https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/description/?favorite=xb9nqhhg
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	mirrorTree(root.Left)
	mirrorTree(root.Right)

	return root
}

//isSymmetric 剑指 Offer 28. 对称的二叉树 https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/description/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return Symmetric(root.Left, root.Right)
}

func Symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left.Val == right.Val {
		return Symmetric(left.Left, right.Right) && Symmetric(left.Right, right.Left)
	} else {
		return false
	}
}

//spiralOrder 剑指 Offer 29. 顺时针打印矩阵 https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/description/?favorite=xb9nqhhg
func spiralOrder(matrix [][]int) []int {
	//如果为空，直接返回空
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	//求矩阵的行和列
	rows, columns := len(matrix), len(matrix[0])
	//初始化visited数组
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		//返回结果的数量
		total = rows * columns
		//返回值
		order = make([]int, total)
		//当前的打印位置
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		//先判断下一个位置是否正确
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		//如果到了边界或者已经访问过
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			//进行转弯，按照右、下、左、上的顺序
			directionIndex = (directionIndex + 1) % 4
		}
		//确认出下一个位置
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]

	}

	return order

}

//MinStack 剑指 Offer 30. 包含min函数的栈 https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/?favorite=xb9nqhhg
//使用了额外空间min []int 来保存每次添加数据，如果添加的数比现在最小的数小，那么就添加，否则继续添加当前最小的数
//进阶 不使用额外空间？？？
type MinStack struct {
	Stack  []int
	Length int
	min    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:  make([]int, 0),
		Length: 0,
		min:    []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min(x, this.min[len(this.min)-1]))
	}

	sort.Ints(this.min)
	this.Length++
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:this.Length-1]
	this.min = this.min[:len(this.min)-1]
	this.Length--
}

func (this *MinStack) Top() int {
	return this.Stack[this.Length-1]
}

//数组复制要用copy
func (this *MinStack) Min() int {
	//tmp := make([]int, this.Length)
	//copy(tmp, this.Stack)
	//sort.Ints(tmp)
	//return tmp[0]

	return this.min[len(this.min)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//func validateStackSequences(pushed []int, popped []int) bool {
//
//}

//levelOrder 剑指 Offer 32 - I. 从上到下打印二叉树 https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/description/?favorite=xb9nqhhg
//二叉树的层次遍历
func levelOrder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	node := root
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res

}

//levelOrder1 剑指 Offer 32 - II. 从上到下打印二叉树 II https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/description/
//分层添加到二维数组
func levelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		//将当前的所有添加到tmp中
		tmp := []int{}
		for i := 0; i < len(queue); i++ {
			tmp = append(tmp, queue[i].Val)
		}

		res = append(res, tmp)
		//清空原来的，添加接下来的
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]

	}

	return res
}

//levelOrder2 剑指 Offer 32 - III. 从上到下打印二叉树 III https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/description/
//负负得正
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	col := 0
	for len(queue) > 0 {
		//将当前的所有添加到tmp中
		tmp := []int{}
		for i := 0; i < len(queue); i++ {
			tmp = append(tmp, queue[i].Val)
		}

		res = append(res, tmp)
		//清空原来的，添加接下来的
		l := len(queue)
		if col%2 != 0 {
			for i := l - 1; i >= 0; i-- {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
		} else if col%2 == 0 {
			for i := l - 1; i >= 0; i-- {
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
			}
		}

		queue = queue[l:]
		col++
	}

	return res
}

//verifyPostorder 剑指 Offer 33. 二叉搜索树的后序遍历序列 https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/?favorite=xb9nqhhg
func verifyPostorder(postorder []int) bool {
	return true
}

//pathSum 剑指 Offer 34. 二叉树中和为某一值的路径 https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/description/
//匿名函数+递归（深度搜索）
func pathSum(root *TreeNode, target int) (res [][]int) {
	tmp := []int{}
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}

		tmp = append(tmp, root.Val)
		defer func() { tmp = tmp[:len(tmp)-1] }()

		if target-root.Val == 0 && root.Left == nil && root.Right == nil {
			cur := make([]int, len(tmp))
			copy(cur, tmp)
			res = append(res, cur)
			return
		}
		dfs(root.Left, target-root.Val)
		dfs(root.Right, target-root.Val)
	}

	dfs(root, target)
	return res
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

//copyRandomList 剑指 Offer 35. 复杂链表的复制 https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/description/?favorite=xb9nqhhg
//递归进行拷贝 哈希表用来防止重复拷贝
func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

//Permutation 剑指 Offer 38. 字符串的排列 https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/?favorite=xb9nqhhg
//组合问题
//从左往右第一个未被填入的字符
func Permutation(s string) []string {
	str := []byte(s)
	res := []string{}

	//sort.Slice(str, func(i, j int) bool { return str[i] < str[j] })

	var backtrack func()

	length := len(str)

	tmp := make([]byte, 0, length)
	visited := make([]bool, length)

	backtrack = func() {
		if len(tmp) == length {

			res = append(res, string(tmp))
			return
		}

		for j := 0; j < len(visited); j++ {
			if visited[j] == true || j > 0 && visited[j-1] == false && str[j] == str[j-1] {
				continue
			}
			tmp = append(tmp, str[j])
			visited[j] = true
			backtrack()
			tmp = tmp[:len(tmp)-1]
			visited[j] = false
		}

	}

	backtrack()

	return res

}

//majorityElement 剑指 Offer 39. 数组中出现次数超过一半的数字 https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/?favorite=xb9nqhhg
func majorityElement(nums []int) int {
	length := len(nums)
	dic := map[int]int{}

	for i := 0; i < len(nums); i++ {
		dic[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		if dic[nums[i]] > length/2 {
			return nums[i]
		}
	}

	return 0
}

//GetLeastNumbers 剑指 Offer 40. 最小的k个数 https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/description/?favorite=xb9nqhhg
//考的排序
func GetLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

//maxSubArray 剑指 Offer 42. 连续子数组的最大和 https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/?favorite=xb9nqhhg
//思路：如果你能让我更好，那我们在一起，否则舍弃，直到走完，看看一路走来哪个过程最好
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	ans := math.MinInt32

	for _, d := range dp {
		if d > ans {
			ans = d
		}
	}

	return ans
}

//TranslateNum 剑指 Offer 46. 把数字翻译成字符串 https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/description/?favorite=xb9nqhhg
//动态规划
//有条件的青蛙跳格子(斐波那契数列) 只有当数字匹配的满足>=10 <=25时，才能跳两格，即dp[i]=dp[i-1]+dp[i-2],否则dp[i]=dp[i-1]
func TranslateNum(num int) int {

	strs := strconv.Itoa(num)
	dp := make([]int, len(strs)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= len(strs); i++ {
		if strs[i-2:i] >= "10" && strs[i-2:i] <= "25" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}

	}

	return dp[len(dp)-1]
}

//MaxValue 剑指 Offer 47. 礼物的最大价值 https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/description/?favorite=xb9nqhhg
//同leetcode62题 62. 不同路径 https://leetcode.cn/problems/unique-paths/
func MaxValue(grid [][]int) int {

	row, col := len(grid), len(grid[0])
	dp := make([]int, col)
	dp[0] = grid[0][0]
	for i := 1; i < col; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < row; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < col; j++ {
			dp[j] = max(dp[j-1], dp[j]) + grid[i][j]
		}
	}

	return dp[len(dp)-1]
}
