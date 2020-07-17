package DynamicProgramming

import "testing"

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
