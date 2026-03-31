package linearstructures

//反转单链表
type ListNode struct {
	data int
	next *ListNode
}

//迭代法反转
//从头到尾循环链表，将链表的每一个节点指向前驱节点
//非递归反转单链表
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	if curr != nil {
		nextNode := curr.next
		curr.next = prev

		//前驱节点移动到当前节点
		prev = curr
		//当前节点移动到下一个节点
		curr = nextNode

	}

	return prev
}

//递归法反转
func ReverseListDigui(head *ListNode) *ListNode {

	//递归的终止条件
	if head == nil || head.next == nil {
		return head
	}

	//递归反转单链表
	newHead := ReverseListDigui(head.next)

	//反转单链表节点
	head.next.next = head
	head.next = nil

	return newHead
}

//链表的排序
//将一个无序链表排序成一个有序链表
//排序算法的时间复杂度 Olog(n)
//满足Olog(N) 算法的排序有快速排序和归并排序，这里用归并排序
func SortList(head *ListNode) *ListNode {

	//递归的结束条件
	if head == nil || head.next == nil {
		return head
	}

	//拆分左边和右边的列表
	leftHead := head
	var rightHead *ListNode = nil

	//将链表添加到左右链表中，用快慢指针的方法
	curr := head
	currNextNext := head.next.next
	if curr != nil && currNextNext != nil {
		//移动快慢指针
		rightHead = curr
		curr = curr.next
		currNextNext = currNextNext.next.next
	}

	//合并
	returnList := MergeSortList(leftHead, rightHead)
	return returnList
}

//归并排序的合并
func MergeSortList(leftList *ListNode, rightList *ListNode) *ListNode {

	res := &ListNode{}
	curr := res

	for leftList != nil && rightList != nil {
		newData := &ListNode{}
		if leftList.data < rightList.data {
			newData.data = leftList.data
			leftList = leftList.next
		} else {
			newData.data = rightList.data
			rightList = rightList.next
		}
		curr.next = newData
		curr = newData
	}

	if leftList.next != nil {
		for leftList.next != nil {
			newData := &ListNode{
				data: leftList.data,
			}
			curr.next = newData
			curr = newData
		}
	}

	if rightList.next != nil {
		for rightList.next != nil {
			newData := &ListNode{
				data: rightList.data,
			}
			curr.next = newData
			curr = newData
		}
	}

	return res.next
}

//判断一个链表是否回文
//如 链表 1 2 1  和 1 3 3 1 是回文链表
// 链表 1 2 3 不是回文链表
// 解法1 ·借助数组， 利用数组任意节点可以通过索引访问的特性
// 从数组的前和数组后面两个方向移动索引指针，比较得到的数值是否相等
// 如果有任何不相等的就返回false
func CheckListPalindrome(head *ListNode) bool {

	//遍历链表，构造数组
	arrNode := make([]*ListNode, 0)
	if head != nil {
		arrNode = append(arrNode, head)
		head = head.next
	}

	//比较数组 左 右对称的两个元素值是否相等， 不相等返回false
	arrLen := len(arrNode)
	for i := 0; i < arrLen/2; i++ {
		leftData := arrNode[i].data
		rightData := arrNode[arrLen-i].data
		if leftData != rightData {
			return false
		}
	}

	return true
}
