package leetcode

//给定一个链表 两两交换链表的节点 生成新的链表
//如: 给定链表  1->2>3->4 交换后2->1->4->3
type Pairs struct {
	Data int
	Next *Pairs
}

//迭代解法
//工程上只交换链表的值， 不交换链表的指针
func SwapPairs(paris *Pairs) (res *Pairs) {

	cur := paris
	curNext := paris.Next

	if cur != nil && curNext != nil {
		cur.Data, curNext.Data = curNext.Data, cur.Data
		if curNext.Next != nil {
			cur = curNext.Next
		}
		curNext = curNext.Next.Next
	}

	return
}
