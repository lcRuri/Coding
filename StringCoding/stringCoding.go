package StringCoding

import "strings"

//LongestPalindrome 5. 最长回文子串 https://leetcode.cn/problems/longest-palindromic-substring/
//思路：中心扩散
func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}

	}

	return s[start : end+1]

}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

//MyAtoi 8. 字符串转换整数 (atoi)
//https://leetcode.cn/problems/string-to-integer-atoi/solution/go-by-intelligent-shterngeg-5rk9/
func MyAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1

	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}

	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}

	for ; i < n; i++ {
		if s[i] < 48 || s[i] > 57 {
			break
		}

		result = result*10 + int(s[i]-'0')
		if sign*result < MinInt32 {
			return MinInt32
		}
		if sign*result > MaxInt32 {
			return MaxInt32
		}
	}

	return sign * result
}

//IsValid 20. 有效的括号 https://leetcode.cn/problems/valid-parentheses/ 9/27
//思路：map+栈 建立对应关系 https://leetcode.cn/problems/valid-parentheses/solution/valid-parentheses-fu-zhu-zhan-fa-by-jin407891080/
func IsValid(s string) bool {
	//长度不为偶数必然false
	if len(s)%2 != 0 {
		return false
	}
	//初始化时为1，防止越界
	stack := make([]int32, 1)
	dic := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, v := range s {
		_, ok := dic[v]
		//如果map里面有k对应的v，则将k对应的v压入栈中
		if ok {
			stack = append(stack, v)
		}
		//如果不存在，去栈中查看栈顶的k对应的v是否和当前的v相同
		if !ok {
			//如果不相同，则不对应，返回false
			if dic[stack[len(stack)-1]] != v {
				return false
			} else { //否则删除栈顶元素，继续访问下一个v
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) != 1 {
		return false
	}
	return true
}

//IsFlipedString https://leetcode.cn/problems/string-rotation-lcci/ 9/29
//https://leetcode.cn/problems/string-rotation-lcci/solution/zi-fu-chuan-lun-zhuan-by-leetcode-soluti-kc8z/
func IsFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 && len(s2) == 0 {
		return true
	}
	return strings.Contains(s1+s1, s2)

}

//romanToInt 13. 罗马数字转整数 https://leetcode.cn/problems/roman-to-integer/
func RomanToInt(s string) int {

	dir := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	arr := make([]string, 0)
	for i := 0; i < len(s); i++ {
		arr = append(arr, string(s[i]))
	}

	sum := 0
	cur, next := 0, 0
	for i := 0; i < len(arr); i++ {
		cur = dir[arr[i]]
		if i+1 < len(arr) {
			next = dir[arr[i+1]]
		}

		if next > cur {
			sum += next - cur
			i++
		} else {
			sum += cur
		}
		next = 0

	}

	return sum
}

//minAddToMakeValid 921. 使括号有效的最少添加 https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/
func MinAddToMakeValid(s string) (ans int) {
	stack := make([]int32, 0)
	cnt := 0
	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
			continue
		}
		if v == ')' {
			if len(stack) > 0 {
				stack = append(stack[:len(stack)-1], stack[len(stack):]...)
			} else {
				cnt++
			}
		}
	}

	return cnt + len(stack)

}

//MergeAlternately 1768. 交替合并字符串 https://leetcode.cn/problems/merge-strings-alternately/
func MergeAlternately(word1 string, word2 string) string {
	//res := ""
	//
	//if len(word1) > len(word2) {
	//	i := 0
	//	for ; i < len(word2); i++ {
	//		res = res + string(word1[i]) + string(word2[i])
	//	}
	//
	//	res = res + word1[i:]
	//} else if len(word1) == len(word2) {
	//	i := 0
	//	for ; i < len(word2); i++ {
	//		res = res + string(word1[i]) + string(word2[i])
	//	}
	//} else {
	//	i := 0
	//	for ; i < len(word1); i++ {
	//		res = res + string(word1[i]) + string(word2[i])
	//	}
	//
	//	res = res + word2[i:]
	//}
	//
	//return res

	n, m := len(word1), len(word2)
	ans := make([]byte, 0, n+m)
	for i := 0; i < n || i < m; i++ {
		if i < n {
			ans = append(ans, word1[i])
		}
		if i < m {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

//reverseString 344. 反转字符串 https://leetcode.cn/problems/reverse-string/description/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=4s8s8zs
func reverseString(s []byte) {
	//使用了额外空间
	//tmp := make([]byte, 0)
	//for i := len(s) - 1; i >= 0; i-- {
	//	tmp = append(tmp, s[i])
	//}
	//copy(s, tmp)

	//不使用额外空间
	for i := 0; i < len(s)/2; i++ {
		tmp := s[len(s)-i-1]
		s[len(s)-i-1] = s[i]
		s[i] = tmp
	}
}

func ReverseWords(s string) string {
	//使用额外空间
	res := make([]byte, 0)
	start, end := 0, 0
	for i, i2 := range s {
		if string(i2) == " " {
			end = i
			for j := end - 1; j >= start; j-- {
				res = append(res, s[j])
			}
			res = append(res, s[end])
			start = end + 1
		}
	}

	for i := len(s) - 1; i >= start; i-- {
		res = append(res, s[i])
	}

	return string(res)
}
