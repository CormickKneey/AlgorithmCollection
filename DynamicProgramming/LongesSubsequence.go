package DynamicProgramming

/*
最长递增子序列

示例:
input:  [10,9,2,5,3,7,101,18]
output: 4
explain: [2,3,7,101]
 */
func LIS(nums []int)int{{
	res := 0
	dp := make([]int,len(nums)+1)
	for i,_ := range dp{
		dp[i] = 1
	}
	//状态转移方程:
	//if num[i] > num[j] dp[i] = max(dp[i],dp[j]+1)
	for i:=0;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[i] > nums[j]{
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
		res = max(res,dp[i])
	}
	return res
}}

func max(i,j int)int{
	if i<j{
		return j
	}
	return i
}


/*
leetcode32:最长有效率括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
*/
func LongestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i - 2] + 2
				} else {
					dp[i] = 2    //"()" dp[1] = 2
				}
			} else if i - dp[i - 1] > 0 && s[i - dp[i - 1] - 1] == '(' {
				if i - dp[i - 1] >= 2 {
					dp[i] = dp[i - 1] + dp[i - dp[i - 1] - 2] + 2   //前边还有字符串
				} else {
					dp[i] = dp[i - 1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}


/*
1. 最长回文字串
2. 快排/冒泡
3. KMP
 */

func LongestPalindrome(s string)string{
	f := func(s string,i,j int)(int,int) {
		for i >= 0 && j<len(s) && s[i] == s[j]{
			i--
			j++
		}
		return i+1,j-i-1 //开始位置 与 长度
	}
	maxlen := 0
	start := 0
	for i:=0;i<len(s);i++{
		s1,len1 := f(s,i,i)
		s2,len2 := f(s,i,i+1)
		s,len := s1,len1
		if len2>len1{
			s,len = s2,len2

		}
		if len2 > maxlen{
			maxlen = len
			start = s
		}
	}
	return s[start:start+maxlen]
}



