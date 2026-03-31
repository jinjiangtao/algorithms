package string

//字母异位词
func testStrSmal(str1 string, str2 string) bool {
	str1Map := make(map[string]int)
	str2map := make(map[string]int)
	for _, v := range str1 {
		str1Map[string(v)]++
	}
	for _, v := range str2 {
		str2map[string(v)]++
	}
	for k, v := range str1Map {
		if v != str2map[k] {
			return false
		}
	}
	return true
}
