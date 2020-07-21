package Records

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



func reverseList(head *ListNode) *ListNode {
	var pre ListNode
	pre.Next = head
	cur := head
	for cur.Next != nil{
		temp := cur.Next
		cur.Next = &pre
		pre = *cur
		cur = temp
	}
	return cur
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