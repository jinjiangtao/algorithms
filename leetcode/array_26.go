package leetcode

//题目移除已经排好序数组中的重复的元素
//如 [1,2,2,3,3,4] 变成了[1,2,3,4]
//用快慢指针法
//慢指针负责一步一步的往后移动
//快指针负责扫描后面的元素， 如果后面的元素和当前元素不同，移动慢指针

func RomoveSameNumber(arr []int) []int {
	slow := 0

	for fast := 1; fast < len(arr); fast++ {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
	}

	//得到不重复的数组
	arr = arr[0:slow]
	return arr
}
