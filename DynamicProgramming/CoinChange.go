package DynamicProgramming

/*
Leetcode:518
给定不同面额的硬币和一个总金额，写出可以凑出总金额的硬币组合数.

Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation:
5 = 5
5 = 2+2+1
5 = 2+1+1+1
5 = 1+1+1+1+1
 */

func CoinChange(amount int,coins []int)int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}
	return dp[amount]
}