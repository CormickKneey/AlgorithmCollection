package Records

import (
	"math"
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

func Quick_sort(q []int, l, r int) {
	if l >= r { // 终止条件
		return
	}
	x := q[(l+r)>>1] // 确定分界点
	i, j := l-1, r+1 // 两个指针，因为do while要先自增/自减
	for i < j {      // 每次迭代
		for { // do while 语法
			i++ // 交换后指针要移动，避免没必要的交换
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j { // swap 两个元素
			q[i], q[j] = q[j], q[i]
		}
	}
	Quick_sort(q, l, j) // 递归处理左右两段
	Quick_sort(q, j+1, r)

}

/*
求n的平方根x，x精确到m位小数。例如求n精确到小数点后4位的平方根。
*/

func Sqrt(num float64,pre float64)float64{
	y := num/2.0
	low := 0.0
	up := num
	for math.Abs(y * y - num) > pre{
		// fmt.Println(count, y)
		if y*y > num{
			up = y
			y = low+(y-low)/2
		}else{
			low = y
			y = up-(up-y)/2
		}
	}
	return y
}