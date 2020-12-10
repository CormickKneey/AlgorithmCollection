package Records

import "strings"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}


func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur!= nil{
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}


func ReverseWords(words string)string{
	wd := strings.Split(words," ")
	res := make([]string,len(wd))
	for i:=len(wd)-1;i>=0;i--{
		res[len(wd)-i] = wd[i]
	}
	str := strings.Join(res," ")
	return str
}

func ReverseInt(num int)int{
	res := 0
	for num!=0{
		res *= 10
		res += num%10
		num /=10
	}
	return res
}



func GetRoute(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if 0 == i || 0 == j {   //再最左边or最上边--一条路直达
 				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j - 1] + dp[i - 1][j]   //左边的或者上面的
			}
		}
	}
	return dp[m - 1][n - 1]
}

