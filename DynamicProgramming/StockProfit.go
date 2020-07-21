package DynamicProgramming

//	买卖股票的最佳时机


/*
Leetcode 121
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成【一笔】交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
注意：你不能在买入股票前卖出股票。
 */

func MaxProfit1(prices []int)int{    // len(nums) >= 2
	/*
	dp[i][0] = max(dp[i-1][1]+prices[i],dp[i-1][0])
	dp[i][1] = max(dp[i-1][1],-prices[i])
	 */
	dp_0,dp_1 := 0,-prices[0]  // dp_0 不持有股票  dp_1持有股票
	for i:=0;i<len(prices);i++{
		dp_0 = max(dp_1+prices[i],dp_0)
		dp_1 = max(dp_1,-prices[i])
	}
	return dp_0
}

/*
Leetcode 122
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成[更多的交易]（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 */

func MaxProfit2(prices []int)int { // len(nums) >= 2
	/*
		dp[i][0] = max(dp[i-1][1]+prices[i],dp[i-1][0])
		dp[i][1] = max(dp[i-1][1],dp[i][0]-prices[i])
	*/
	dp_0,dp_1 := 0,-prices[0]  // dp_0 不持有股票  dp_1持有股票
	for i:=0;i<len(prices);i++{
		dp_0 = max(dp_1+prices[i],dp_0)
		dp_1 = max(dp_1,dp_0-prices[i])
	}
	return dp_0
}

/*
Leetcode 123
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成[两笔]交易。
注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 */

func MaxProfit3(prices []int)int{
	/*
	dp[i][2][0] = max(dp[i-1][2][0],dp[i-1][2][1]+prices[i-1])
	dp[i][2][1] = max(dp[i-1][1][0]-prices[i-1],dp[i-1][2][1])
	dp[i][1][0] = max(dp[i-1][1][0],dp[i-1][1][1]+prices[i-1])
	dp[i][1][1] = max(dp[i-1][1][1],-prices[i])
	 */
	dp_1_0,dp_2_0 := 0,0
	dp_1_1,dp_2_1 := -prices[0],-prices[0]
	for i:=0;i<len(prices);i++{
		dp_2_0 = max(dp_2_0,dp_2_1+prices[i])
		dp_2_1 = max(dp_1_0-prices[i],dp_2_1)
		dp_1_0 = max(dp_1_0,dp_1_1+prices[i])
		dp_1_1 = max(dp_1_1,-prices[i])
	}
	return dp_2_0
}

