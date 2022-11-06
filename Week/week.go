package Week

//maximumSubarraySum 6230. 长度为 K 子数组中的最大和 https://leetcode.cn/contest/weekly-contest-318/problems/maximum-sum-of-distinct-subarrays-with-length-k/
//https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/solutions/1951940/hua-dong-chuang-kou-by-endlesscheng-m0gm/
//滑动窗口+map
func maximumSubarraySum(nums []int, k int) int64 {
	//preArr := make([]int, len(nums))
	//
	//preArr[0] = nums[0]
	//for i := 1; i < len(nums); i++ {
	//	preArr[i] = nums[i] + preArr[i-1]
	//
	//}
	//
	//i := k - 1
	//j := -1
	//var maxs int64
	//
	//for ; i < len(nums); i++ {
	//	var tmp int64
	//	//判断是否重复
	//	flag := 0
	//	dup := make(map[int]int, k)
	//	for l := j + 1; l < i; l++ {
	//		dup[nums[l]]++
	//		if dup[nums[l]] > 1 {
	//			flag = 1
	//		}
	//	}
	//
	//	if dup[nums[i]] > 0 || flag == 1 {
	//		j++
	//		continue
	//	} else {
	//		if j == -1 {
	//			tmp = int64(preArr[i])
	//			j++
	//		} else {
	//			tmp = int64(preArr[i] - preArr[j])
	//			j++
	//		}
	//
	//	}
	//
	//	if tmp > maxs {
	//		maxs = tmp
	//	}
	//
	//}
	//
	//return maxs
	sum := 0
	cnt := map[int]int{}

	for _, v := range nums[:k-1] {
		cnt[v]++
		sum += v
	}

	var res int64
	for i := k - 1; i < len(nums); i++ {
		//判断是否重复
		cnt[nums[i]]++
		sum += nums[i]
		if len(cnt) < k {
			x := nums[i+1-k]
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
			sum -= nums[i+1-k]
			continue
		}

		if int64(sum) > res {
			res = int64(sum)
		}
		delete(cnt, nums[i+1-k])
		sum -= nums[i+1-k]

	}

	return res
}

func applyOperations(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}

	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			size++
			continue
		}
		res = append(res, nums[i])
	}

	for i := 0; i < size; i++ {
		res = append(res, 0)

	}

	return res
}
