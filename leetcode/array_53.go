package leetcode

//题目 leetcode 53 题
//给定一个数组， 找出连续和最大的子数组
//解法1 暴力解法
//kadane 算法
//curmax 已当前结尾的最大子数组和
//maxsum 全局最大子数组和
//对每个数只有2种选择
//1、加入到前面子数组
//2、重新开始计算curmax
//该算法是通过谈心策略来实现的
func GetArrayMaxValue(arr []int) (sumMax int) {

	curMax := arr[0]
	sumMax = arr[0]

	for i := 1; i < len(arr); i++ {
		if (sumMax + arr[i]) < curMax {
			curMax = arr[i]
		} else {
			curMax = curMax + arr[i]
		}
		if curMax > sumMax {
			sumMax = curMax
		}
	}

	return
}
