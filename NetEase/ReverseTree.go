package NetEase

import "fmt"

type Node struct {
	Val int
	Left *Node
	Right *Node
	High int
}

func reverseTree(n,m int){
	root,_ := bulidTree(n)
	node,_ := getNode(m,root)
	fmt.Print(node.Val)
}

func bulidTree(n int) (*Node, error) {
	vals := make([]int,n)
	for i,_ := range vals{
		vals[i] = i+1
	}
	length := len(vals)
	if length <=0 {
		return nil, fmt.Errorf("empty node")
	}
	var head, tmpNode *Node
	var brotherNodeList []*Node
	for _, v := range vals{
		node := &Node{Val: v}
		if head == nil {
			head = node
			tmpNode = node
		}else{
			brotherNodeList = append(brotherNodeList, node)
			node.High = tmpNode.High + 1
			if tmpNode.Left == nil {
				tmpNode.Left = node
			} else if tmpNode.Right == nil{
				tmpNode.Right = node
				tmpNode = brotherNodeList[0]
				brotherNodeList = brotherNodeList[1:]
			}
		}
	}
	return head, nil
}

func getNode(m int,root *Node)(*Node,error){
	cur := make([]*Node,0)
	cur = append(cur,root)
	if root.Val == m{
		return root,nil
	}
	queue := []*Node{root}
	for len(queue)>0{
		lengh := len(queue)
		for _,v := range queue{
			if v.Val == m {
				return v,nil
			}
			if v.Right != nil{
				queue = append(queue,v.Right)
			}
			if v.Left != nil{
				queue = append(queue,v.Left)
			}
		}
		queue = queue[lengh:]
	}
	return nil,fmt.Errorf("no result")
}
