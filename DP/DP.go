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

//MaxProfitII 122. 买卖股票的最佳时机 II https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func MaxProfitII(prices []int) int {
	//dp[i][j] i表示第i天是否持有的 j=0表示不持有股票 j=1表示持有股票
	//dp[0][0]=0 dp[0][1]=-7
	//dp[1][0]=0/-6 第1天持有不持有可以是前一天没有买或者前一天买了然后卖出 dp[1][1]=-7/-1 第1天持有股票 可以是第1天没有没有买入持续前面的状态，也可以是前面一天就没有然后第1天买入
	//dp[2][0]=0/4 dp[2][1]=-1/-5d

	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

//minSwap 801. 使序列递增的最小交换次数 https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing/
//https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing/solution/shi-xu-lie-di-zeng-de-zui-xiao-jiao-huan-ux2y/
func minSwap(nums1, nums2 []int) int {
	n := len(nums1)
	a, b := 0, 1
	for i := 1; i < n; i++ {
		at, bt := a, b
		a, b = n, n
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			a = min(a, at)
			b = min(b, bt+1)
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			a = min(a, bt)
			b = min(b, at+1)
		}
	}
	return min(a, b)
}

func Rob(nums []int) int {
	//dp[i]=max(dp[i-1],dp[i-2]+nums[i])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) > 1 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	res := dp[0]
	for _, v := range dp {
		res = max(res, v)
	}

	return res
}

//MaxSubArray 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/
//https://leetcode.cn/problems/maximum-subarray/solution/dong-tai-gui-hua-fen-zhi-fa-python-dai-ma-java-dai/
func MaxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	//for i := 0; i < len(nums); i++ {
	//	tmp := 0
	//	dp[i] = nums[i]
	//	for j := i; j < len(nums); j++ {
	//		tmp += nums[j]
	//		if tmp > dp[i] {
	//			dp[i] = tmp
	//		}
	//
	//	}
	//}

	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	res := dp[0]
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
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

//NthUglyNumber 264. 丑数 II https://leetcode.cn/problems/ugly-number-ii/ 9/28
//思路：动态规划 https://leetcode.cn/problems/ugly-number-ii/solution/chou-shu-ii-by-leetcode-solution-uoqd/
//表示下一个丑数是当前指针指向的丑数乘以对应的质因数。初始时，三个指针的值都是 11。
func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min3(x2, x3, x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}

	return dp[n]

}

//GetKthMagicNumber 面试题 17.09. 第 k 个数 https://leetcode.cn/problems/get-kth-magic-number-lcci/  9/28
func GetKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		x3, x5, x7 := dp[p3]*3, dp[p5]*5, dp[p7]*7
		dp[i] = min3(x3, x5, x7)
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
		if dp[i] == x7 {
			p7++
		}
	}

	return dp[k]
}
func min3(x, y, z int) int {
	var mi = 0
	if x <= y {
		mi = x
	} else {
		mi = y
	}

	if z < mi {
		mi = z
	}

	return mi

}

//GenerateParenthesis 22. 括号生成 9/30
func GenerateParenthesis(n int) []string {
	return nil
}
