package linkedlist_test

import (
	"GoAlgo/linkedlist"
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	list := linkedlist.NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	if list.Len() != 5 {
		t.Errorf("list length should be 5, got %d", list.Len())
	}
}

func TestLinkedList_RemoveHead(t *testing.T) {
	list := linkedlist.NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.RemoveHead()
	list.RemoveHead()
	if list.Len() != 3 {
		t.Errorf("list length should be 3, got %d", list.Len())
	}
}

func TestLinkedList_RemoveItem(t *testing.T) {
	list := linkedlist.NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.RemoveItem(list.Head())
	list.RemoveItem(list.Head())
	if list.Len() != 3 {
		t.Errorf("list length should be 3, got %d", list.Len())
	}
}

func TestPrint(t *testing.T) {
	list := linkedlist.NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	if list.Print() != "5->4->3->2->1" {
		t.Errorf("list print should be 5->4->3->2->1, got %s", list.Print())
	}
}

func TestReverse(t *testing.T) {
	list := linkedlist.NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.Reverse()
	if list.Print() != "1->2->3->4->5" {
		t.Errorf("list print should be 1->2->3->4->5, got %s", list.Print())
	}
}
