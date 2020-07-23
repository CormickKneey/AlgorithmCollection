package Records

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	var (
		str = "2+3*9-9+4/2"
		expected = 22
	)
	actual := Calculator(str)
	if actual != expected{
		t.Errorf("Actual : %d; expected %d", actual , expected)
	}
}


func TestSubset(t *testing.T) {
	var (
		nums = []int{1,2,3}
		//expected = [][]int{{},{1},{1,2},{1,3},{1,2,3},{2,3}}
	)
	actual := Subset2(nums)
	for _,v := range actual{
		for _,key := range v{
			fmt.Print(key)
			fmt.Print("\t")
		}
		print("\n")
	}
}


func TestReverseList(t *testing.T) {
	var (
		head = ListNode{Val: 1}
		p1 = ListNode{Val: 2}
		p2 = ListNode{Val: 3}
		p3 = ListNode{Val: 4}
	)
	head.Next = &p1
	p1.Next = &p2
	p2.Next = &p3
	p3.Next = nil
	res := ReverseList(&head)
	for res!= nil{
		fmt.Print(res.Val)
		res = res.Next
	}
}

func TestCombine(t *testing.T) {
	var(
		k = 2
		n = 5
	)
	actual := Combine(n,k)
	for _,v := range actual{
		for _,k := range v{
			fmt.Print(k)
			fmt.Print("\t")
		}
		fmt.Print("\n")
	}
}


func TestQuick_sort(t *testing.T) {
	var(
		nums = []int{2,3,5,1,18,7,9,11,87,34,23}
	)
	Quick_sort(nums,0,len(nums)-1)
	for _,v := range nums{
		fmt.Printf("\t %d",v)
	}
}



func TestCompare(t *testing.T) {

}
