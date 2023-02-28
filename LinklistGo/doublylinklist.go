package main

import "fmt"

// Define the Node struct
type Node struct {
	data int
	next *Node
	prev *Node
}

// Define the LinkedList struct
type LinkedList struct {
	head *Node
	tail *Node
}

// Define a method to create a new node
func NewNode(data int) *Node {
	return &Node{data: data, next: nil, prev: nil}
}

// Define a method to create a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}

// Define a method to add a node to the front of the linked list
func (ll *LinkedList) AddFront(data int) {
	newNode := NewNode(data)

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	ll.head.prev = newNode
	newNode.next = ll.head
	ll.head = newNode
}

// Define a method to remove the front node from the linked list
func (ll *LinkedList) RemoveFront() {
	if ll.head == nil {
		return
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	ll.head = ll.head.next
	ll.head.prev = nil
}

// Define a method to add a node to the end of the linked list
func (ll *LinkedList) AddEnd(data int) {
	newNode := NewNode(data)

	if ll.tail == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	ll.tail.next = newNode
	newNode.prev = ll.tail
	ll.tail = newNode
}

// Define a method to remove the last node from the linked list
func (ll *LinkedList) RemoveEnd() {
	if ll.tail == nil {
		return
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	ll.tail = ll.tail.prev
	ll.tail.next = nil
}

// Define a method to insert a node at a specific index
func (ll *LinkedList) InsertAt(data int, index int) {
	newNode := NewNode(data)

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	if index == 0 {
		ll.AddFront(data)
		return
	}

	current := ll.head
	for i := 1; i < index && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		ll.AddEnd(data)
		return
	}

	newNode.next = current.next
	newNode.prev = current
	current.next = newNode
	newNode.next.prev = newNode
}

// Define a method to remove a node at a specific index
func (ll *LinkedList) RemoveAt(index int) {
	if ll.head == nil {
		return
	}

	if index == 0 {
		ll.RemoveFront()
		return
	}

	current := ll.head
	for i := 1; i < index && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		return
	}

	if current == ll.tail {
		ll.RemoveEnd()
		return
	}

	current.prev.next = current.next
	current.next.prev = current.prev
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

	ll.AddEnd(10)
	ll.AddEnd(20)
	ll.AddEnd(30)

	ll.PrintList() // Output: 10 20 30

	ll.AddFront(5)

	ll.PrintList() // Output: 5 10 20 30

	ll.RemoveAt(2)

	ll.PrintList() // Output: 5 10 30

	ll.InsertAt(15, 1)

	ll.PrintList() // Output: 5 15 10 30

	ll.RemoveFront()

	ll.PrintList() // Output: 15 10 30

	ll.RemoveEnd()

	ll.PrintList() // Output: 15 10
}
