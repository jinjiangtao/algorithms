package recursion

// MaxValue 递归求数组中的最大值
// 时间复杂度O(n)
// 空间复杂度O(logn)
// 每次都是求数组中数组中的最大值
func MaxValue(arr []int, start, end int) int {
	if start == end {
		return arr[start]
	}
	mind := start + ((end - start) >> 2)
	leftMax := MaxValue(arr, start, mind)
	rightMax := MaxValue(arr, mind+1, end)
	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

//递归求阶乘
func factorialNubmer(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialNubmer(n-1)
}

//递归求斐波那契数列
func fibonacciNubmer(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacciNubmer(n-1) + fibonacciNubmer(n-2)
}
