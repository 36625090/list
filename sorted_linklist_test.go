package list

import "testing"

func TestNewSortedLinkList(t *testing.T) {
	list := NewSortedLinkList[int]()
	list.Add(1)
	list.Add(3)
	list.Add(2)
	list.Add(8)
	list.Add(4)
	list.Add(5)
	list.Add(100)
	list.Add(17)
	list.Add(0)
	for node := list.tail; node != nil; node = node.prev {
		t.Log(node.Val)
	}

	println("======")
	for node := list.head; node != nil; node = node.next {
		t.Log(node.Val)
	}
}
