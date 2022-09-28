package main

import (
	"coding/DP"
	"fmt"
	"go/scanner"
	"go/token"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var m sync.Mutex
var set = make(map[int]bool, 0)

func printOnce(num int, i int) {
	m.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num, i)
	}
	set[num] = true
	m.Unlock()
}

var mu = sync.Mutex{}

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func main() {
	//findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	//s := string(1)
	//fmt.Println(s)
	//url := "http://127.0.0.1:9090/add"
	//
	//param := addParam{
	//	X: 100,
	//	Y: 20,
	//}
	//
	//bytes, _ := json.Marshal(param)
	//
	//resp, _ := http.Post(url, "application/json", bytes2.NewReader(bytes))
	//defer resp.Body.Close()
	//
	//bytesResp, _ := ioutil.ReadAll(resp.Body)
	//var res addResult
	//json.Unmarshal(bytesResp, &res)
	//fmt.Println(res.Data)
	//fmt.Println(trimMean([]int{9, 7, 8, 7, 7, 8, 4, 4, 6, 8, 8, 7, 6, 8, 8, 9, 2, 6, 0, 0, 1, 10, 8, 6, 3, 3, 5, 1, 10, 9, 0, 7, 10, 0, 10, 4, 1, 10, 6, 9, 3, 6, 0, 0, 2, 7, 0, 6, 7, 2, 9, 7, 7, 3, 0, 1, 6, 1, 10, 3}))

	//fmt.Println(uniquePaths(3, 7))
	//maximumSwap(99386)

	//fmt.Println(maxLengthBetweenEqualCharacters("aa"))
	//fmt.Println(string(rune(97 + 0)))

	//obstacleGrid := [][]int{{0, 0}}
	//fmt.Println(test(obstacleGrid))
	//test(obstacleGrid)

	//test2(obstacleGrid)

	//fmt.Println(frequencySort([]int{2, 3, 1, 3, 2})
	//fmt.Println(1 << 3)
	//fmt.Println(numSquares(12))

	//fmt.Println(DP.MinimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	//fmt.Println(canFormArray([]int{37, 69, 3, 74, 46}, [][]int{{37, 69, 3, 74, 46}}))
	//fmt.Println(DP.MaxProduct([]int{0, 2}))
	//fmt.Println(rotatedDigits(20))

	//a := make([]int, 0)
	//fmt.Println(a) //[0 0]
	//a = append(a, 33)
	//fmt.Println(a)

	//fmt.Println(DayTest.Decrypt([]int{2, 4, 9, 3}, -2))
	//fmt.Println(StringCoding.LongestPalindrome("babad"))
	//fmt.Println(StringCoding.MyAtoi("-222"))
	//fmt.Println(DoubelPtr.ThreeSum([]int{0, 0, 0}))
	//fmt.Println(CheckPermutation("aa", "bb"))
	//fmt.Println(StringCoding.IsValid("){"))
	//fmt.Println(strings.Contains("(([]){})", "()"))

	//fmt.Println(strings.Replace("()", "()", "", -1))

	fmt.Println(DP.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
}

//面试题 01.02. 判定是否互为字符重排 9/27
func CheckPermutation(s1 string, s2 string) bool {
	res := true
	if len(s1) != len(s2) {
		return false
	}

	com1 := make(map[int32]int, 0)
	com2 := make(map[int32]int, 0)
	for _, v := range s1 {
		com1[v]++
	}
	for _, v := range s2 {
		com2[v]++
	}

	for _, v := range s1 {
		if com1[v] != com2[v] {
			res = false
		}
	}

	return res
}

//面试题 17.19. 消失的两个数字 https://leetcode.cn/problems/missing-two-lcci/
//思路:map
func missingTwo(nums []int) []int {
	res := make([]int, 0)

	sort.Ints(nums)
	mp := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = 1
	}

	for i := 1; i < 30000; i++ {
		if mp[i] != 1 {
			res = append(res, i)
		}

		if len(res) == 2 {
			return res
		}
	}

	return res
}

//1640. 能否连接形成数组 https://leetcode.cn/problems/check-array-formation-through-concatenation/
func canFormArray(arr []int, pieces [][]int) bool {
	m := len(pieces)
	for i := 0; i < m; i++ {
		if len(pieces[i]) > 1 {
			n := len(pieces[i])
			start := find(pieces[i][0], arr)
			if start == -1 {
				return false
			}
			for k := 0; k < n; k++ {
				if arr[start] != pieces[i][k] {
					return false
				}
				if start == len(arr)-1 && k == 0 {
					return false
				} else {
					start++
				}

			}
		}
		if len(pieces[i]) == 1 {
			location := find(pieces[i][0], arr)
			if location == -1 {
				return false
			}
		}
	}

	return true
}

func find(x int, arr []int) int {
	var location = -1
	for j := 0; j < len(arr); j++ {
		if x == arr[j] {
			location = j
			return location
		}
	}

	return location
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func frequencySort(nums []int) []int {
	numsMap := make(map[int]int, len(nums))

	for _, v := range nums {
		numsMap[v]++
	}

	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return numsMap[a] < numsMap[b] || numsMap[a] == numsMap[b]
	})

	return nums

}

func test(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) != 1 {
		return 1
	}
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) == 1 {
		return 0
	}
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
		} else {
			obstacleGrid[0][i] = 1
		}

	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = 1
		}

	}

	for j := 1; j < len(obstacleGrid); j++ {
		for i := 1; i < len(obstacleGrid[1]); i++ {
			if obstacleGrid[j][i] == 1 {
				obstacleGrid[j][i] = 0
			} else {
				obstacleGrid[j][i] = obstacleGrid[j-1][i] + obstacleGrid[j][i-1]
			}

		}
	}

	m := len(obstacleGrid) - 1
	n := len(obstacleGrid[m]) - 1
	return obstacleGrid[m][n]

}

func test2(obstacleGrid [][]int) {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return 1
}

func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	for i := 0; i < 26; i++ {
		sub := string(rune(97 + i))
		if cur := strings.LastIndex(s, sub) - strings.Index(s, sub) - 1; cur > ans {
			ans = cur
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	var tag int
	dup := make(map[int]int, 5)

	//排序
	sort.Ints(nums)
	//排序完如果第一个数大于0，无结果

	if nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {

		return [][]int{{0, 0, 0}}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			tag = i
			break
		}
	}

	//正常情况
	for i := 0; i < tag; i++ {

		for j := tag; j < len(nums); j++ {
			for k := 1; k < len(nums)-1; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					dup[i]++
					dup[j]++
					dup[k]++
					if dup[i] >= 1 && dup[k] >= 1 && dup[j] > 1 {
						continue
					}
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}

	}

	return res

}

func uniquePaths(m int, n int) int {

	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		arr[0][i] = 1
	}

	for i := 0; i < m; i++ {
		arr[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[i][j] = arr[i-1][j] + arr[i][j-1]
		}
	}

	res := arr[m-1][n-1]

	return res

}

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	sum := 0
	for _, x := range arr[n/20 : 19*n/20] {
		sum += x
	}
	return float64(sum*10) / float64(n*9)
}

func maximumSwap(num int) int {
	snum := strconv.Itoa(num)
	ints := make([]int, 0)
	pre := make([]int, len(snum))
	var fl int
	for i := 0; i < len(snum); i++ {
		t := string(snum[i])
		m, _ := strconv.Atoi(t)
		ints = append(ints, m)
	}

	copy(pre, ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	for i := 0; i < len(pre); i++ {
		if ints[i] > pre[i] {
			tmp := pre[i]

			for k := 0; k < len(pre); k++ {
				if ints[i] == pre[k] {
					fl = k

				}
			}
			pre[i] = ints[i]
			pre[fl] = tmp
			break
		}
	}

	byt := make([]string, len(pre))
	for i := 0; i < len(pre); i++ {
		byt[i] = strconv.Itoa(pre[i])
	}

	res, _ := strconv.Atoi(strings.Join(byt, ""))
	return res

}

func longestCommonPrefix(strs []string) string {

	var res string

	m := []byte(strs[0])
	size := len(strs)

	i := 1
	for {
		if len(strs) == 1 {
			return strs[0]
		}
		n := strs[i]
		similar := make([]byte, 0)
		l := len(n)
		if len(m) < len(n) {
			l = len(m)
		}
		for x := 0; x < l; x++ {
			if m[0] != n[0] {
				return ""
			}
			if m[x] != n[x] {
				break
			} else if m[x] == n[x] {
				similar = append(similar, m[x])
			}

		}

		m = similar
		tmp := m
		i++

		if i > size-1 {
			res = string(tmp)
			return res
		}

	}
}

func specialArray(nums []int) int {
	sort.Ints(nums)
	size := len(nums)
	for i := 0; i < size; i++ {
		num := nums[i]
		for x := 0; x <= num; x++ {

			n := len(nums[i:])
			if x == n {
				if i > 0 && x == nums[i-1] {
					return -1
				}
				return x
			}
		}

	}

	return -1
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	m := strconv.Itoa(x)
	res := make([]uint8, 0)
	for i := len(m) - 1; i >= 0; i-- {
		res = append(res, m[i])
	}
	tmp := string(res)
	reserve, _ := strconv.Atoi(tmp)
	fmt.Println(reserve)
	if x == reserve {
		return true
	}

	return false

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64
	var k int
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l := len(nums1)
	k = int(l / 2)

	if l%2 != 0 {
		res = float64(nums1[k])
	} else {
		a := nums1[k]
		b := nums1[k-1]
		res = float64((a + b) / 2)
	}

	return res

}

func reverse(x int) int {
	//0的情况
	if x == 0 {
		return x
	}
	//结尾为0的情况
	if x%10 == 0 && x != 0 {
		x = x / 10
	}
	//转为字符串
	str := strconv.Itoa(x)
	size := len(str)
	tag := 0
	res := make([]byte, 0)
	for i := size - 1; i >= 0; i-- {
		if str[i] == 45 {
			tag = 1
			break
		}
		res = append(res, str[i])

	}
	s := string(res)
	atoi, _ := strconv.Atoi(s)
	if tag == 1 {
		atoi = -atoi
	}

	return atoi

}

func ExampleScanner_Scan() {
	// src is the input that we want to tokenize.
	src := []byte("cos(x) + 1i*sin(x) // Euler")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	// output:
	// 1:1	IDENT	"cos"
	// 1:4	(	""
	// 1:5	IDENT	"x"
	// 1:6	)	""
	// 1:8	+	""
	// 1:10	IMAG	"1i"
	// 1:12	*	""
	// 1:13	IDENT	"sin"
	// 1:16	(	""
	// 1:17	IDENT	"x"
	// 1:18	)	""
	// 1:20	;	"\n"
	// 1:20	COMMENT	"// Euler"
}

func lengthOfLongestSubstring(s string) int {
	start, res, tmp := 0, 0, 0
	//i相当于右边界，j相当于左边界
	for i := 0; i < len(s); i++ {
		tmp = i - start + 1
		for j := start; j < i; j++ {
			if s[j] == s[i] {
				tmp = i - start
				start = j + 1
				break
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}
