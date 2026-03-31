package sort

func MergeSort(arr []int) (res []int) {
	if len(arr) <= 1 {
		return arr
	}
	res = make([]int, 0)

	//分
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	//治
	left = MergeSort(left)
	right = MergeSort(right)

	return res
}

// 合并两个有序数组
// 数组left 的所有值一定是小于数组right 的所有值的。
func Merge(left, right []int) (res []int) {
	res = make([]int, 0, len(left)+len(right))

	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	// 合并剩余的元素
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
