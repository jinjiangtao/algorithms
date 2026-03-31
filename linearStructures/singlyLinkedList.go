package linearstructures

//单链表节点
type SinglyLinkedNode struct {
	data any
	next *SinglyLinkedNode
}

// SinglyLinkedList 单链表
type SinglyLinkedList struct {
	head *SinglyLinkedNode
	tail *SinglyLinkedNode
	size int
}

func (s *SinglyLinkedList) GetSize() int {
	return s.size
}

func (s *SinglyLinkedList) Add(data any) {
	node := &SinglyLinkedNode{
		data: data,
	}
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.next = node
		s.tail = node
	}
	s.size++
}
