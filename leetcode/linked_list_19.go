package leetcode

//给定一个链表，删除链表倒数第n 的元素。
//解法1
//先遍历一遍链表，获取链表的长度
//再遍历一遍链表，要删除节点的位置就是从头节点开始的第 len-n 个元素
//找到这个元素并删除
type RemoveNode struct {
	data int
	next *RemoveNode
}

func RemoveListedNode(head *RemoveNode, removeN int) *RemoveNode {

	listed := 0
	curr := head
	for curr != nil {
		listed++
		curr = head.next
	}

	var pre *RemoveNode = nil
	curr1 := head
	listRight := 0

	for curr1 != nil {
		nextNode := curr1.next

		if listRight == removeN {
			pre.next = nextNode
			return head
		}

		pre = curr1
		curr1 = nextNode
	}

	return head
}

//解法2
//利用快慢指针法
//快指针移动一步，慢指针移动N步
//当快指针移动到结束位置的时候
//慢指针移动到了删除位置的前一个
