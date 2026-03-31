package leetcode

//最小栈操作
//设计一个栈，支持pop 和push 操作，且有一个方法支持getMin 方法随时获取栈的最小元素

//解法 设计两个栈结构，一个正常存数据， 另外一个存当前入栈时的较小数据。
//辅助栈的核心是每一层，都记住了当时栈结构里面的最小值
//所以弹出的时候弹出最小的就行了。
type TwoStack struct {
	Stack1  []int
	Struct2 []int
}

//入栈
func (s *TwoStack) PushInt(value int) {

	s.Stack1 = append(s.Stack1, value)
	minValue := value

	if minValue > s.Struct2[len(s.Struct2)-1] {
		minValue = s.Struct2[len(s.Struct2)-1]
	}

	s.Struct2 = append(s.Struct2, minValue)

}

//出栈
func (s *TwoStack) PopInt(value int) {
	s.Stack1 = s.Stack1[0 : len(s.Stack1)-1]
	s.Struct2 = s.Struct2[0 : len(s.Struct2)-1]
}

//getMinInt
//返回最小
func (s *TwoStack) GetMinInt() int {
	return s.Struct2[len(s.Struct2)-1]
}
