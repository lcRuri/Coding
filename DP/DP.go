package DP

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MinimumTotal 120. 三角形最小路径和 https://leetcode.cn/problems/triangle/
func MinimumTotal(triangle [][]int) int {

	n := len(triangle)

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	f[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}
	return ans
}

//NumSquares 279. 完全平方数 https://leetcode.cn/problems/perfect-squares/
//思路：dp[i] 表示i的完全平方和的最少数量，dp[i - j * j] + 1表示减去一个完全平方数j的完全平方之后的数量加1就等于dp[i]，只要在dp[i], dp[i - j * j] + 1中寻找一个较少的就是最后dp[i]的值。
func NumSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

// MaxProduct 152. 乘积最大子数组 https://leetcode.cn/problems/maximum-product-subarray/ 2002/9/25
//思路：动态规划 维护一个前n-1个数的最大值和最小值切片
func MaxProduct(nums []int) int {
	l := len(nums)
	res := make([][]int, 2)
	//res[0]维护最大值 res[1]维护最小值
	for i := range res {
		res[i] = make([]int, 1)
	}

	res[0][0] = nums[0]
	res[1][0] = nums[0]

	for i := 1; i < l; i++ {
		maxTmp := res[0][i-1]
		minTmp := res[1][i-1]

		//最大值不一定是正数乘以前n-1个的最大值，也可能是第n个数为负数乘以前n-1个数的最小值
		//求最大值
		maxi := max(nums[i]*maxTmp, max(nums[i]*minTmp, nums[i]))

		res[0] = append(res[0], maxi)

		//求最小值
		mini := min(nums[i]*minTmp, min(nums[i]*maxTmp, nums[i]))

		res[1] = append(res[1], mini)

	}

	max := res[0][0]
	for i := 0; i < l; i++ {
		if res[0][i] > max {
			max = res[0][i]
		}
	}

	return max

}

//MaxProduct1 优化 由于第 i 个状态只和第 i - 1 个状态相关，根据「滚动数组」思想，我们可以只用两个变量来维护 i - 1 时刻的状态，一个维护 f_{max} ，一个维护 f_{min}
func MaxProduct1(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans

}

func MaxProfit1(prices []int) int {
	var profit int = 0

	for i := 0; i < len(prices)-1; i++ {
		pay := prices[i]
		for j := i + 1; j < len(prices); j++ {
			mo := prices[j]
			profit = max(profit, mo-pay)
		}

	}

	return profit
}

// MaxProfit 121. 买卖股票的最佳时机 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {
	mi := prices[0]
	ma := 0
	for i := 1; i < len(prices); i++ {
		ma = max(ma, prices[i]-mi)
		mi = min(mi, prices[i])
	}

	return ma
}
