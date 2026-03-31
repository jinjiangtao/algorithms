package sort

//快速排序
func QuickSort(arr []int) (res []int) {
	if len(arr) <= 1 {
		return arr
	}
	res = make([]int, 0)
	//选中一个中间数作为基准数
	pivod := arr[0]

	//将数组分为三部分，小于基准数的，等于基准数的，大于基准数的
	mid := make([]int, 0)
	left := make([]int, 0)
	right := make([]int, 0)

	//将数组分为三部分，小于基准数的，等于基准数的，大于基准数的
	for _, v := range arr[1:] {
		if v < pivod {
			left = append(left, v)
		} else if v == pivod {
			mid = append(mid, v)
		} else {
			right = append(right, v)
		}
	}

	//递归左边的和右边的数组
	left = QuickSort(left)
	right = QuickSort(right)

	//合并左边的数组，等于基准数的数组，右边的数组
	res = append(res, left...)
	res = append(res, mid...)
	res = append(res, right...)

	return res
}
