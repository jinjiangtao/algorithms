package leetcode

//给定一个数组[1,2,0,9,0,5]
//移动后的数组编程[1,2,9,5,0,0]
//将一个数组中的0 元素都移动到数组末尾
//解法1 开辟一个 新的存储空间，将数组的非0 元素放到头部，0元素放到尾部
//时间复杂度 Olog(N)

//解法2 将所有的非0 元素找到，将所有的非0 元素，一次写入到数组0-index 的位置，
//记录index 位置
//index到len(arr)-1 的位置补零
func ArraryZreoMove(array []int) []int {
	index := 0
	//找到数组不为0 的数据放到数组0-index 的索引上
	for _, v := range array {
		if v != 0 {
			array[index] = v
			index++
		}
	}

	//将数组index 后面的索引设置为0
	for k := index + 1; k < len(array); k++ {
		array[k] = 0
	}

	//返回数组
	return array
}

//第2种解法直接交互数组中0 元素的位置
