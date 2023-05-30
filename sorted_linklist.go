package list

type SortedLinkList[T any] struct {
	LinkList[T]
}

func NewSortedLinkList[T any]() *SortedLinkList[T] {
	return &SortedLinkList[T]{}
}

func (list *SortedLinkList[T]) Add(currVal T) {
	list.Lock()
	node := &Node[T]{Val: currVal}
	if list.size == 0 {
		list.head = node
		list.tail = node
	} else {
		if Compare(currVal, list.head.Val, CompareModeLess) || Compare(currVal, list.head.Val, CompareModeEqual) {
			node.next = list.head
			list.head.prev = node
			list.head = node
		} else {
			move := list.head
			for ; move != nil && Compare(currVal, move.Val, CompareModeGreater); move = move.next {
			}
			if move == nil {
				node.prev = list.tail
				list.tail.next = node
				list.tail = node
			} else {
				move.prev.next = node
				node.prev = move.prev
				node.next = move
				move.prev = node
			}
		}
	}
	list.size += 1
	list.Unlock()
}
