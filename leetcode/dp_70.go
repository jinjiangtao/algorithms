package leetcode

//动态规划
//这个问题具有重叠子结构
//问题的最优解， 就是子问题的最优解。
//leetcode 70 题
//爬楼梯  有n 个台阶的楼梯， 我们每次走1步或2步， 问走到楼顶我们一共有多少种走法
// F(n) = F(n-1) + F(n-2)

//解法1 递归解法
func GetNumber1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return GetNumber1(n-1) + GetNumber1(n-2)
}

//解法2 利用记忆法搜索
//就是利用一个map 将已经计算得到的结构记录下来
func GetNumber2(n int, numberMap map[int]int) int {
	if value, ok := numberMap[n]; ok {
		return value
	}
	if n == 0 {
		numberMap[0] = 0
		return 0
	}
	if n == 1 {
		numberMap[1] = 1
		return 1
	}
	numberMap[n] = GetNumber2(n-1, numberMap) + GetNumber2(n-2, numberMap)
	return numberMap[n]
}

// 解法3 利用动态规划方式 将前两次的结果计算出来
func GetNumber3(n int) int {
	a := 0
	b := 1
	if n <= 1 {
		return n
	}
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return a + b
}
