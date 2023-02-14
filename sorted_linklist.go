package list

import (
	"github.com/36625090/lambda/compare"
)

type SortedLinkList[T any] struct {
	LinkList[T]
}

func NewSortedLinkList[T any]() *SortedLinkList[T] {
	return &SortedLinkList[T]{}
}

func (list *SortedLinkList[T]) Add(val T) {
	list.Lock()
	node := &Node[T]{Val: val}
	if list.size == 0 {
		list.head = node
	} else {
		curr := list.head
		if compare.Compare(val, curr.Val, compare.CmpModeLess) {
			node.next = curr
			list.head = node
		} else {
			for curr.next != nil && compare.Compare(val, curr.next.Val, compare.CmpModeGreater) {
				curr = curr.next
			}
			node.next = curr.next
			curr.next = node
		}
	}
	list.size += 1
	list.Unlock()
}
