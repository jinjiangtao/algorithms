package stack

//给定一个括号数组 ，判断是否匹配,匹配规则就像编程代码中的左右括号是否对等出现。
//如{()[]} 匹配
//如{[(]} 不匹配
func stackUse(strs []string) bool {
	stack := NewStack()
	for _, v := range strs {
		top := stack.Pop()
		if top != nil && v == top.(string) {
			continue
		} else {
			stack.Push(top)
			stack.Push(v)
		}
	}

	if stack.len == 0 {
		return true
	}

	return false
}

//
