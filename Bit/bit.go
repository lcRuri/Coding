package Bit

//IsPowerOfTwo 231. 2 的幂 https://leetcode.cn/problems/power-of-two/description/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs
func IsPowerOfTwo(n int) bool {
	//for i := 1; i <= n; {
	//
	//	if n == i {
	//		return true
	//	}
	//	i = i << 1
	//}
	//
	//return false

	//不使用循环
	return n > 0 && n&(n-1) == 0
}

func HammingWeight(num uint32) int {
	res := 0

	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			res++
		}
	}
	return res
}

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}
