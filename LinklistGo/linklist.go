package main

import (
	"fmt"
)

// Node struct represents a node in the linked list
type Node struct {
	data interface{}
	next *Node
}

// LinkedList struct represents a linked list
type LinkedList struct {
	head *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

// AddFront adds a node to the front of the linked list
func (ll *LinkedList) AddFront(data interface{}) {
	newNode := &Node{data, nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}

// RemoveFront removes the front node from the linked list
func (ll *LinkedList) RemoveFront() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.next
}

// AddEnd adds a node to the end of the linked list
func (ll *LinkedList) AddEnd(data interface{}) {
	newNode := &Node{data, nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// RemoveEnd removes the last node from the linked list
func (ll *LinkedList) RemoveEnd() {
	if ll.head == nil {
		return
	}
	if ll.head.next == nil {
		ll.head = nil
		return
	}
	current := ll.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

// InsertAt inserts a node at a specific index
func (ll *LinkedList) InsertAt(data interface{}, index int) {
	newNode := &Node{data, nil}
	if index == 0 {
		newNode.next = ll.head
		ll.head = newNode
	} else {
		current := ll.head
		for i := 0; i < index-1 && current != nil; i++ {
			current = current.next
		}
		if current == nil {
			return
		}
		newNode.next = current.next
		current.next = newNode
	}
}

// RemoveAt removes a node at a specific index
func (ll *LinkedList) RemoveAt(index int) {
	if ll.head == nil {
		return
	}
	if index == 0 {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}
	if current == nil || current.next == nil {
		return
	}
	current.next = current.next.next
}

// PrintList prints all the nodes in the linked list
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	ll := NewLinkedList()

	// Add nodes to the front
	ll.AddFront(3)
	ll.AddFront(2)
	ll.AddFront(1)
	ll.PrintList() // Should print "1 2 3"

	// Remove the front node
	ll.RemoveFront()
	ll.PrintList() // Should print "2 3"

	// Add nodes to the end
	ll.AddEnd(4)
	ll.AddEnd(5)
	ll.PrintList() // Should print "2 3 4 5"

	// Remove the last node
	ll.RemoveEnd()
	ll.PrintList() // Should print "2 3 4"

	// Insert a node at a specific index
	ll.InsertAt(7, 1)
	ll.PrintList() // Should print "2 7 3 4"

	// Remove a node at a specific index
	ll.RemoveAt(2)
	ll.PrintList() // Should print "2 7 4"
}
