package leetcode

// 给定数组和目标值，找到满足目标值的数组下标
// 如 []int{1,2,3,4,5,10} 找到满足11 的下标
type Num struct {
	dataOne  int
	indexOne int
	dataTwo  int
	indexTwo int
}

// 时间复杂度O(n)
func FindMapNumber(numberList []int, target int) (list []Num) {
	dataMap := make(map[int]int)

	//先定义一个数组
	for k, v := range numberList {
		dataMap[v] = k
	}

	//找到数组中和等于特定值的一对
	for k, v := range numberList {
		needNumber := target - v
		if twoData, ok := dataMap[needNumber]; ok {
			findData := Num{
				dataOne:  v,
				indexOne: k,
				dataTwo:  twoData,
				indexTwo: dataMap[needNumber],
			}
			list = append(list, findData)
		}

	}

	return list
}

// 时间复杂度O(n平方)
func FindMapNumber2(numberList []int, target int) (list []Num) {

	for k1, v1 := range numberList {
		for k2, v2 := range numberList {
			if v1+v2 == target {
				data := Num{
					dataOne:  v1,
					indexOne: k1,
					dataTwo:  v2,
					indexTwo: k2,
				}
				list = append(list, data)
			}
		}
	}

	return list
}
