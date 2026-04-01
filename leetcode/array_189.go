package leetcode

//给定一个数组， 给定K 将数组的每个元素后移动K个位置
//给定数组[1,2,3,4,5,6,7] 给定K=3 移动后变成了 [5,6,7,1,2,3,4]

//解法1
//新建一个数组，数组i 的位置上的数值就是 i = (i+k)%n
func ArrayMoveK(arr []int, k int) (res []int) {

	res = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		newi := (k + i) % len(arr)
		res[i] = arr[newi]
	}

	return res
}

//解法2
//三步反转法
//反转1 将数组反转
//从K 位置反转前面的元素
//从K 位置反转后面的元素
