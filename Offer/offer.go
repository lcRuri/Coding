package Offer

import (
	"math"
	"sort"
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
