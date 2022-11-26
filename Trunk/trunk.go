package Trunk

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
