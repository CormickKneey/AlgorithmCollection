package Records

import (
	"strconv"
)

/*
处理四则运算
input: "2+3*9-9+4/2"
output: 22
 */


func Calculator(s string)int{
	stack := make([]int,1)  //存数字,make([]int,0)会到-1
	operator := []byte{'+'} //组成搭配 '+1'，'*10'
	num := "0"              // 直接用string转数字，[]byte 不太好处理。。。。
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num += string(s[i])
		}
		if s[i]-'0' > '9' || i == len(s)-1 {  //把屁股加上
			curNum, _ := strconv.Atoi(num)
			nextNum := stack[len(stack)-1] // 待操作数
			stack  = stack[:len(stack)-1]
			if '*' == operator[len(operator)-1] {
				stack = append(stack, nextNum*curNum)
			} else if '/' == operator[len(operator)-1] {
				stack = append(stack, nextNum/curNum)
			}
			if '+' == operator[len(operator)-1] {
				stack = append(stack, nextNum)
				stack = append(stack, curNum)
			} else if '-' == operator[len(operator)-1] {
				stack = append(stack, nextNum)
				stack = append(stack, -curNum)
			}
			operator = operator[:len(operator)-1] //去掉操作完的符号
			operator = append(operator, s[i])
			num = ""
		}
	}

	res := 0
	for _,v := range(stack) {
		res += v
	}
	return res
}
