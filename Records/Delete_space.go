package Records

import "strconv"

/*
1.用熟悉的语言写一个函数，把字符串里面的空格全部去掉，
并返回删除的空格的个数，要求空间复杂度O(1)，
时间复杂度O(n)，函数原型：int delete_space(char *str)
 */




/*
2.编写一个函数求给定两个正整数的二进制表示中不同位数的个数。
例如给定数字8和9，8的二进制表示是1000，9的二进制表示是1001，则不同位数的个数为1
 */


func Compare(p,q int)int{
	if q < p{
		q,p = p,q
	}// q > p
	pcode := numTocode(p)
	qcode := numTocode(q)
	temp := 0
	res := 0
	for i:=0;i<len(qcode);i++{
		if temp == len(pcode) && qcode[i] == '1'{  // p 已经结束
			res++
			continue
		}
		if pcode[temp] != qcode[temp]{
			res++
		}
		temp++
	}
	return res
}

func numTocode(x int)string{
	res := ""
	for ;x > 0;x/=2{
		flag := x % 2
		res = res + strconv.Itoa(flag)
	}
	return res
}


/*
3.编写一个函数将字符串转换为整数（即atoi）

 */
func mockatoi(str string)int{
	res := 0
	temp := len(str)
	for i := 0; i < len(str); i++ {
		curNum := str[i] - '0'
		res += temp*10*int(curNum)
		temp--
	}
	return res
}


/*
设计并实现一个QQ用户（QQ用户id是一个32位整数）在线统计的API，至少包含以下功能：
用户上线登记，下线取消登记；快速判断某用户是否在线（高频调用接口）；
快速统计当前多少用户在线；提供历史最高在线数的接口；统计最近5分钟内上线的用户数
 */

type requestParams struct {
	ID string
	method string   // login logout isOnline numOnline highestOnline recentOnline
}
type responseParams struct {
	status string
	extInfo struct{
	}
}

type userInfo struct {
	ID string
	isLogin bool
	lastLogintime string
}

func QQAPi(param requestParams){
	if param.method == "login" || param.method == "logout"{

	}
}



/*
5.实现一个LRUCache，LRU即Least Recently Used算法，分析你所实现的算法的时间和空间复杂度
 */

type LRUCache struct {
	size int
	cap int
}

/*

 */






