package leetcode

//leetcode 75 题
//给定一个数组，数组里面只包含 [1,0,2] 多个元素，数组中包含多个这样的元素
//要求对数组中的元素进行排序，0的元素都排到最前面，1 排到中间，2 排到最后
//[1,0,1,2,0,2] -> [0,0,1,1,2,2]
//解1
//暴力解法 创建3个数组， 0的放到一个数组， 1的放到一个数组，2的放到一个数组
//合并数组
//荷兰国旗问题， 拿到题目后可以用排序算法。

//解2
func ArrarySrotThree(arr []int) []int {
	res := make([]int, 0)

	zeroArr := make([]int, 0)
	oneArr := make([]int, 0)
	twoArr := make([]int, 0)

	for _, v := range res {
		switch v {
		case 0:
			zeroArr = append(zeroArr, 0)
		case 1:
			oneArr = append(oneArr, v)
		case 2:
			twoArr = append(twoArr, v)
		}
	}

	//将数据合并返回
	res = append(res, zeroArr...)
	res = append(res, oneArr...)
	res = append(res, twoArr...)

	return res
}

//三指针法
//用三个指针 low current high
//开始的是low = 0， current=0, high = len(arr)
//遍历数组
//遇到 1 位置补变， 移动current
//遇到 0 交换 low 和current 位置的值，右移low 和current
//遇到 2 交换 high 和current 的值，左移high 右移current
func ArrarySrotThree2(arr []int) []int {

	low := 0
	current := 0
	high := len(arr) - 1

	for current < high {
		v := arr[current]
		switch v {
		case 0:
			arr[low], arr[current] = arr[current], arr[low]
			low++
			current++
		case 1:
			current++
		case 2:
			arr[high], arr[current] = arr[current], arr[high]
			high--
		}
	}

	return arr
}
