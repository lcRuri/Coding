package Bs

import (
	"math"
)

//二分法一般有left right mid
//判断mid和要比较值的大小，如果mid>target,left=mid+1/right=mid-1,重新算出mid
//判断mid和要比较值的大小，如果mid<target,left=mid+1/right=mid-1,重新算出mid
//选择left=mid+1/right=mid-1要看题目的要求,也有可能不是这两个表达式
//还有边界问题，思考查找的元素的边界

//Search 704. 二分查找 https://leetcode.cn/problems/binary-search/
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		num := nums[mid]

		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

//SearchInsertI 35. 搜索插入位置 https://leetcode.cn/problems/search-insert-position/
func SearchInsertI(nums []int, target int) int {
	exist, res := find(nums, target)
	if exist {
		return res
	}
	if nums[0] > target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	i := 0
	for ; i < len(nums)-1; i++ {

		if target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}
	if i == len(nums)-1 {
		return len(nums)
	}

	return 0
}

func find(nums []int, target int) (bool, int) {
	for k, v := range nums {
		if v == target {
			return true, k
		}
	}

	return false, -1
}

func SearchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

//IsPerfectSquare 367. 有效的完全平方数 https://leetcode.cn/problems/valid-perfect-square/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=42xxjpm
func IsPerfectSquare(num int) bool {
	left := 0
	right := num

	for left <= right {
		mid := (right-left)/2 + left

		if mid*mid < num {
			left = mid + 1
		} else if mid*mid > num {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

//FindTheDistanceValue 1385. 两个数组间的距离值 https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=42xxjpm
//对于一个arr1[i],在arr2数组中不存在元素arr2[j]到arr1[i]的距离小于等于 d（即!(arr1[i] - d<= arr2[j] <= arr[i] + d)） 二分法
func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0

	for i := 0; i < len(arr1); i++ {
		flag := 0
		for j := 0; j < len(arr2); j++ {
			abs := int(math.Abs(float64(arr1[i] - arr2[j])))
			if abs <= d {
				flag = 1
			}
		}
		if flag == 0 {
			res++
		}
	}

	return res
}

//MySqrt 69. x 的平方根 https://leetcode.cn/problems/sqrtx/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=42xxjpm
func MySqrt(x int) int {
	left := 0
	right := x
	ans := x
	for left <= right {
		mid := (right-left)/2 + left

		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

//NextGreatestLetter 744. 寻找比目标字母大的最小字母 https://leetcode.cn/problems/find-smallest-letter-greater-than-target/
//https://leetcode.cn/problems/find-smallest-letter-greater-than-target/comments/
func NextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1
	for left <= right {
		mid := (right-left)/2 + left

		if target < letters[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return letters[left%len(letters)]
}

//PeakIndexInMountainArray 山脉数组的峰顶索引 https://leetcode.cn/problems/peak-index-in-a-mountain-array/
func PeakIndexInMountainArray(arr []int) int {
	//res := -1
	//for i := 1; i < len(arr)-1; i++ {
	//	if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
	//		res = i
	//		return res
	//	}
	//}
	//return res

	res := -1
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			res = i
			return res
		}
	}
	return res
}

func isBadVersion(version int) bool {
	if version >= 2 {
		return true
	}
	return false
}

//FirstBadVersion 278. 第一个错误的版本 https://leetcode.cn/problems/first-bad-version/
//n=3 bad=2
func FirstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := (right-left)/2 + left

		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			right = mid //答案在区间 [left, mid] 中
		}
	}

	return left
}

//SearchRange 未解决34. 在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=42xxjpm
func SearchRange(nums []int, target int) []int {
	return nil
}

//ArrangeCoins 441. 排列硬币 https://leetcode.cn/problems/arranging-coins/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=42xxjpm
func ArrangeCoins(n int) int {
	//row := make([]int, 0)
	num := 0
	res := 0
	for i := 1; res <= n; i++ {
		res += i
		if res <= n {
			num = i
		} else {
			break
		}
	}

	return num
}

//FindKthPositive 1539. 第 k 个缺失的正整数 https://leetcode.cn/problems/kth-missing-positive-number/
//https://leetcode.cn/problems/kth-missing-positive-number/solution/di-k-ge-que-shi-de-zheng-zheng-shu-by-leetcode-sol/
func FindKthPositive(arr []int, k int) int {
	misCount := 0
	current := 1
	ptr := 0
	res := 0
	for misCount = 0; misCount < k; current++ {
		if current == arr[ptr] {
			if ptr+1 < len(arr) {
				ptr++
			}
		} else {
			misCount++
			res = current
		}
	}

	return res
}
