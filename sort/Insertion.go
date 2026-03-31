package sort

//插入
//利用一个空数组
//将需要排序的数组插入到空数组中
//每次插入一个元素，将新的数据插入到空数组中正确的位置
//时间复杂度O(n^2)
//空间复杂度O(1)
func InsertionSort(arr []int) (res []int) {
	res = make([]int, len(arr))
	for _, v1 := range arr {
		if len(res) == 0 {
			res = append(res, v1)
			continue
		}
		for k, v2 := range res {
			if v1 < v2 {
				res = append(res[:k], append(res[k:], v1)...)
				break
			}
		}
	}

	return res
}
