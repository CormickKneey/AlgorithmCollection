package DynamicProgramming

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T){
	var(
		amount = 5
		coins = []int{1,2,5}
		expected = 4
	)
	actual := CoinChange(amount,coins)
	if actual != expected{
		t.Errorf("Actual : %d; expected %d",  actual, expected)
	}
}

func TestLIS(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil{
			t.Errorf("Error!!!!!")
		}
	}()
	var(
		nums = []int{10,9,2,5,3,7,101,18}
		expected = 4
	)
	actual := LIS(nums)
	if actual != expected{
		t.Errorf("Actual : %d; expected %d", actual , expected)
	}
}

func TestCanPartition(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil{
			t.Errorf("Error!!!!!")
		}
	}()
	var(
		nums1 = []int{1,5,11,5}
		nums2 = []int{1,5,11,5,4}
		expected1 = true
		expected2 = false
	)
	actual1 := CanPartition(nums1)
	actual2 := CanPartition(nums2)
	if actual1 != expected1 || actual2 != expected2{
		t.Errorf("\nActual1 : %v; expected1 %v\n2 : %v; expected2 %v", actual1 , expected1,actual2,expected2)
	}
}

func TestLongestValidParentheses(t *testing.T) {
	var(
		str1 = "()))((()(())))"
		str2 = "))(()()))()"
		expected1 = 10
		expected2 = 6
	)
	actual1 := LongestValidParentheses(str1)
	actual2 := LongestValidParentheses(str2)
	if actual1 != expected1 || actual2 != expected2{
		t.Errorf("\nActual1 : %v; expected1 %v\n2 : %v; expected2 %v", actual1 , expected1,actual2,expected2)
	}
}

func TestMaxProfit(t *testing.T) {
	var (
		prices = []int{7,1,5,3,6,4}
		prices2 = []int{3,3,5,0,0,3,1,4}
		expected1 = 5
		expected2 = 7
		expected3 = 6
	)
	actual1 := MaxProfit1(prices)
	actual2 := MaxProfit2(prices)
	actual3 := MaxProfit3(prices2)
	if actual1 != expected1 || actual2 != expected2 || actual3 != expected3{
		t.Errorf("\nActual1 : %v; expected1 %v" +
			"\nActual2 : %v; expected2 %v" +
			"\nActual2 : %v; expected2 %v", actual1 , expected1,actual2,expected2,actual3,expected3)
	}
}

func TestLongestPalindrome(t *testing.T) {
	var(
		s = "asdcasddsaadeaaaedas"
	)
	fmt.Printf(LongestPalindrome(s))
}
