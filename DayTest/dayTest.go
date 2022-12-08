package DayTest

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Decrypt 1652. 拆炸弹 https://leetcode.cn/problems/defuse-the-bomb/ 2022/9/24
//思路:生成一个和code一样的切片添加到后面
func Decrypt(code []int, k int) []int {
	prel := len(code)
	res := make([]int, 0)
	if k == 0 {
		for i := 0; i < prel; i++ {
			code[i] = 0
			res = append(res, 0)
		}

	}

	code1 := make([]int, prel)
	copy(code1, code)
	code = append(code, code1...)

	if k > 0 {
		for i := 0; i < prel; i++ {
			last := i + k
			sum := 0
			for j := i + 1; j <= last; j++ {
				tmp := code[j]
				sum += tmp

			}
			code[i] = sum
		}

		for i := 0; i < prel; i++ {
			res = append(res, code[i])
		}

		return res
	}

	if k < 0 {
		for i := prel; i < len(code); i++ {
			last := i + k
			sum := 0
			for j := i - 1; j >= last; j-- {
				tmp := code[j%prel]
				sum += tmp
			}
			code[i] = sum
		}
		for i := prel; i < len(code); i++ {
			res = append(res, code[i])
		}
		return res
	}

	return res
}

//优化
func Decrypt1(code []int, k int) []int {
	return []int{}
}

// RotatedDigits 788. 旋转数字 https://leetcode.cn/problems/rotated-digits/ 2022/9/25
func RotatedDigits(n int) int {
	var number = 0
	//0起始 -1 存在347不正常 2意味着2个数都是0 或1 或8
	var flag = 0
	//2，5，6，9
	for i := 0; i <= n; i++ {
		s := strconv.Itoa(i)
		for j := 0; j < len(s); j++ {
			//3 4 7
			if s[j] == 51 || s[j] == 52 || s[j] == 55 {
				flag = -1
				break
			}
			//0 1 8
			if s[j] == 48 || s[j] == 49 || s[j] == 56 {
				flag++
				continue
			}
		}
		if flag == -1 {
			flag = 0
			continue
		}
		if flag == len(s) {
			flag = 0
			continue
		}
		number++
		flag = 0
	}

	return number

}

//Interpret 1678. 设计 Goal 解析器
func Interpret(command string) string {
	replaceAll := strings.ReplaceAll(command, "()", "o")
	return strings.ReplaceAll(replaceAll, "(al)", "al")
}

func CountConsistentStrings(allowed string, words []string) int {
	dic := make(map[string]int, len(allowed))
	res := 0
	for _, v := range allowed {
		dic[string(v)]++
	}

	for _, word := range words {
		for _, s := range word {
			_, ok := dic[string(s)]
			if !ok {
				res--
				break
			}
		}
		res++
	}

	return res
}

func HalvesAreAlike(s string) bool {
	l := len(s)

	dic := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}

	cnt1, cnt2 := 0, 0

	for i := 0; i < l/2; i++ {
		_, ok := dic[s[i]]

		if ok {
			cnt1++
		}
	}

	for i := l / 2; i < l; i++ {
		_, ok := dic[s[i]]

		if ok {
			cnt2++
		}
	}

	if cnt1 != cnt2 {
		return false
	}

	return true

}

//CustomSortString 791. 自定义字符串排序 https://leetcode.cn/problems/custom-sort-string/description/
func CustomSortString(order string, s string) string {
	//arr := make([]byte, len(s))
	//for i := 0; i < len(s); i++ {
	//	arr[i] = s[i]
	//}
	//
	//sort.Slice()
	//res := ""
	//for i := 0; i < len(order); i++ {
	//	for j := 0; j < len(arr); j++ {
	//		if arr[j] == order[i] {
	//			res += string(arr[j])
	//			arr = append(arr[:j], arr[j+1:]...)
	//			j = -1
	//		}
	//	}
	//}
	//
	//res += string(arr)
	//
	//return res

	val := map[byte]int{}
	for i, c := range order {
		val[byte(c)] = i + 1
	}

	t := []byte(s)

	//go里面的比较器
	sort.Slice(t, func(i, j int) bool {
		return val[t[i]] < val[t[j]]
	})

	return string(t)
}

//IsIdealPermutation 775. 全局倒置与局部倒置 https://leetcode.cn/problems/global-and-local-inversions/
func IsIdealPermutation(nums []int) bool {
	//m1 := 0
	//m2 := 0
	//
	//// for i := 0; i < len(nums)-1; i++ {
	//// 	if nums[i] > nums[i+1] {
	//// 		m2++
	//// 	}
	//// }
	//
	//for i := 0; i < len(nums)-1; i++ {
	//	if nums[i] > nums[i+1] {
	//		m1++
	//		m2++
	//	}
	//	for j := i + 2; j < len(nums); j++ {
	//		if nums[i] > nums[j] {
	//			m1++
	//			if m1>m2{
	//				return false
	//			}
	//		}
	//	}
	//}
	//
	//
	//
	//if m1 != m2 {
	//	return false
	//}
	//
	//return true

	//只要有一个局部倒置就会有一个全局倒置，所有只要全局倒置存在一个不相邻的，则两者数量不相等，返回false
	n := len(nums)
	mins := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > mins {
			return false
		}

		mins = min(mins, nums[i])
	}

	return true

}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//NumMatchingSubseq 792. 匹配子序列的单词数 https://leetcode.cn/problems/number-of-matching-subsequences/
func NumMatchingSubseq(s string, words []string) int {
	//res := 0
	//
	//for i := 0; i < len(words); i++ {
	//	m := 0
	//	flag := 0
	//	for j := 0; j < len(words[i]); {
	//		if words[i][j] == s[m] {
	//			m++
	//			flag++
	//			j++
	//		} else {
	//			m++
	//		}
	//
	//		if m == len(s) {
	//			break
	//		}
	//	}
	//	if flag == len(words[i]) {
	//		res++
	//	}
	//}
	//
	//return res

	ans := 0
	d := [26][]string{}
	for _, w := range words {
		d[w[0]-'a'] = append(d[w[0]-'a'], w)
	}
	for _, c := range s {
		q := d[c-'a']
		d[c-'a'] = nil
		for _, t := range q {
			if len(t) == 1 {
				ans++
			} else {
				d[t[1]-'a'] = append(d[t[1]-'a'], t[1:])
			}
		}
	}
	return ans

}

func largestAltitude(gain []int) int {
	high := 0
	res := 0
	for _, h := range gain {
		high = high + h
		if high > res {
			res = high
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

//1122
func nthMagicalNumber(n int, a int, b int) int {
	num := 0
	for i := 2; ; i++ {
		if i%a == 0 || i%b == 0 {
			num++
		}
		if num == n {
			return i % (int(math.Pow(float64(10), float64(9))) + 7)
		}
	}

	return 0
}

//countBalls 1742. 盒子中小球的最大数量 https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/
func countBalls(lowLimit, highLimit int) (ans int) {
	count := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		count[sum]++
		ans = max(ans, count[sum])
	}
	return
}

//NumSubarrayBoundedMax 795. 区间子数组个数 https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/description/
func NumSubarrayBoundedMax(nums []int, left int, right int) int {

	res := 0
	l1, l2 := -1, -1

	for i, num := range nums {
		if num >= left && num <= right {
			l1 = i
		}
		if num > right {
			l2 = i
			l1 = -1
		}

		if l1 != -1 {
			res += l1 - l2
		}
	}

	return res
}

//Check 1752. 检查数组是否经排序和轮转得到 https://leetcode.cn/problems/check-if-array-is-sorted-and-rotated/
func Check(nums []int) bool {
	mins := nums[0]
	flag := 0
	n := len(nums)
	tmp := make([]int, n)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= mins && nums[i] != nums[i-1] {
			flag = i
			mins = nums[i]
		}
	}

	x := n - flag

	for i := 0; i < len(nums); i++ {
		loc := (i + x) % n
		tmp[loc] = nums[i]
	}

	for i := 1; i < len(tmp); i++ {
		if tmp[i] < tmp[i-1] {
			return false
		}
	}

	return true
}

//LargestSumOfAverages 813. 最大平均值和的分组 https://leetcode.cn/problems/largest-sum-of-averages/
func LargestSumOfAverages(nums []int, k int) float64 {
	sort.Ints(nums)
	res := 0

	if k == 1 {
		for _, num := range nums {
			res += num
		}
		return float64(res / len(nums))
	}

	for i := 0; i < k-1; i++ {
		res += nums[len(nums)-1-i]
	}
	var tmp float64
	for i := 0; i < len(nums)-k+1; i++ {
		tmp += float64(nums[i])
	}

	t := tmp / float64(len(nums)-k+1)
	return float64(res) + t

}

//NearestValidPoint 1779. 找到最近的有相同 X 或 Y 坐标的点 https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/description/
func NearestValidPoint(x int, y int, points [][]int) int {
	res := -1
	distance := 32767
	tmp := distance + 1
	for i := 0; i < len(points); i++ {

		if points[i][0] == x || points[i][1] == y {
			tmp = int(math.Abs(float64(x-points[i][0])) + math.Abs(float64(y-points[i][1])))
		}
		if tmp < distance {
			distance = tmp
			res = i
		} else if tmp == distance {
			if i < res {
				res = i
			}
		}

	}

	return res
}

//SecondHighest 1796. 字符串中第二大的数字 https://leetcode.cn/problems/second-largest-digit-in-a-string/description/
func SecondHighest(s string) int {
	res := -1
	min := '0'
	for _, v := range s {
		if v >= '0' && v <= '9' {
			if v > min {
				min = v
			}
		}
	}

	tmp := '!'
	for _, v := range s {
		if v >= '0' && v <= min-1 {
			if v > tmp {
				tmp = v
			}
		}
	}
	if tmp == '!' {
		return -1
	}

	res, _ = strconv.Atoi(string(tmp))
	return res
}

//NumDifferentIntegers 1805. 字符串中不同整数的数目 https://leetcode.cn/problems/number-of-different-integers-in-a-string/
func NumDifferentIntegers(word string) int {
	res := 0
	words := []byte(word)
	dic := make(map[string]int, 0)
	for i := 0; i < len(words); i++ {
		if words[i] >= 'a' && words[i] <= 'z' {
			words[i] = ' '
		}
	}

	for i := 0; i < len(words); i++ {
		s := ""
		flag := 0

		for i < len(words) && words[i] != ' ' {
			if words[i] == '0' && flag == 0 {
				i++
				continue
			}
			s += string(words[i])
			flag = 1
			i++
		}
		if flag == 1 {
			//num, _ := strconv.Atoi(s)
			//itoa := strconv.Itoa(num)
			_, ok := dic[s]
			if !ok {
				res++
				dic[s]++
			}
		}

	}

	return res
}

func NumDifferentIntegers1(word string) int {
	set := map[string]struct{}{}
	for _, s := range strings.FieldsFunc(word, unicode.IsLower) {
		for len(s) > 1 && s[0] == '0' {
			s = s[1:]
		}
		set[s] = struct{}{}
	}
	return len(set)
}

//SquareIsWhite 1812. 判断国际象棋棋盘中一个格子的颜色 https://leetcode.cn/problems/determine-color-of-a-chessboard-square/
func SquareIsWhite(coordinates string) bool {
	if (coordinates[0]-'a')%2 == 0 {
		if coordinates[1] == '1' || coordinates[1] == '3' || coordinates[1] == '5' || coordinates[1] == '7' {
			return false
		} else {
			return true
		}
	} else {
		if coordinates[1] == '1' || coordinates[1] == '3' || coordinates[1] == '5' || coordinates[1] == '7' {
			return true
		} else {
			return false
		}
	}
}

//CheckPowersOfThree 1780. 判断一个数字是否可以表示成三的幂的和 https://leetcode.cn/problems/check-if-number-is-a-sum-of-powers-of-three/description/
func CheckPowersOfThree(n int) bool {
	num := n
	left := 0
	for num != 0 {
		left = num % 3
		num = num / 3
		if left > 1 {
			return false
		}
	}

	return true
}
