package leetcode

//给定多个字符串， 字符串是英文字符
//写一个函数找到字符串的最长公共前缀
//如给定 abcssj ，abcsfdf，abmm 的最大公共前缀是ab
//公共扫描法， 从每个字符串第一个个字符开始扫描，遇到一个就
//遍历每个字符串的字段
func LongestCommonPrefix(strList []string) string {

	prefix := strList[0]

	for i := 0; i < len(prefix); i++ {

		for j := 1; j < len(strList); j++ {
			if strList[j][i] != prefix[i] {
				prefix = prefix[:i]
				break
			}
		}

		if prefix == "" {
			return ""
		}
	}

	return prefix
}
