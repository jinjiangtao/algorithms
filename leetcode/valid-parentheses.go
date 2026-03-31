package leetcode

import "slices"

//有三种括号 [ { ( 判断括号是否合法
//如果括号，左右对称，成对出现就是合法的号，不是成对出现就不是合法的。
//如 [{[()]}] 是合法的
// [{] 是不合法的
// 解题思路
// 1循环这个扩号列表，如果碰见左括号就入栈， 如果碰见右括号就出栈
// 2当右括号的时候判断出栈的括号是否和当前的括号是否对称，如不对称返回false
// 3 最后判断栈是否为空，如不为空返回false

func ValidParentheses(strList []string) bool {
	//用slices 模拟栈结构
	stackList := make([]string, 0)
	leftStr := []string{"{", "[", "("}

	for _, v := range strList {
		if slices.Contains(leftStr, v) {
			stackList = append(stackList, v)
		} else {
			if len(stackList) > 0 {
				str := stackList[len(stackList)-1]
				stackList = stackList[:len(stackList)-1]
				if str != v {
					return false
				}
			} else {
				return false
			}
		}
	}

	if len(stackList) > 0 {
		return false
	}
	return true
}
