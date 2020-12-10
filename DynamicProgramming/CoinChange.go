package DynamicProgramming

// 背包问题以及变体
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

type Spot struct {
	x int
	y int
}

type Spots []Spot

func (s Spots)Len()int{
	return len(s)
}

func (s Spots)Less(i,j int)bool{
	return s[i].y>s[j].y
}

func (s Spots)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}



func CoinChange(amount int,coins []int)int {
	n := len(coins)
	dp := make([]int, amount+1)    // dp[i] == true 凑出总额为i的组合数
	dp[0] = 1		// 凑出0元的组合 [0]

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {  //凑出j元时，还差着前一个硬币呢(or 更多
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}
	return dp[amount]
}

/*
Leetcode:416
分割等和子集，将非空正整数组分为两个元素和相等的子集。

Input: [1,5,11,5]
Output: true
Explanation:
[1,5,5]
[11]
*/
func CanPartition(nums []int)bool{
	sum := 0
	for _,v := range nums{
		sum += v
	}
	if sum % 2 != 0{
		return false
	}
	sum /= 2 // 背包目标数
	dp := make([]bool,sum+1) // dp[i] == true 塞到i数量的分组是可行的
	dp[0] = true
	for i:=0;i<len(nums);i++{
		for j:=sum;j>=0;j--{
			if j-nums[i] >= 0{ // 能装下or正好装下
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[sum]
}





