package Arr

//RemoveDuplicates 26. 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
//思路：双指针
func RemoveDuplicates(nums []int) int {

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

//Rotate 189. 轮转数组 https://leetcode.cn/problems/rotate-array/
//https://leetcode.cn/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
func Rotate(nums []int, k int) {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[(i+k)%len(nums)] = nums[i]
	}

	copy(nums, res)

}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

//ContainsDuplicate 217. 存在重复元素 https://leetcode.cn/problems/contains-duplicate/
func ContainsDuplicate(nums []int) bool {
	dic := make(map[int]int, 0)

	for _, v := range nums {
		dic[v]++
	}

	flag := 0
	for v := range dic {
		if dic[v] >= 2 {
			flag = 1
		}
	}

	if flag == 0 {
		return false
	}

	return true
}

func SingleNumber(nums []int) int {
	//使用map
	//dic := make(map[int]int, 0)
	//
	//for _, v := range nums {
	//	dic[v]++
	//}
	//
	//flag := 0
	//for v := range dic {
	//	if dic[v] < 2 {
	//		flag = v
	//	}
	//}
	//
	//return flag

	//排序
	//sort.Ints(nums)
	//i := 0
	//for ; i < len(nums)-1; i = i + 2 {
	//	if nums[i] != nums[i+1] {
	//		return nums[i]
	//	}
	//}
	//
	//return nums[i]

	//位运算
	single := 0
	for _, v := range nums {
		single ^= v
	}

	return single
}
