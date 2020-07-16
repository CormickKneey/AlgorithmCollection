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
	dp := make([]int,0)
	dp[0] = 1
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