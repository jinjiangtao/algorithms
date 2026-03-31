package sort

//冒泡排序
//时间复杂度O(n^2)
//空间复杂度O(1)
//从第一个元素开始，将每个元素和后面的所有元素进行比较
//如果当前元素大，就交换位置
//完成的效果就是数组的从大道小排序
func BubbleSort(arr []int) (res []int) {
	for k1, v1 := range arr {
		for k2, v2 := range arr[k1+1:] {
			if v2 > v1 {
				arr[k1], arr[k2] = v2, v1
			}
		}
	}
	return arr
}
