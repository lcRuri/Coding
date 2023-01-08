package Trunk

import "sort"

//Subsets 78. 子集 https://leetcode.cn/problems/subsets/description/?envType=study-plan&id=suan-fa-ji-chu&plan=algorithms&plan_progress=1ah3sii&languageTags=golang
func Subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {

		for _, v := range res {
			tmp := []int{}
			tmp = append(tmp, nums[i])
			tmp = append(tmp, v...)
			res = append(res, tmp)
		}

	}

	return res
}

var tmp = []int{}
var res = [][]int{}

func Arrange(arr []int) (int, [][]int) {
	res = [][]int{}
	backTrunk(len(arr), arr)
	return len(res), res

}

func backTrunk(n int, arr []int) {
	if len(tmp) == 3 {
		c := make([]int, 3)
		copy(c, tmp)
		res = append(res, c)
		//tmp = tmp[:len(tmp)-1]
		return
	}
	for i := 0; i < n; i++ {
		tmp = append(tmp, arr[i])
		backTrunk(len(arr), arr)
		//每一层节点退出后，删除最后一个值
		tmp = tmp[:len(tmp)-1]
	}
}

//PermuteUnique 47. 全排列 II https://leetcode.cn/problems/permutations-ii/description/
//组合
//利用visited判断是否被访问过
//最重要：if visited[i] == true || i > 0 && visited[i-1] == false && nums[i] == nums[i-1] {
//				continue
//			}
func PermuteUnique(nums []int) (res [][]int) {
	length := len(nums)

	perm := make([]int, 0, length)
	visited := make([]bool, length)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var backtrunk func()

	backtrunk = func() {
		if len(perm) == length {
			c := make([]int, length)
			copy(c, perm)
			res = append(res, c)
			return
		}

		for i := 0; i < len(visited); i++ {
			if visited[i] == true || i > 0 && visited[i-1] == false && nums[i] == nums[i-1] {
				continue
			}

			perm = append(perm, nums[i])
			visited[i] = true
			backtrunk()
			perm = perm[:len(perm)-1]
			visited[i] = false
		}
	}

	backtrunk()
	return
}
