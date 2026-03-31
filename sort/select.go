package sort

// 选择排序
// 在数组中选择最小的元素将它交换到数组的开头，
// 每次选择一个元素，将新的数据插入到空数组中正确的位置
// 时间复杂度O(n^2)
// 空间复杂度O(1)
// 选择排序， 每次从数组中选择出最小的数组，放到新开辟的空间
// 完成从小到大的排序
//需要额外的空间来存储排序后的数组
func SelectSort(arr []int) (res []int) {
	var minValue int
	res = make([]int, len(arr))

	for k, v1 := range arr {
		minValue = v1
		for _, v2 := range arr[k+1:] {
			if v2 < minValue {
				minValue = v2
			}
		}
		res[k] = minValue
	}
	return res
}
