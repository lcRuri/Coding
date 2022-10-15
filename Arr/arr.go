package Arr

import "fmt"

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

//SingleNumber 只出现一次的数字 https://leetcode.cn/leetbook/read/top-interview-questions-easy/x21ib6/
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

//Intersect 两个数组的交集 II https://leetcode.cn/problems/intersection-of-two-arrays-ii/
func Intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				res = append(res, nums2[j])
				nums2 = append(nums2[:j], nums2[j+1:]...)
				break
			}
		}
	}

	return res
}

//PlusOne 66. 加一 https://leetcode.cn/problems/plus-one/
//思路：https://leetcode.cn/problems/plus-one/solution/java-shu-xue-jie-ti-by-yhhzw/
func PlusOne(digits []int) []int {
	//res := make([]int, 0)
	//var num int64
	//for i := 0; i < len(digits); i++ {
	//	num += int64(digits[i]) * int64(math.Pow(float64(10), float64(len(digits)-i-1)))
	//}
	//num++
	//
	//s := strconv.FormatInt(int64(num), 10)
	//for i := 0; i < len(s); i++ {
	//	tmp, _ := strconv.Atoi(s[i : i+1])
	//	res = append(res, tmp)
	//
	//}
	//
	//return res

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}

	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}

//MoveZeroes 移动零 https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2ba4i/
func MoveZeroes(nums []int) {
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			size++
		}
	}

	for i := 0; i < size; i++ {
		nums = append(nums, 0)
	}

	fmt.Println(nums)
}

//IsValidSudoku 36. 有效的数独 https://leetcode.cn/problems/valid-sudoku/
func IsValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true

}
