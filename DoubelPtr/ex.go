package DoubelPtr

import "sort"

//ThreeSum 15. 三数之和 https://leetcode.cn/problems/3sum/
//思路：排序+双指针+去重 参考：https://leetcode.cn/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/
func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	for v := range res {
		res[v] = make([]int, 0)
	}
	sort.Ints(nums)
	if nums[0] > 0 {
		return [][]int{}
	}

	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	for i := 0; i < n; i++ {
		L := i + 1
		R := n - 1
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				//去重
				for L < R && nums[L] == nums[L+1] {
					L = L + 1
				}
				//去重
				for L < R && nums[R] == nums[R-1] {
					R = R - 1
				}
				L++
				R--
			}
			//防止越界
			if L < R && nums[i]+nums[L]+nums[R] > 0 {
				R--
			}
			//防止越界
			if L < R && nums[i]+nums[L]+nums[R] < 0 {
				L++
			}
		}

	}

	return res
}
