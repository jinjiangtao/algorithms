package leetcode

//动态规划
//重叠子问题， 最优子结构
//一个问题的最优解， 可以通过子问题的最优解得到
//斐波那契数列问题 从第2个数开始， 后面的每一个数都是前面两个数的和
// 0,1,2,3,5,8...
//f(n) = f(n-2) + f(n-2)
//解法1 递归解法
func GetFibonacciNumber(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	return GetFibonacciNumber(number-1) + GetFibonacciNumber(number-2)
}

//自顶像下的动态规划
//递归+记忆化
func GetFibonacciNumberMap(number int, numberMap map[int]int) int {
	if number <= 1 {
		return 1
	}
	if value, ok := numberMap[number]; ok {
		return value
	}
	numberMap[number] = GetFibonacciNumberMap(number-1, numberMap) + GetFibonacciNumberMap(number-2, numberMap)
	return numberMap[number]
}

//自下而上的动态规划
//申请一个数组记录开始只有f(0) + f(1)
//后面的数都是前面的两个数的和
//在循环的过程中，动态扩容数组的大小
//这个思想就是动态规划
func GetFibonaccBySlice(number int) int {
	if number <= 1 {
		return 1
	}
	slices := []int{0, 1}
	for i := 2; i <= number; i++ {
		slices = append(slices, slices[i-1]+slices[i-2])
	}
	return slices[number]
}
