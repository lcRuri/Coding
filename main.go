package main

import (
	"bytes"
	"coding/DayTest"
	"fmt"
	"go/scanner"
	"go/token"
	"math"
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

	//mt.Println(DP.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(DP.NthUglyNumber(10))
	//fmt.Println(DP.GetKthMagicNumber(1))

	//fmt.Println(StringCoding.IsFlipedString("aba", "bab"))
	//fmt.Println(StringCoding.IsFlipedString("abcd", "acdb"))
	//setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})

	//fmt.Println(reformatNumber("123 4-567"))
	//fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
	//fmt.Println(canTransform("LXXLXRLXXL", "XLLXRXLXLX"))
	//fmt.Println(canTransform("XL", "LX"))
	//fmt.Println(canTransform("LXXXX", "XXXLX"))

	//fmt.Println(commonFactors(8, 30))
	//fmt.Println(maxSum([][]int{{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9}}))
	//fmt.Println(maxSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))

	//fmt.Println(minimizeXor(1))
	//fmt.Println(sizen(2))
	//fmt.Println(checkOnesSegment("11111111"))
	//str := "XVL"
	//fmt.Println(string(str[0]))
	//fmt.Println(StringCoding.RomanToInt("MDCXCV"))

	//fmt.Println(StringCoding.MinAddToMakeValid("()()"))

	//split := strings.Split("9001 discuss.leetcode.com", " ")
	//fmt.Println(split[0])
	//atoi, _ := strconv.Atoi(split[0])
	//fmt.Println(atoi)
	//fmt.Println(reflect.TypeOf(atoi))
	//domain := strings.Split(split[1], ".")
	//for _, v := range domain {
	//	fmt.Println(v)
	//}

	//fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))

	//fmt.Println(maxAscendingSum([]int{10, 20, 30, 40, 50}))
	//fmt.Println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
	//fmt.Println(maxAscendingSum([]int{100, 10, 1}))

	//fmt.Println(DP.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Println(DP.MaxSubArray([]int{1}))
	//fmt.Println(DP.MaxSubArray([]int{5, 4, -1, 7, 8}))

	//fmt.Println(DP.Rob([]int{1, 2}))

	//fmt.Println(advantageCount([]int{15448, 14234, 13574, 19893, 6475}, []int{14234, 6475, 19893, 15448, 13574}))
	//fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
	//fmt.Println(advantageCount([]int{718967141, 189971378, 341560426, 23521218, 339517772}, []int{967482459, 978798455, 744530040, 3454610, 940238504}))

	//fmt.Println(hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))
	//fmt.Println(hardestWorker(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}))
	//fmt.Println(hardestWorker(145, [][]int{{114, 5}, {115, 7}, {50, 9}, {105, 11}, {18, 13}, {47, 16}, {48, 18}, {39, 19}}))
	//fmt.Println(hardestWorker(70, [][]int{{36, 3}, {1, 5}, {12, 8}, {25, 9}, {53, 11}, {29, 12}, {52, 14}}))

	//fmt.Println(findArray([]int{5, 2, 0, 3, 1}))
	//fmt.Println(findArray([]int{13}))
	//fmt.Println(5 ^ 7)
	//fmt.Println(5 ^ 2)

	//fmt.Println(scoreOfParentheses("()"))
	//fmt.Println(scoreOfParentheses("(())"))
	//fmt.Println(scoreOfParentheses("(()(()))"))
	//fmt.Println(scoreOfParentheses("()(())"))
	//fmt.Println(scoreOfParentheses("()()"))
	//fmt.Println(scoreOfParentheses("((()()))"))

	//fmt.Println(areAlmostEqual("bank", "kanb"))
	//fmt.Println(areAlmostEqual("attack", "defend"))
	//fmt.Println(areAlmostEqual("kelb", "kelb"))
	//fmt.Println(areAlmostEqual("abcd", "dcba"))
	//fmt.Println(areAlmostEqual("acb", "aab"))

	//fmt.Println(Arr.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Println(DP.MaxProfitII([]int{7, 1, 5, 3, 6, 4}))

	//fmt.Println(MaxChunksToSorted([]int{4, 3, 2, 1, 0}))
	//fmt.Println(MaxChunksToSorted([]int{1, 0, 2, 3, 4}))
	//fmt.Println(MaxChunksToSorted([]int{1, 2, 0, 3}))
	//fmt.Println(MaxChunksToSorted([]int{2, 0, 1, 3}))
	//fmt.Println(MaxChunksToSorted([]int{0, 4, 5, 2, 1, 3}))

	//Arr.Rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	//fmt.Println(Arr.ContainsDuplicate([]int{3, 3}))

	//fmt.Println(Arr.SingleNumber([]int{2, 2, 1}))
	//fmt.Println(Arr.SingleNumber([]int{4, 1, 2, 1, 2}))

	//fmt.Println(Arr.Intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	//fmt.Println(Arr.Intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))

	//fmt.Println(Arr.PlusOne([]int{9}))
	//fmt.Println(buildArray([]int{1, 2}, 4))
	//fmt.Println(buildArray([]int{1, 2, 3}, 3))

	//Arr.MoveZeroes([]int{0, 0, 1})
	//fmt.Println(CountTime("?5:00"))
	//fmt.Println(CountTime("0?:0?"))
	//fmt.Println(CountTime("??:??"))

	//var v int64 = 15
	//formatInt := strconv.FormatInt(v, 2)
	//fmt.Println(formatInt[0] == '1')
	//fmt.Println(math.Pow(float64(10), float64(9)) + 7)
	//fmt.Println(productQueries(806335498, [][]int{{11, 11}, {1, 8}}))

	//fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))

	//fmt.Println(reserve(1001002))

	//fmt.Println(countDistinctIntegers([]int{1, 13, 10, 12, 31}))
	//fmt.Println(countDistinctIntegers([]int{2, 2, 2}))

	//fmt.Println(sumOfNumberAndReverse(181))
	//fmt.Println(sumOfNumberAndReverse(0))

	//fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	//fmt.Println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))

	//mik := find(1, []int{1, 3, 5, 2, 7, 5})
	//mak := find(5, []int{1, 3, 5, 2, 7, 5})
	//
	//fmt.Println(mik, mak)

	//fmt.Println(totalFruit([]int{1, 0}))
	//fmt.Println(totalFruit([]int{1, 2, 1}))
	//fmt.Println(totalFruit([]int{0, 1, 2, 1}))
	//fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
	//fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))

	//fmt.Println(Bs.Search([]int{-1, 0, 3, 5, 9, 12}, 2))
	//fmt.Println(Bs.Search([]int{-1, 0, 3, 5, 9, 12}, 3))

	//fmt.Println(atMostNGivenDigitSet([]string{"1", "4", "9"}, 1000000000))
	//fmt.Println(atMostNGivenDigitSet([]string{"5", "7", "8"}, 59))

	//fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
	//fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))

	//fmt.Println(kthGrammar(5, 9))

	//fmt.Println(Bs.SearchInsert([]int{1, 3, 5, 6}, 2))
	//fmt.Println(Bs.SearchInsert([]int{1, 3, 5, 6}, 7))
	//fmt.Println(Bs.SearchInsert([]int{1}, 2))
	//fmt.Println(Bs.SearchInsert([]int{1, 3}, 0))

	//fmt.Println(Bs.MySqrt(8))

	//fmt.Println(Bs.NextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))

	//fmt.Println(Bs.FirstBadVersion(3))

	//fmt.Println(Bs.SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(Bs.SearchRange([]int{1}, 1))
	//fmt.Println(Bs.SearchRange([]int{2, 2}, 2))

	//fmt.Println(StringCoding.MergeAlternately("abc", "pqr"))
	//fmt.Println(StringCoding.MergeAlternately("ab", "pqrs"))
	//fmt.Println(StringCoding.MergeAlternately("abcd", "pq"))
	//s := "01:15"
	//fmt.Println(strconv.Atoi(s[3:]))

	//fmt.Println(HaveConflict([]string{"10:00", "11:00"}, []string{"14:00", "15:00"}))

	//fmt.Println(subarrayGCD([]int{9, 3, 1, 2, 6, 3}, 3))

	//fmt.Println(minCost([]int{1, 3, 5, 2}, []int{2, 3, 1, 14}))
	//fmt.Println(minCost([]int{2, 2, 2, 2}, []int{2, 3, 1, 14}))

	//arr := []int{0, 1, 2, 3, 4, 5}
	//fmt.Println(arr[:1], arr[1:])

	//fmt.Println(Arr.PartitionDisjoint([]int{5, 0, 3, 8, 6}))
	//fmt.Println(Arr.PartitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
	//fmt.Println(Arr.PartitionDisjoint([]int{1, 1}))

	//fmt.Println(Bs.PeakIndexInMountainArray([]int{0, 1, 0}))
	//fmt.Println(Bs.PeakIndexInMountainArray([]int{0, 10, 5, 2}))

	//fmt.Println(Bs.ArrangeCoins(5))
	//fmt.Println(Bs.ArrangeCoins(8))
	//fmt.Println(Bs.ArrangeCoins(11))

	//fmt.Println(Bs.FindKthPositive([]int{2, 3, 4, 7, 11}, 5))
	//fmt.Println(Bs.FindKthPositive([]int{1, 2, 3, 4}, 2))
	//fmt.Println(Bs.FindKthPositive([]int{5, 6, 7, 8, 9}, 9))

	//dic := make(map[string]string, 10)
	//dic["A"] = "1"
	//dic["A"] = "1"
	//dic["B"] = "1"
	//
	//fmt.Println(dic)

	//fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))

	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(twoSum([]int{2, 3, 4}, 6))
	//fmt.Println(twoSum([]int{-1, 0}, -1))

	//fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	//fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
	//fmt.Println(sumSubarrayMins([]int{71, 55, 82, 55}))

	//fmt.Println(averageValue([]int{1, 3, 6, 10, 12, 15}))
	//fmt.Println(averageValue([]int{1, 2, 4, 7, 10}))

	//fmt.Println(mostPopularCreator([]string{"alice", "bob", "alice", "chris"}, []string{"one", "two", "three", "four"}, []int{5, 10, 5, 4}))
	//fmt.Println(mostPopularCreator([]string{"alice", "alice", "alice"}, []string{"a", "b", "c"}, []int{1, 2, 2}))

	//fmt.Println(magicalString(7))
	//fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))

	//fmt.Println(maxRepeating("ababc", "ba"))
	//fmt.Println(maxRepeating("ababc", "ab"))
	//fmt.Println(maxRepeating("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"))

	//fmt.Println(StringCoding.ReverseWords("Let's take LeetCode contest"))

	//左闭右开
	//ints := []int{0, 1, 2, 3, 4, 5, 6}
	//fmt.Println(ints[0:2])
	//fmt.Println(ints[1:2])
	//fmt.Println(ints[0:])
	//fmt.Println(ints[1:])
	//fmt.Println(ints[:3])

	//fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
	//fmt.Println(applyOperations([]int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272}))

	//fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	//fmt.Println(maximumSubarraySum([]int{9, 9, 9, 1, 2, 3}, 3))
	//fmt.Println(maximumSubarraySum([]int{1, 1, 1, 1, 1, 1}, 1))
	//fmt.Println(maximumSubarraySum([]int{5, 3, 3, 1, 1}, 3))

	//fmt.Println(DayTest.Interpret("G()(al)"))

	//fmt.Println(DayTest.CountConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))

	//fmt.Println(DayTest.HalvesAreAlike("book"))

	//fmt.Println(Dfs.FloodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	//fmt.Println(Dfs.FloodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0))

	//fmt.Println(Dfs.MaxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))

	//fmt.Println(DayTest.CustomSortString("jftiugkz", "kfiukutzjg"))

	//fmt.Println(Arr.Combine(4, 2))
	//fmt.Println(Arr.Combine(1, 1))

	//fmt.Println(Arr.Permute([]int{1, 2, 3}))

	//fmt.Println(DP.Rob1([]int{2, 7, 9, 3, 1}))

	//fmt.Println(Bit.IsPowerOfTwo(8))

	//fmt.Println(Bit.HammingWeight(uint32(11111111111111111111111111111101)))

	//fmt.Println(DP.MinimumTotal1([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))

	//fmt.Println(Bfs.UpdateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))

	//fmt.Println(Bfs.OrangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))

	//fmt.Println(DayTest.IsIdealPermutation([]int{1, 2, 0}))

	//fmt.Println(DayTest.NumMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))

	//fmt.Println(Bs.SearchRange([]int{5, 7, 7, 8}, 8))
	//fmt.Println(Bs.SearchRange([]int{1}, 1))
	//fmt.Println(Bs.SearchRange([]int{}, 0))

	//fmt.Println(Bs.Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

	//fmt.Println(Bs.SearchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	//fmt.Println(Bs.SearchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))

	//fmt.Println(Bs.FindMin([]int{4, 5, 6, 7, 0, 1, 2}))
	//fmt.Println(Bs.FindMin([]int{3, 1, 2}))

	//list := List.MakeList([]int{1, 1})
	//duplicates := List.DeleteDuplicates(list)
	//for duplicates != nil {
	//	fmt.Println(duplicates.Val)
	//	duplicates = duplicates.Next
	//}

	//fmt.Println(Arr.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))

	//fmt.Println(StringCoding.BackspaceCompare("ab##", "c#d#"))

	//fmt.Println(StringCoding.FindAnagrams("cbaebabacd", "abc"))

	//fmt.Println(Arr.MinSubArrayLen(5, []int{2, 3, 1, 1, 1, 1, 1}))
	//fmt.Println(Arr.MinSubArrayLen(6, []int{10, 2, 3}))
	//fmt.Println(Arr.MinSubArrayLen(15, []int{1, 2, 3, 4, 5}))
	//fmt.Println(Arr.MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))

	//fmt.Println(Bfs.NumIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))

	//fmt.Println(Bfs.FindCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
	//fmt.Println(Bfs.FindCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
	//fmt.Println(Bfs.FindCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	//fmt.Println(Bfs.FindCircleNum([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}))

	//fmt.Println(DayTest.NumSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))

	//fmt.Println(Trunk.Subsets([]int{1, 2, 3}))
	//fmt.Println(DayTest.Check([]int{3, 4, 5, 1, 2}))
	//fmt.Println(DayTest.Check([]int{6, 10, 6}))
	//fmt.Println(DayTest.Check([]int{7, 9, 1, 1, 1, 3}))

	//fmt.Println(DayTest.LargestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
	//fmt.Println(DayTest.LargestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4))

	//fmt.Println(DayTest.NearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))

	//fmt.Println(DayTest.SecondHighest("ck077"))
	//fmt.Println(Trunk.Arrange([]int{1, 2, 3, 4}))

	//fmt.Println(DayTest.NumDifferentIntegers1("a1b01c001"))

	//fmt.Println(DayTest.SquareIsWhite("a1"))
	//fmt.Println(DayTest.SquareIsWhite("h3"))

	//fmt.Println(DayTest.CheckPowersOfThree(11))
	//fmt.Println(DayTest.CheckPowersOfThree(21))
	//fmt.Println(DayTest.CheckPowersOfThree(91))

	fmt.Println(DayTest.BeautySum("aabcb"))
	fmt.Println(DayTest.BeautySum("aabcbaa"))

	//s := "aabcb"
	//fmt.Println(s[0:5])

}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		tmp := Min(height[l], height[r]) * (r - l)
		if tmp > res {
			res = tmp
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

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

//reachNumber 754. 到达终点数字 https://leetcode.cn/problems/reach-a-number/
//如果一直相加等于target，那么直接返回
//如果相加之后sum-target的余数为偶数，那么说明反转前面一个偶数就行，放回index
//如果余数不是偶数，一直加，知道满足是偶数返回
func reachNumber(target int) int {
	sum := 0
	index := 0
	target = int(math.Abs(float64(target)))

	for sum < target {
		index++
		sum += index
	}

	if sum == target {
		return index
	}

	for (sum-target)%2 != 0 {
		index++
		sum += index
	}

	return index
}

//maxRepeating 1668. 最大重复子字符串 https://leetcode.cn/problems/maximum-repeating-substring/
func maxRepeating(sequence string, word string) int {
	k := 0

	for i := 0; i < len(sequence); i++ {
		j := i
		m := 0
		tmp := 0
		for j < len(sequence) && m < len(word) && sequence[j] == word[m] {
			if sequence[j] == word[m] {
				j++
				m++
			} else {
				i = j - 1
				break
			}

			if m == len(word) {
				m = 0
				tmp++

			}
		}
		if tmp > k {
			k = tmp
		}
	}

	return k
}

//arrayStringsAreEqual 1662. 检查两个字符串数组是否相等 https://leetcode.cn/problems/check-if-two-string-arrays-are-equivalent/
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1, s2 := "", ""
	for _, w1 := range word1 {
		s1 += w1
	}
	for _, w2 := range word2 {
		s2 += w2
	}

	if s1 == s2 {
		return true
	}
	return false
}

//magicalString 481. 神奇字符串 https://leetcode.cn/problems/magical-string/
//https://leetcode.cn/problems/magical-string/solution/by-endlesscheng-z8o1/
func magicalString(n int) int {
	b := make([]byte, 0, n+1)
	b = append(b, 1, 2, 2)
	tmp := []byte{2}
	for i := 2; len(b) < n; i++ {
		//tmp[0] = b[len(b)-1] ^ 3
		tmp[0] ^= 3
		b = append(b, bytes.Repeat(tmp, int(b[i]))...)
	}

	return bytes.Count(b[:n], []byte{1})
}

//317
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	res := [][]string{}
	dic := make(map[string]int, 0)

	for i, creator := range creators {
		dic[creator] += views[i]
	}
	maxs := 0
	for _, v := range dic {
		if v >= maxs {
			maxs = v
		}
	}

	for i, creator := range creators {
		if dic[creator] == maxs {

			k := i
			for j := i + 1; j < len(creators); j++ {
				if creators[j] == creator && views[j] > views[k] {
					k = j
				}
			}

			res = append(res, []string{creator, ids[k]})
		}

		delete(dic, creator)

	}

	return res
}

//317
func averageValue(nums []int) int {
	res := 0
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 && nums[i]%3 == 0 {
			res += nums[i]
			size++
		}
	}
	if size == 0 {
		return 0
	}
	return res / size
}

//subArrayRanges 未完成2104. 子数组范围和 https://leetcode.cn/problems/sum-of-subarray-ranges/
//func subArrayRanges(nums []int) int64 {
//	return 0
//}

//sumSubarrayMins 907. 子数组的最小值之和 https://leetcode.cn/problems/sum-of-subarray-minimums/
func sumSubarrayMins(arr []int) (ans int) {

	////num := make([]int, 0)
	//res := 0
	//
	////暴力解法
	////k := int(math.Pow(float64(10), float64(9))) + 7
	////for i := 0; i < len(arr); i++ {
	////	mins := arr[i]
	////	for j := i; j < len(arr); j++ {
	////		mins = Min(arr[j], mins)
	////		//num = append(num, mins)
	////		res += mins
	////		if res >= k {
	////			res = res % k
	////		}
	////
	////	}
	////}
	//k := int(math.Pow(float64(10), float64(9))) + 7
	////单调栈 求每个元素的辐射范围
	////注意⚠️：在计算左边界或者右边界时将一侧设置为求解小于等于E的元素，目的是为了解决当一个子数组中有两个最小值元素时（比如[3,1,2,4,1]中有两个1），不重复且不遗漏地统计每一个子数组。
	////
	//
	//for i := 0; i < len(arr); i++ {
	//	left, right := i, i
	//
	//	tmp := arr[i]
	//	for {
	//		if left-1 >= 0 && arr[left-1] >= tmp {
	//			left--
	//		} else {
	//			break
	//		}
	//
	//	}
	//	for {
	//		if right+1 < len(arr) && arr[right+1] > tmp {
	//			right++
	//		} else {
	//			break
	//		}
	//	}
	//	res += tmp * (right - i + 1) * (i - left + 1)
	//	if res >= k {
	//		res = res % k
	//	}
	//}
	//return res

	//官方题解
	const mod int = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	monoStack := []int{}
	for i, x := range arr {
		for len(monoStack) > 0 && x <= arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && arr[i] < arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n - i
		} else {
			right[i] = monoStack[len(monoStack)-1] - i
		}
		monoStack = append(monoStack, i)
	}
	for i, x := range arr {
		ans = (ans + left[i]*right[i]*x) % mod
	}
	return

}

//twoSum
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		}
	}

	return []int{}
}

//arraySign 1822. 数组元素积的符号 https://leetcode.cn/problems/sign-of-the-product-of-an-array/
func arraySign(nums []int) int {
	product := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			product = product * 1
		}
		if nums[i] < 0 {
			product = product * (-1)
		} else if nums[i] == 0 {
			return 0
		}
	}

	if product > 0 {
		return 1
	}
	return -1
}

////20221026
func shortestSubarray(nums []int, k int) int {

	return -1
}

//20221025
func shortestBridge(grid [][]int) int {

	return 0
}

//minCost 未完成6216. 使数组相等的最小开销 https://leetcode.cn/contest/weekly-contest-316/problems/minimum-cost-to-make-array-equal/
func minCost(nums []int, cost []int) int64 {
	//maxc := cost[0]
	//loc := 0
	//for i := 1; i < len(cost); i++ {
	//	if cost[i] > maxc {
	//		maxc = cost[i]
	//		loc = i
	//	}
	//}
	//
	//toBe := nums[loc]
	//
	//var res int64 = 0
	//for i := 0; i < len(nums); i++ {
	//	abs := int(math.Abs(float64(toBe - nums[i])))
	//	res += int64(abs * cost[i])
	//}

	var res int64 = 0
	costs := []int64{}
	cpy := make([]int, len(cost))
	dic := map[int]int{}

	for i, v := range cost {
		dic[v] = i
	}

	copy(cpy, cost)
	sort.Ints(cpy)
	for i := 0; i < len(cpy)/2; i++ {
		var tmp int64 = 0
		toBe := dic[cpy[len(cpy)-i-1]]
		for j := 0; j < len(nums); j++ {
			abs := int(math.Abs(float64(nums[toBe] - nums[j])))
			tmp += int64(abs * cost[j])
		}
		costs = append(costs, tmp)
	}

	res = costs[0]
	for _, v := range costs {
		if v < res {
			res = v
		}
	}

	return res
}

//subarrayGCD 未完成 6224. 最大公因数等于 K 的子数组数目 https://leetcode.cn/contest/weekly-contest-316/problems/number-of-subarrays-with-gcd-equal-to-k/
func subarrayGCD(nums []int, k int) int {
	return 0
}

//HaveConflict 6214. 判断两个事件是否存在冲突 https://leetcode.cn/contest/weekly-contest-316/problems/determine-if-two-events-have-conflict/
func HaveConflict(event1 []string, event2 []string) bool {
	//starth1, _ := strconv.Atoi(event1[0][:2])
	//endh1, _ := strconv.Atoi(event1[1][:2])
	//starth2, _ := strconv.Atoi(event2[0][:2])
	//endh2, _ := strconv.Atoi(event2[1][:2])
	//if endh1 == starth2 || endh2 == starth1 {
	//	return true
	//}

	start1, _ := strconv.Atoi(event1[0][:2] + event1[0][3:])
	end1, _ := strconv.Atoi(event1[1][:2] + event1[1][3:])
	start2, _ := strconv.Atoi(event2[0][:2] + event2[0][3:])
	end2, _ := strconv.Atoi(event2[1][:2] + event2[1][3:])

	if (start2 >= start1 && start2 <= end1) || (end2 >= start1 && end2 <= end1) {
		return true
	}
	if (start1 >= start2 && start1 <= end2) || (end1 >= start2 && end1 <= end2) {
		return true
	}

	return false
}

//kthGrammar 779. 第K个语法符号 https://leetcode.cn/problems/k-th-symbol-in-grammar/
func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	return k&1 ^ 1 ^ kthGrammar(n-1, (k+1)/2)
}

//https://leetcode.cn/problems/k-th-symbol-in-grammar/solution/by-lcbin-2eyj/
func kthGrammarI(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k <= (1 << (n - 2)) {
		return kthGrammar(n-1, k)
	}
	return kthGrammar(n-1, k-(1<<(n-2))) ^ 1
}

func countStudents(students []int, sandwiches []int) int {
	j := 0
	i := 0
	flag := 0
	for {

		if students[i] != sandwiches[j] {
			tmp := students[i]
			students = students[i+1:]
			students = append(students, tmp)
			flag++
		}
		if students[i] == sandwiches[j] {
			sandwiches = append(sandwiches[:j], sandwiches[j+1:]...)
			students = append(students[:i], students[i+1:]...)
			flag = 0
		}

		if flag >= len(students) {
			break
		}
		if len(students) == 0 {
			return 0
		}
	}

	return len(students)
}

func atMostNGivenDigitSet(digits []string, n int) int {
	res := 0

	if n <= 10 {
		for i := 0; i < len(digits); i++ {
			num, _ := strconv.Atoi(digits[i])
			if num <= n {
				res++
			}
		}

		return res
	}

	s := ""
	for i := 0; i < len(strconv.Itoa(n)); i++ {
		s = s + digits[0]
	}
	high, _ := strconv.Atoi(s)
	if high > n {
		high = len(strconv.Itoa(n)) - 1
	} else {
		high = len(strconv.Itoa(n))
	}

	//l := len(digits)
	for i := 0; i < high; i++ {
		////res += int(math.Pow(float64(l), float64(high-i)))
		//x := 0
		//for j := 0; j < l; j++ {
		//	atoi, _ := strconv.Atoi(digits[j])
		//	if atoi*int(math.Pow(float64(10), float64(high-1))) < n {
		//		x++
		//	}
		//
		//}
		//
		//res += int(math.Pow(float64(x), float64(high-i)))
	}

	return res
}

//func totalFruit(fruits []int) (ans int) {
//	cnt := map[int]int{}
//	left := 0
//	for right, x := range fruits {
//		cnt[x]++
//		for len(cnt) > 2 {
//			y := fruits[left]
//			cnt[y]--
//			if cnt[y] == 0 {
//				delete(cnt, y)
//			}
//			left++
//		}
//		ans = max(ans, right-left+1)
//	}
//	return
//}

//错误？？？？？？？？？？？？？？？？？？？？？？
//totalFruit 904. 水果成篮 https://leetcode.cn/problems/fruit-into-baskets/
func totalFruit(fruits []int) (ans int) {
	left := 0
	cnt := map[int]int{}
	for right, x := range fruits {
		cnt[x]++
		if len(cnt) > 2 {
			y := cnt[fruits[left]]
			delete(cnt, y)
			left = right - 1

		}
		ans = max(ans, right-left+1)
	}

	return
}

//func countSubarrays(nums []int, minK int, maxK int) int64 {
//	//res:=make([][]int,0)
//	//
//	//for i:= range res {
//	//	res[i]=make([]int,0)
//	//}
//
//	var res int64
//	arr := make([]int, 0)
//	mik := find(minK, arr)
//	mak := find(maxK, arr)
//	//for i := 0; i < len(nums); i++ {
//	//	if nums[i] >= minK && nums[i] <= maxK {
//	//		arr = append(arr, nums[i])
//	//
//	//	}
//	//}
//
//	return res
//}

func sumOfNumberAndReverse(num int) bool {

	for i := 0; i <= num; i++ {
		if i+reserve(i) == num {
			fmt.Println(i, reserve(i))
			return true
		}
	}

	return false
}

func countDistinctIntegers(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		nums = append(nums, reserve(nums[i]))
	}

	dic := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		dic[nums[i]]++
	}

	return len(dic)

}

func reserve(x int) int {
	itoa := strconv.Itoa(x)
	res := make([]byte, 0)
	for i := len(itoa) - 1; i >= 0; i-- {
		res = append(res, itoa[i])
	}
	atoi, _ := strconv.Atoi(string(res))
	return atoi

}

func findMaxK(nums []int) int {

	tmp := -1

	for i := 0; i < len(nums); i++ {

		for j := i; j < len(nums); j++ {
			if nums[j]+nums[i] == 0 {
				if int(math.Abs(float64(nums[i]))) > tmp {
					tmp = int(math.Abs(float64(nums[i])))
				}
			}

		}
	}
	return tmp
}

//productQueries 6208. 有效时间的数目 https://leetcode.cn/contest/biweekly-contest-89/problems/number-of-valid-clock-times/
func productQueries(n int, queries [][]int) []int {
	powers := make([]int, 0)
	ans := make([]int, len(queries))
	for i := 0; i < len(ans); i++ {
		ans[i] = 1
	}
	format := strconv.FormatInt(int64(n), 2)
	for i := len(format) - 1; i >= 0; i-- {
		if format[i] == '1' {
			powers = append(powers, int(math.Pow(float64(2), float64(len(format)-1-i))))
		}
	}

	for i := 0; i < len(queries); i++ {
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			ans[i] *= powers[j]
			ans[i] = ans[i] % int(math.Pow(float64(10), float64(9))+7)
		}

	}

	return ans
}

//CountTime 6208. 有效时间的数目 https://leetcode.cn/contest/biweekly-contest-89/problems/number-of-valid-clock-times/
func CountTime(time string) int {
	res := 1
	if time[4] == '?' {
		res = 10
	}
	if time[3] == '?' {
		res = res * 6
	}

	if time[0] == '?' && time[1] == '?' {
		res = res * 24
	}
	if time[0] == '?' && time[1] != '?' {
		num, _ := strconv.Atoi(string(time[1]))
		if num <= 3 {
			res = res * 3
		} else {
			res = res * 2
		}
	}

	if time[0] != '?' && time[1] == '?' {
		num, _ := strconv.Atoi(string(time[0]))
		if num == 2 {
			res = res * 4
		} else {
			res = res * 10
		}
	}

	return res

}

//buildArray 1441. 用栈操作构建数组 https://leetcode.cn/problems/build-an-array-with-stack-operations/
func buildArray(target []int, n int) []string {
	res := make([]string, 0)
	j := 0
	for i := 1; i <= n; i++ {
		if i == target[j] {
			res = append(res, "Push")
			j++
			if j > len(target)-1 {
				break
			}
		} else {
			res = append(res, "Push", "Pop")
		}
	}

	return res
}

//MaxChunksToSorted 769. 最多能完成排序的块 https://leetcode.cn/problems/max-chunks-to-make-sorted/
//思路：贪心 https://leetcode.cn/problems/max-chunks-to-make-sorted/solution/zui-duo-neng-wan-cheng-pai-xu-de-kuai-by-gc4k/
//在0-i范围内的额数全部出现，就可以划分一个chunk
func MaxChunksToSorted(arr []int) (ans int) {
	mx := 0
	for i, x := range arr {
		if x > mx {
			mx = x
		}
		if mx == i {
			ans++
		}
	}
	return

}

//areAlmostEqual 1790. 仅执行一次字符串交换能否使两个字符串相等 https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal/
//思路：仅仅交换一次就可以相同意味着两个字符串只有两个位置不同，i，j记录s2两个位置的不同，当i，j都有值的时候，意味着已经有两个位置不同，再来的话就是false，最后判断是否这两个位置互换是一样的
func areAlmostEqual(s1 string, s2 string) bool {
	//if s1 == s2 {
	//	return true
	//}
	//num := 0
	//tmp := -1
	//for i := 0; i < len(s1); i++ {
	//	if s1[i] != s2[i] {
	//		if tmp != -1 {
	//			ex := s2[i]
	//			if ex != s1[tmp] || s1[i] != s2[tmp] {
	//				return false
	//			}
	//			tmp = -1
	//			num++
	//			continue
	//
	//		}
	//		tmp = i
	//	}
	//}
	//
	//if tmp != -1 || num > 1 {
	//	return false
	//}
	//return true

	i, j := -1, -1

	for idx := range s2 {
		if s1[idx] != s2[idx] {
			if i < 0 {
				i = idx
			} else if j < 0 {
				j = idx
			} else {
				return false
			}
		}
	}

	return i < 0 || j > 0 && s1[i] == s2[j] && s1[j] == s2[i]
}

//scoreOfParentheses 856. 括号的分数 https://leetcode.cn/problems/score-of-parentheses/
//func scoreOfParentheses(s string) int {
//	out := 0
//	inner := 0
//	res := 0
//
//	stack := make([]int32, 0)
//
//	for _, v := range s {
//		if v == '(' && len(stack) == 0 {
//			stack = append(stack, v)
//			out++
//			continue
//		}
//		if v == '(' && len(stack) != 0 {
//			stack = append(stack, v)
//			inner++
//			continue
//		}
//		if v == ')' {
//			stack = stack[:len(stack)-1]
//			if len(stack) == 0 && inner == 0 {
//				res += out
//				out = 0
//				inner = 0
//			}
//			if len(stack) == 0 && inner != 0 {
//				res += inner * 2
//				out = 0
//				inner = 0
//			}
//		}
//
//	}
//
//	//if out > 1 {
//	//	return out + inner*2
//	//}
//	//if out == 1 && inner != 0 {
//	//	return inner * 2
//	//}
//	//if out == 1 && inner == 0 {
//	//	return out
//	//}
//
//	return res
//}
//scoreOfParentheses 856. 括号的分数 https://leetcode.cn/problems/score-of-parentheses/solution/gua-hao-de-fen-shu-by-leetcode-solution-we6b/
func scoreOfParentheses(s string) int {
	st := []int{0}
	for _, c := range s {
		if c == '(' {
			st = append(st, 0)
		} else {
			v := st[len(st)-1]
			st = st[:len(st)-1]
			st[len(st)-1] += max(2*v, 1)
		}
	}
	return st[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//findArray 6201. 找出前缀异或的原始数组 https://leetcode.cn/contest/weekly-contest-314/problems/find-the-original-array-of-prefix-xor/
//思路：异或的交换律
//a^b = c 则 a^c=b, b^c=a
func findArray(pref []int) []int {
	arr := make([]int, len(pref))
	arr[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		arr[i] = pref[i-1] ^ pref[i]

	}
	return arr
}

//hardestWorker 6200. 处理用时最长的那个任务的员工 https://leetcode.cn/contest/weekly-contest-314/problems/the-employee-that-worked-on-the-longest-task/
func hardestWorker(n int, logs [][]int) int {
	time := logs[0][1]
	flag := logs[0][0]
	for i := 1; i < len(logs); i++ {
		if logs[i][1]-logs[i-1][1] > time {
			time = logs[i][1] - logs[i-1][1]
			flag = logs[i][0]
		}
		if logs[i][1]-logs[i-1][1] == time {
			if logs[i][0] < flag {
				flag = logs[i][0]
			}
		}

	}
	return flag
}

func advantageCount(nums1 []int, nums2 []int) []int {
	l := len(nums1)
	if l == 1 {
		return nums1
	}

	sort.Ints(nums1)
	res := make([]int, 0)

	idx1 := 0
	idx2 := 0
	for {
		if len(res) == l {
			break
		}
		if nums1[len(nums1)-1] <= nums2[idx2] {
			res = append(res, nums1[0])
			nums1 = append(nums1[:0], nums1[1:]...)
			idx2++
			continue
		}
		if len(res) == l {
			break
		}

		if nums1[idx1] > nums2[idx2] {
			res = append(res, nums1[idx1])
			nums1 = append(nums1[:idx1], nums1[idx1+1:]...)
			idx1 = 0
			idx2++
		} else {
			idx1++
		}
	}

	return res
}

//maxAscendingSum 1800. 最大升序子数组和 https://leetcode.cn/problems/maximum-ascending-subarray-sum/
func maxAscendingSum(nums []int) int {
	//tmp := nums[0]
	//res := nums[0]
	//
	//for i := 1; i < len(nums); i++ {
	//	if nums[i] > nums[i-1] {
	//		tmp += nums[i]
	//	} else {
	//		tmp = nums[i]
	//	}
	//	if tmp > res {
	//		res = tmp
	//	}
	//
	//}
	//
	//return res

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}

	max := dp[0]
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

//subdomainVisits 811. 子域名访问计数 https://leetcode.cn/problems/subdomain-visit-count/
func subdomainVisits(cpdomains []string) []string {
	dic := make(map[string]int, 0)

	for _, v := range cpdomains {
		//以空格先分割字符串
		split := strings.Split(v, " ")
		cnt, _ := strconv.Atoi(split[0])
		s := split[1]
		dic[s] += cnt

		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				dic[s[i+1:]] += cnt
			}
		}
	}

	res := make([]string, 0)
	for k, v := range dic {
		s := strconv.Itoa(v) + " " + k
		res = append(res, s)
	}

	return res
}

//1784. 检查二进制字符串字段 https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones/ 10/3
//https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones/solution/jian-cha-er-jin-zhi-zi-fu-chuan-zi-duan-b1phi/
func checkOnesSegment(s string) bool {
	//res := true
	//flag := 0
	//
	//for i := 0; i < len(s); i++ {
	//	if s[i] == 49 {
	//		if i+1 < len(s) && s[i+1] == 49 {
	//
	//			continue
	//		}
	//		flag++
	//	}
	//
	//}
	//if flag > 1 {
	//	res = false
	//	return res
	//}
	//
	//return res
	return !strings.Contains(s, "01")
}

//minimizeXor 6194. 最小 XOR https://leetcode.cn/contest/weekly-contest-313/problems/minimize-xor/
//先求出num2的二进制中1的个数
//将1个数从小到大排列满足和num1异或最小
func minimizeXor(num1 int, num2 int) int {

	size2, _ := sizen(num2)
	_, res := sizen(num1)

	count := 0
	for i := 1; i <= num2; i++ {
		n, com := sizen(i)
		if n != size2 {
			i++
			continue
		}

		//com=rev(com)
		//res=rev(res)
		mlen := len(com)
		f := 0
		if len(res) > mlen {
			mlen = len(res)
			f = 1
		}

		if f == 0 {
			for k := 0; k < mlen-len(res); k++ {
				res = append(res, 0)
			}
		}

		if f == 1 {
			for k := 0; k < mlen-len(com); k++ {
				com = append(com, 0)
			}
		}

		for j := 0; j < mlen; j++ {
			tmp := 0
			if com[j] != res[j] {
				tmp += 1 * 2 &^ (i - 1)
			} else {
				tmp += 0 * 2 &^ (i - 1)
			}
			if tmp > count {
				count = tmp
			}
		}

	}

	return count
}

func rev(slice []int) []int {
	fmt.Println(slice)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func sizen(num2 int) (size1 int, res []int) {
	size1 = 0
	res = make([]int, 0)
	//先求出num2的二进制中1的个数
	tmp := num2
	for tmp/2 != 0 {
		x := tmp % 2
		if x == 1 {
			size1++
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
		tmp = tmp / 2
	}
	if tmp%2 == 1 {
		size1++
		res = append(res, 1)
	}

	return size1, res
}

//maxSum 6193. 沙漏的最大总和 https://leetcode.cn/contest/weekly-contest-313/problems/maximum-sum-of-an-hourglass/
func maxSum(grid [][]int) int {
	maxRes := 0

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m && i+2 < m; i++ {
		for j := 0; j < n && j+2 < n; j++ {
			tmp := grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if tmp > maxRes {
				maxRes = tmp
			}
		}
	}

	return maxRes

}

//commonFactors 6192. 公因子的数目 https://leetcode.cn/contest/weekly-contest-313/problems/number-of-common-factors/
func commonFactors(a int, b int) int {
	l := b
	flag := 1

	if a < b {
		l = a
	}

	for i := 2; i <= l; i++ {
		if a%i == 0 && b%i == 0 {
			flag++
		}
	}

	return flag
}

//canTransform  777. 在LR字符串中交换相邻字符 https://leetcode.cn/problems/swap-adjacent-in-lr-string/ 10/2
func canTransform(start string, end string) bool {
	l := len(start)
	if l == 1 {
		if start[0] != end[0] {
			return false
		}
	}
	if l == 2 && (!strings.Contains(start, "X") || !strings.Contains(end, "X")) {
		return false
	}
	flag := true

	for i := 0; i < l; i++ {
		if start[i] != end[i] {
			if i == l-1 {
				return false
			}
			if start[i+1] != end[i] || start[i] != end[i+1] {
				flag = false
			}
			if start[i] == 88 && start[i+1] == 76 {
				i++
			} else if start[i] == 82 && start[i+1] == 88 {
				i++
			} else {
				return false
			}
		}
	}

	return flag
}

//reformatNumber https://leetcode.cn/problems/reformat-phone-number/ 10/1
func reformatNumber(number string) string {

	number1 := strings.ReplaceAll(number, "-", "")
	number2 := strings.ReplaceAll(number1, " ", "")

	if len(number2) == 2 {
		return number2
	}

	if len(number2) == 3 {
		return number2
	}

	if len(number2) == 4 {
		return number2[:2] + "-" + number2[2:]
	}

	res := make([]string, 0)
	flag := 0

	for i := 0; i < len(number2)-4; i = i + 3 {
		res = append(res, number2[i:i+3])
		flag = i + 3
	}

	if len(number2)-flag == 2 || len(number2)-flag == 3 {
		res = append(res, number2[flag:])
	}

	if len(number2)-flag == 4 {
		res = append(res, number2[flag:flag+2], number2[flag+2:])

	}
	s := res[0]

	for i := 1; i < len(res); i++ {
		s += "-" + res[i]
	}

	return s
}

//setZeroes 面试题 01.08. 零矩阵 https://leetcode.cn/problems/zero-matrix-lcci/ 9/30
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	//
	//matrix1 := make([][]int, m)
	//for  v := range matrix1 {
	//	matrix1[v]=make([]int,n)
	//}
	//
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		matrix1[i][j]=1
	//	}
	//}

	h := make(map[int]bool, 0)
	l := make(map[int]bool, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				h[i] = true
				l[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		if h[i] == true {
			seth0(matrix[i])
		}
	}

	for i := 0; i < n; i++ {
		if l[i] == true {
			setl0(matrix, i)
		}
	}

}
func seth0(m []int) {
	for i := 0; i < len(m); i++ {
		m[i] = 0
	}
}
func setl0(mat [][]int, x int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if j == x {
				mat[i][j] = 0
			}
		}
	}
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
