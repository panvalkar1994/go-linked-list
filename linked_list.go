package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	len  int
}

// Insert data in linked list at head
func (l *List) Insert(d int) {
	node := &Node{
		data: d,
		next: nil,
	}
	node.next = l.head
	l.head = node
	l.len += 1
}

// Display Linked List data
func (l *List) Display() {
	node := l.head
	for node != nil {
		fmt.Printf("%d->", node.data)
		node = node.next
	}

	if l.head == nil {
		fmt.Print(".\n")
	}
}

// Append data at end of linked list
func (l *List) Append(d int) {
	if l.len == 0 {
		l.Insert(d)
		return
	}

	cur_node := l.head
	for cur_node.next != nil {
		cur_node = cur_node.next
	}

	node := &Node{
		data: d,
		next: nil,
	}
	cur_node.next = node
	l.len += 1
}

// Get Last Element
func (l *List) Get_Tail() (int, error) {
	if l.head == nil {
		return 0, errors.New("Empty List")
	}
	cur_node := l.head
	for cur_node.next != nil {
		cur_node = cur_node.next
	}

	return cur_node.data, nil

}
