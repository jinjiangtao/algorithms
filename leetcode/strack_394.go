package leetcode

//字符串解码
//给定一个字符串，将字符串解码
//例如给定 aa2[b] 就变成 了 aabb
// ab2[bc] 就变 abbcbc
// ab2[2[ac]] 变成了 ab2[acac] -> abacacacac
//注意 中括号会嵌套
//输入合法不需要处理边界情况
//这个也是利用栈的结构
//循环这个字符串
//碰见普通字符串，最佳到返回字符串
//碰见数字计算重复的次数
//碰见左
func StringDecode(str string) (res string) {

	res = ""
	num := 0
	numStract := make([]int, 0)
	preStrStract := make([]string, 0)

	for _, c := range str {
		if c > '0' && c < '9' {
			num = num*10 + int(c)
		} else if c == '[' {
			//入栈
			numStract = append(numStract, num)
			preStrStract = append(preStrStract, res)
			//重置
			res = ""
			num = 0
		} else if c == ']' {

			//从栈中取出来需要拼接的字符串
			preStr := preStrStract[len(preStrStract)-1]
			numPop := numStract[len(numStract)-1]

			//按照需要拼接的次数拼接字符串
			needStr := ""
			for i := 0; i < numPop; i++ {
				needStr = needStr + res
			}
			res = preStr + needStr

			//弹出栈操作
			preStrStract = preStrStract[0 : len(preStrStract)-1]
			numStract = numStract[0 : len(numStract)-1]

		} else {
			res = res + string(c)
		}
	}

	return res
}
