package Bs

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
