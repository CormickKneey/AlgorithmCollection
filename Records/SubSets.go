package Records

/*
Leetcode78:给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
示例:
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
 */
// 算法1：回溯, 两次递归，注意传递指针，否则会导致slice扩容，结果不对
func Subset(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)
	result = append(result, item)
	generate(0, nums, &item, &result)	//[0,[1,2,3],[],[[]]]
	return result
}

func generate(i int, nums []int, item *[]int, result *[][]int) {
	if i >= len(nums) {
		return
	}
	*item = append(*item, nums[i])  //[1]
	temp := make([]int, len(*item))
	for i, v := range *item {
		temp[i] = v //temp[0] = 1
	}
	*result = append(*result, temp)
	generate(i+1, nums, item, result)
	*item = (*item)[:len(*item)-1]  //撤销选择
	generate(i+1, nums, item, result)
	return
}



func Subset2(nums []int)[][]int{
	result := make([][]int, 0)
	item := make([]int, 0)
	result = append(result,item)
	for i:=0;i<len(nums);i++{
		l := len(result)
		for j:=0;j<l;j++{
			curSet := result[j]
			curSet = append(curSet,nums[i])
			result = append(result,curSet)
		}
	}
	return result
}


/*
	输入 n，k 返回[1,n]中k个数字的组合
 */

func Combine(n,k int)[][]int{
	track := make([]int,0)
	res := make([][]int,0)
	backTrack(n,k,1,&track,&res)
	return res
}
func backTrack(n,k,start int,track *[]int,res *[][]int){
	if k == len(*track){  //到底树底部
		s := make([]int, len(*track))
		copy(s, *track)
		*res = append(*res, s)
		return
	}
	for i:=start;i<=n;i++{
		*track = append(*track,i)
		backTrack(n,k,i+1,track,res)
		*track = (*track)[:len(*track)-1]
	}
}


func Getone(n int)int{
	account := 0
	for n!=1{
		if n%2 == 0{
			n /=2
			account++
		}else if n==3{
			account += 2
			return account
		}else{
			n++
			n/=2
			account +=2
		}
	}
	return account
}
