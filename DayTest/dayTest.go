package DayTest

import (
	"container/heap"
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

//BeautySum 1781. 所有子字符串美丽值之和 https://leetcode.cn/problems/sum-of-beauty-of-all-substrings/description/
func BeautySum(s string) int {
	left, right := 0, 1
	res := 0
	for left < len(s)-2 {
		res += count(s[left : right+1])
		right++
		if right == len(s) {
			left++
			right = left + 1
		}
	}

	return res
}

//计算一个字符串的美丽值
func count(s string) int {
	cou := [26]int{}

	for i := 0; i < len(s); i++ {
		cou[s[i]-'a']++
	}
	maxs, mins := 0, 1000
	for _, v := range cou {
		if v > 0 {
			if v > maxs {
				maxs = v
			}
			if v < mins {
				mins = v
			}
		}

	}

	return maxs - mins
}

//checkIfPangram 1832. 判断句子是否为全字母句 https://leetcode.cn/problems/check-if-the-sentence-is-pangram/description/
func checkIfPangram(sentence string) bool {
	cnt := [26]int{}

	for i := 0; i < len(sentence); i++ {
		cnt[sentence[i]-'a']++
	}

	for i := 0; i < len(cnt); i++ {
		if cnt[i] == 0 {
			return false
		}
	}

	return true
}

//递归超时
//func DistanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
//	res := []bool{}
//
//	//通过n和edgeList构建图的邻接矩阵
//	edges := make([][]int, n)
//	for i := range edges {
//		edges[i] = make([]int, n)
//	}
//	//遍历edgeList填充数据
//	for _, edge := range edgeList {
//		x, y, cost := edge[0], edge[1], edge[2]
//		if edges[x][y] != 0 {
//			if cost < edges[x][y] {
//				edges[x][y], edges[y][x] = cost, cost
//			}
//		} else {
//			edges[x][y], edges[y][x] = cost, cost
//		}
//
//	}
//	//找到路
//	for _, query := range queries {
//		st, dst := query[0], query[1]
//		cost := query[2]
//
//		flag := 0
//		visit := make([]int, n)
//		if findEdge(edges, st, dst, cost, &flag, visit) == true {
//			res = append(res, true)
//		} else {
//			res = append(res, false)
//		}
//	}
//
//	return res
//}
//
//func findEdge(edges [][]int, st int, dst int, cost int, flag *int, visit []int) bool {
//	//for k := 0; k < len(edges); k++ {
//	//first := len(edges)
//	for i := 0; i < len(edges); i++ {
//		if visit[i] == 1 {
//			i++
//			continue
//		}
//		visit[st] = 1
//		if edges[st][i] != 0 {
//			//first = i
//			if i == dst && edges[st][i] < cost {
//				*flag = 1
//				return true
//			} else if edges[st][i] < cost {
//				//st = i
//				//i = -1
//				findEdge(edges, i, dst, cost, flag, visit)
//			}
//		}
//		if *flag == 1 {
//			return true
//		}
//
//	}
//	//visit = make([]int, len(edges))
//	//if first == len(edges) {
//	//	continue
//	//} else {
//	//	visit[first] = 1
//	//}
//
//	//}
//
//	return false
//
//}
//distanceLimitedPathsExist 1697. 检查边长度限制的路径是否存在 https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/description/
//并查集
func DistanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool { return edgeList[i][2] < edgeList[j][2] })

	// 并查集模板
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		fa[find(from)] = find(to)
	}

	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	// 按照 limit 从小到大排序，方便离线
	sort.Slice(queries, func(i, j int) bool { return queries[i][2] < queries[j][2] })

	ans := make([]bool, len(queries))
	k := 0
	for _, q := range queries {
		for ; k < len(edgeList) && edgeList[k][2] < q[2]; k++ {
			merge(edgeList[k][0], edgeList[k][1])
		}
		ans[q[3]] = find(q[0]) == find(q[1])
	}
	return ans
}

//GetLucky 1945. 字符串转化后的各位数字之和 https://leetcode.cn/problems/sum-of-digits-of-string-after-convert/
func GetLucky(s string, k int) int {
	num := []int{}
	res := 0
	for _, v := range s {
		x, y := 0, 0
		tmp := int(v - 'a' + 1)
		if tmp >= 10 {
			x = tmp / 10
			y = tmp % 10
			num = append(num, x, y)
		} else {
			num = append(num, tmp)
		}

	}

	for i := 0; i < k; i++ {
		sum := 0
		for j := 0; j < len(num); j++ {
			sum += num[j]
		}
		res = sum
		num = []int{}
		for sum != 0 {
			tmp := sum % 10
			num = append(num, tmp)
			sum = sum / 10
		}

	}
	return res
}

//minElements 1785. 构成特定和需要添加的最少元素 https://leetcode.cn/problems/minimum-elements-to-add-to-form-a-given-sum/description/
func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	need := goal - sum
	if need == 0 {
		return 0
	}

	k := int(math.Abs(float64(need))) / limit
	l := int(math.Abs(float64(need))) % limit
	if k == 0 {
		return 1
	} else {
		if l == 0 {
			return k
		} else {
			return k + 1
		}

	}

}

//CanChoose 1764. 通过连接另一个数组的子数组得到一个数组 https://leetcode.cn/problems/form-array-by-concatenating-subarrays-of-another-array/description/
func CanChoose(groups [][]int, nums []int) bool {
	flag := 0
	size := 0
	for i := 0; i < len(groups); i++ {
		num := 0
		k := 0
		t := flag
		j := flag
		for ; j < len(nums); j++ {
			if nums[j] == groups[i][k] {
				num++
				k++
			} else {
				j = t
				k = 0
				num = 0
				t++
			}
			if num == len(groups[i]) {
				//if j-num+1 <= flag {
				//	return false
				//}
				flag = j + 1
				size++
				break
			}

		}
	}

	if size == len(groups) {
		return true
	} else {
		return false
	}

}

//minOperations 1775. 通过最少操作次数使数组的和相等 https://leetcode.cn/problems/equal-sum-arrays-with-minimum-number-of-operations/
//??
func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	for _, num := range nums1 {
		sum1 += num
	}
	for _, num := range nums2 {
		sum2 += num
	}
	//sort.Ints(nums1)
	//sort.Ints(nums2)
	need := int(math.Abs(float64(sum2 - sum1)))
	if need == 0 {
		return 0
	} else if (len(nums1) == 1 && nums1[0] == 6 && len(nums2) == sum2) || (len(nums2) == 1 && nums2[0] == 6 && len(nums1) == sum1) {
		return -1
	} else {
		return need / 5

	}
}

//repeatedCharacter 2351. 第一个出现两次的字母 https://leetcode.cn/problems/first-letter-to-appear-twice/description/
func repeatedCharacter(s string) byte {
	ints := make([]int, 26)

	for _, i := range s {
		ints[i-'a']++

		if ints[i-'a'] == 2 {
			return byte(i)
		}
	}

	return 0
}

func getNumberOfBacklogOrders(orders [][]int) (ans int) {
	buyOrders, sellOrders := hp{}, hp2{}
	for _, o := range orders {
		price, amount := o[0], o[1]
		if o[2] == 0 {
			for amount > 0 && len(sellOrders) > 0 && sellOrders[0].price <= price {
				if sellOrders[0].left > amount {
					sellOrders[0].left -= amount
					amount = 0
					break
				}
				amount -= heap.Pop(&sellOrders).(pair).left
			}
			if amount > 0 {
				heap.Push(&buyOrders, pair{price, amount})
			}
		} else {
			for amount > 0 && len(buyOrders) > 0 && buyOrders[0].price >= price {
				if buyOrders[0].left > amount {
					buyOrders[0].left -= amount
					amount = 0
					break
				}
				amount -= heap.Pop(&buyOrders).(pair).left
			}
			if amount > 0 {
				heap.Push(&sellOrders, pair{price, amount})
			}
		}
	}
	for _, p := range buyOrders {
		ans += p.left
	}
	for _, p := range sellOrders {
		ans += p.left
	}
	return ans % (1e9 + 7)
}

type pair struct{ price, left int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price > h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type hp2 []pair

func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

//AreNumbersAscending 2042. 检查句子中的数字是否递增 https://leetcode.cn/problems/check-if-numbers-are-ascending-in-a-sentence/
func AreNumbersAscending(s string) bool {
	num := []int{}
	split := strings.Split(s, " ")
	for _, s := range split {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		num = append(num, atoi)
	}

	for i := 1; i < len(num); i++ {
		if num[i] <= num[i-1] {
			return false
		}
	}
	return true
}

func f(big, length int) int {
	if length == 0 {
		return 0
	}
	if length <= big {
		return (2*big + 1 - length) * length / 2
	}
	return (big+1)*big/2 + length - big
}

//maxValue 1802. 有界数组中指定下标处的最大值 https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/description/
//https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/solutions/2042360/you-jie-shu-zu-zhong-zhi-ding-xia-biao-c-aav4/
func maxValue(n, index, maxSum int) int {
	return sort.Search(maxSum, func(mid int) bool {
		left := index
		right := n - index - 1
		return mid+f(mid, left)+f(mid, right) >= maxSum
	})
}

//CountEven 2180. 统计各位数字之和为偶数的整数个数 https://leetcode.cn/problems/count-integers-with-even-digit-sum/
func CountEven(num int) int {
	res := 0
	for i := 2; i <= num; i++ {
		l := 0
		cur := i
		for cur != 0 {
			tmp := cur % 10
			if tmp%2 != 0 {
				l++
			}
			cur = cur / 10
		}
		if l%2 == 0 {
			res++
		}
	}

	return res
}

//minOperations1 1658. 将 x 减到 0 的最小操作数 https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/
//双指针肯定可以 但是写出来有问题
//看了题解 换种思路 求中间连续子数组的最大值
//https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/solutions/2047253/jiang-x-jian-dao-0-de-zui-xiao-cao-zuo-s-hl7u/?languageTags=golang
func minOperations1(nums []int, x int) int {
	//if nums[0] > x && nums[len(nums)-1] > x {
	//	return -1
	//}
	//res := 0
	//l, r := 0, len(nums)-1
	//left := x
	//for l < r {
	//
	//	if nums[l] >= nums[r] && nums[l] <= x {
	//		left = x - nums[l]
	//		l++
	//		res++
	//	} else if nums[r] > nums[l] && nums[r] <= x {
	//		left = x - nums[r]
	//		r--
	//		res++
	//	}
	//
	//	if left == 0 {
	//		return res
	//	}
	//
	//}
	//
	//return res

	//超时
	//t:=[]int{}
	//sum,res:=0,0
	//for _, num := range nums {
	//	sum+=num
	//}
	//
	//if sum<x{
	//	return -1
	//}
	//
	//for i := 0; i <len(nums) ; i++ {
	//	tmp:=0
	//	for j := i; sum-tmp>x&&j<len(nums); j++ {
	//		tmp+=nums[j]
	//		res++
	//	}
	//	if sum-tmp==x{
	//		t=append(t,len(nums)-res)
	//	}
	//	res=0
	//
	//}
	//if len(t)==0{
	//	return -1
	//}
	//mins:=math.MaxInt64
	//for i := 0; i < len(t); i++ {
	//	if t[i]<mins{
	//		mins=t[i]
	//	}
	//}
	//
	//return mins

	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}

	right := 0
	lsum := 0
	rsum := sum
	ans := n + 1

	for left := -1; left < n; left++ {
		if left != -1 {
			lsum += nums[left]
		}
		for right < n && lsum+rsum > x {
			rsum -= nums[right]
			right++
		}
		if lsum+rsum == x {
			ans = min(ans, (left+1)+(n-right))
		}
	}
	if ans > n {
		return -1
	}
	return ans

}

//prefixCount 2185. 统计包含给定前缀的字符串 https://leetcode.cn/problems/counting-words-with-a-given-prefix/description/
func prefixCount(words []string, pref string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		if prefix := strings.HasPrefix(words[i], pref); prefix {
			res++
		}
	}

	return res
}

//ReinitializePermutation 1806. 还原排列的最少操作步数 https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/description/
func ReinitializePermutation(n int) int {
	res := 0
	com := make([]int, n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		com[i] = i
		arr[i] = i
	}
	prem := make([]int, n)
	for {

		copy(prem, arr)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = prem[i/2]
			}
			if i%2 == 1 {
				arr[i] = prem[n/2+(i-1)/2]
			}

		}

		res++
		flag := 1
		for j := 0; j < n; j++ {
			if arr[j] != com[j] {
				flag = 0
				break
			}
		}
		if flag == 1 {
			return res
		}
	}

	return res

}

//Evaluate 1807. 替换字符串中的括号内容 https://leetcode.cn/problems/evaluate-the-bracket-pairs-of-a-string/description/
//使用额外空间 如果没有括号 直接加入 如果有 替换后加入 返回哪个空间里面的数据即可
func Evaluate(s string, knowledge [][]string) string {
	//dic := map[string]string{}
	//for i := 0; i < len(knowledge); i++ {
	//	dic[knowledge[i][0]] = knowledge[i][1]
	//}
	//start, end := 0, 0
	//for i := 0; i < len(s); i++ {
	//
	//}
	//for i := 0; i < len(s); i++ {
	//	w := s[i]
	//	if w == '(' {
	//		start = i
	//	}
	//	if w == ')' {
	//		end = i
	//		rep := s[start+1 : end]
	//		v, ok := dic[rep]
	//		if ok {
	//			old := "(" + rep + ")"
	//			s = strings.Replace(s, old, v, 1)
	//			i = i - int(math.Abs(float64(len(v)-len(old))))
	//			start = 0
	//			end = 0
	//		} else {
	//			old := "(" + rep + ")"
	//			s = strings.Replace(s, old, "?", 1)
	//			i = i - int(math.Abs(float64(1-len(old))))
	//			start = 0
	//			end = 0
	//		}
	//
	//	}
	//}
	//
	//return s
	dict := map[string]string{}
	for _, kd := range knowledge {
		dict[kd[0]] = kd[1]
	}
	ans := &strings.Builder{}
	start := -1
	for i, c := range s {
		if c == '(' {
			start = i
		} else if c == ')' {
			if t, ok := dict[s[start+1:i]]; ok {
				ans.WriteString(t)
			} else {
				ans.WriteByte('?')
			}
			start = -1
		} else if start < 0 {
			ans.WriteRune(c)
		}
	}
	return ans.String()

}

//findingUsersActiveMinutes 1817. 查找用户活跃分钟数 https://leetcode.cn/problems/finding-the-users-active-minutes/description/?languageTags=golang
//??????????????????
func findingUsersActiveMinutes(logs [][]int, k int) (ans []int) {
	ans = make([]int, k)
	sort.Slice(logs, func(i, j int) bool {
		if logs[i][0] == logs[j][0] {
			return logs[i][1] < logs[j][1]
		}
		return logs[i][0] < logs[j][0]
	})
	id := logs[0][0]
	time := logs[0][1]
	cnt := 1
	for i := 1; i < len(logs); i++ {
		if id != logs[i][0] {
			ans[cnt-1]++
			id = logs[i][0]
			time = logs[i][1]
			cnt = 1
		} else if time != logs[i][1] {
			time = logs[i][1]
			cnt++
		}
	}
	ans[cnt-1]++
	return
}

//MinSideJumps 1824. 最少侧跳次数 https://leetcode.cn/problems/minimum-sideway-jumps/
//贪心 尽量选择没有障碍的
func MinSideJumps(obstacles []int) (ans int) {

	path := 2
	t1, t2 := 0, 0
	for i := 1; i < len(obstacles); i++ {
		//当前路上没有障碍
		if obstacles[i] != path {
			continue
		}

		//当前路上存在障碍
		ans++
		//去判断当前列上的另外两条路
		//怎么找到除了path的

		if path == 1 {
			t1 = 2
			t2 = 3
		} else if path == 2 {
			t1 = 1
			t2 = 3
		} else if path == 3 {
			t1 = 1
			t2 = 2
		}
		if obstacles[i-1] == t1 {
			path = t2
			continue
		} else if obstacles[i-1] == t2 {
			path = t1
			continue
		}
		flag := 0
		for j := i + 1; j < len(obstacles); j++ {
			//除去目前所在道路，另外两条都没有障碍
			//谁先遇到障碍，去另外一条路
			if obstacles[j] == t1 {
				path = t2
				flag = 1
				break
			} else if obstacles[j] == t2 {
				path = t1
				flag = 1
				break
			}

		}

		if flag == 0 {
			path = t1
		}
	}

	return
}

func CalculateTax(brackets [][]int, income int) (ans float64) {
	if income == 0 {
		return 0
	}
	for i := 0; i < len(brackets); i++ {
		if i == 0 {
			upper := brackets[i][0]
			percent := brackets[i][1]
			if income < upper {
				ans += float64(income * percent)
			} else {
				ans += float64(upper * percent)
			}

		} else {
			if income < brackets[i][0] && income < brackets[i-1][0] {
				return ans / 100
			}
			if income <= brackets[i][0] {
				upper := income - brackets[i-1][0]
				percent := brackets[i][1]
				ans += float64(upper * percent)
				return ans / 100
			} else {
				upper := brackets[i][0] - brackets[i-1][0]
				percent := brackets[i][1]
				ans += float64(upper * percent)
			}
		}
	}

	return ans / 100
}

//GetSmallestString 1663. 具有给定数值的最小字符串 https://leetcode.cn/problems/smallest-string-with-a-given-numeric-value/description/
//先全部初始化为a，然后从后往前用z替换，知道到某一个位置上数值小于等于26，结束，直接用那个数字替换
func GetSmallestString(n int, k int) string {
	ans := make([]byte, n)
	for i := range ans {
		ans[i] = 'a'
	}
	i, d := n-1, k-n
	for ; d > 25; i, d = i-1, d-25 {
		ans[i] = 'z'
	}
	ans[i] += byte(d)
	return string(ans)
}

//GreatestLetter 2309. 兼具大小写的最好英文字母 https://leetcode.cn/problems/greatest-english-letter-in-upper-and-lower-case/description/
func GreatestLetter(s string) string {
	word1 := make([]int, 26)
	word2 := make([]int, 26)

	for _, i := range s {
		if i >= 97 {
			word1[i-'a']++
		} else {
			word2[i+32-'a']++
		}
	}

	for i := 25; i >= 0; i-- {
		if word1[i] >= 1 && word2[i] >= 1 {
			return string(i + 'A')
		}
	}

	return ""
}

//func waysToMakeFair(nums []int) int {
//
//}
type ListNode struct {
	Val  int
	Next *ListNode
}

//mergeInBetween 1669. 合并两个链表 https://leetcode.cn/problems/merge-in-between-linked-lists/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p := list1
	q := list1
	path := 0
	//先找到第一个
	for p != nil {
		if path == a {
			q.Next = list2
			break
		}
		q = p
		p = p.Next
		path++
	}

	//找到第二个
	for p != nil {
		if path == b {
			p = p.Next
			break
		}
		p = p.Next
		path++
	}

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = p

	return list1

}
