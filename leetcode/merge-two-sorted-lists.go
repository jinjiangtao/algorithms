package leetcode

//合并2个int 切片， 合并前切片是有序的，合并后切片也是有序的
//如切片 []int{1, 2, 6, 7} []int{3,8}
//合并后 []int{1,2,3,6,7,8}
//迭代法合并
//从小到大的合并
//*这个题的合并方法和归并排序的子方法是一的
func MergeTwoSortedNumberList(list1 []int, list2 []int) (resList []int) {

	resList = make([]int, len(list1)+len(list2))

	i := 0
	j := 0

	if len(list1) > i && len(list2) > j {
		if list1[i] < list2[j] {
			resList = append(resList, list1[i])
			i++
		} else {
			resList = append(resList, list2[j])
			j++
		}
	}

	//将剩余的部分合并
	if len(list1) > i {
		resList = append(resList, list1[i:]...)
	}
	if len(list2) > j {
		resList = append(resList, list2[j:]...)
	}

	return resList
}

//合并的是单链表
type Node struct {
	Number int
	Next   *Node
}

func MergeSortList(list1Head *Node, list2Head *Node) (res *Node) {
	res = &Node{}
	cur := res

	for list1Head.Next != nil && list2Head.Next != nil {

		newNode := &Node{}

		if list2Head.Number > list1Head.Number {
			newNode.Number = list1Head.Number
			list2Head = list2Head.Next
		} else {
			newNode.Number = list2Head.Number
			list1Head = list1Head.Next
		}

		cur.Next = newNode
		cur = newNode
	}

	//如果链表1 不为空将剩下的数据添加上
	if list1Head != nil {
		for list1Head.Next != nil {
			newDode := &Node{
				Number: list1Head.Number,
			}
			cur.Next = newDode
			cur = newDode
		}
	}

	//如果链表2 不为空，将剩下的数据添加上
	if list2Head != nil {
		for list2Head.Next != nil {
			newDode := &Node{
				Number: list2Head.Number,
			}
			cur.Next = newDode
			cur = newDode
		}
	}

	return res.Next
}
