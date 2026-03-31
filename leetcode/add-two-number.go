package leetcode

// 两数相加
// 给定一个链表 表示一个正整数，正整数是逆序的表示的且链表的每一位只有一个节点
// 写一个函数求两个链表，且链表已相同的形式返回
// 例如  1,2,3
//
//	2,8,4
//	2,0,8
type listNode struct {
	number int
	next   *listNode
}

func AddTwoNumberList1(data1 *listNode, data2 *listNode) *listNode {

	res := &listNode{}

	cur := res
	carry := 0

	for data1 != nil || data2 != nil || carry != 0 {
		sum := carry

		//添加data 1 的 并移动指针
		if data1.next != nil {
			sum = sum + data1.number
			data1 = data1.next
		}

		//添加data 2 的 并移动指针
		if data2.next != nil {
			sum = sum + data2.number
			data2 = data2.next
		}

		carry = sum / 10

		//求当前数据的值
		//申请下一个链表的存储空间地址
		//并将当前链表的下一个节点指向新申请的链表节点地址
		cur.next = &listNode{number: sum % 10}
		cur = cur.next
	}

	return res.next
}
