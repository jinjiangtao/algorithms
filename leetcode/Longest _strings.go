package leetcode

//leetcode longest-substring-without-repeating-characters
//无重复字符串最大长度
//收入一个字符串，找出这个字符串中无重复子串的最大长度
//如  abcdeabc 的最大长度是4
// aaaaa 的最大长度是1
//解法利用滑动窗口和一个map
// 1、从头开始循环这个字符串
// 2、循环开始的时候判断字符串是否出现过，出现过记录找到right 位置
// 3、map 中记录每个字符串最后的下标位置，此时的位置是left
// 4、用left-right得到不重复字符串的最大长度
// 5、用一个变量记录这个长度，并每次和最新的长度比较得到大的那个。
// 6、返回
func GetLongestStrings(str1 string) (longest int) {

	lastIdx := make(map[rune]int)

	for k, v := range str1 {
		left := 0
		if idx, ok := lastIdx[v]; ok && idx >= left {
			left = idx
		}
		lastIdx[v] = k

		right := k

		//长度是 右边-左边的位置 算头尾所以要加1
		curLong := right - left + 1
		if curLong > longest {
			longest = curLong
		}
	}

	return
}
