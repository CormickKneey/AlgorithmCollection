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