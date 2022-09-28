package StringCoding

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
