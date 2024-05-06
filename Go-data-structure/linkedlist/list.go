package linkedlist

import "fmt"

/// Linked List in Golang

// Linked List Data Structure

type node struct {
	data interface{}
	next *node
}

func newLinkNode(data interface{}) *node {
	return &node{
		data: data,
		next: nil,
	}
}

type LinkedList struct {
	head *node
	size uint
}

// NewLinkedList create a linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

/// Public function of linked list

// Len get the linked list length
func (list *LinkedList) Len() uint {
	return list.size
}

// Insert insert on the head of the linked list
func (list *LinkedList) Insert(data int) *LinkedList {
	newNode := newLinkNode(data)
	newNode.next = list.head
	list.head = newNode
	list.size += 1
	return list
}

// RemoveHead delete head node at the linked list
func (list *LinkedList) RemoveHead() {
	if list.head == nil {
		return
	}
	list.head = list.head.next
	list.size -= 1
}

// RemoveItem delete the n node
func (list *LinkedList) RemoveItem(n *node) {
	if n == nil {
		return
	}
	if n == list.head {
		list.RemoveHead()
		return
	}
	cur := list.head
	for cur.next != nil {
		if cur.next == n {
			cur.next = cur.next.next
			list.size -= 1
			return
		}
		cur = cur.next
	}
}

// Find find the node
func (list *LinkedList) Find(data int) *node {
	if list.head == nil {
		return nil
	}
	cur := list.head
	for cur != nil {
		if cur.data == data {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// Head get the head node
func (list *LinkedList) Head() *node {
	return list.head
}

// Print print the linked list
func (list *LinkedList) Print() string {
	var format string
	cur := list.head
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.data)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	// fmt.Println(format)
	return format
}


/// Alogrithm of linked list

// Reverse reverse the linked list
func (list *LinkedList) Reverse() {
	if list.head == nil || list.head.next == nil { // p || p->next is nil
		return
	}
	var pre *node = nil
	cur := list.head
	for cur != nil {
		next := cur.next
		cur.next = pre  // 指向前一个节点
		pre = cur		// 前一个节点后移
		cur = next
	}
	list.head = pre
}
