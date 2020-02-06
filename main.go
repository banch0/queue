package main

import (
	"fmt"
)

// Node ..
type Node struct {
	value    interface{}
	priority string
	previous *Node
	next     *Node
}

// Queue ...
type Queue struct {
	head       *Node
	tail       *Node
	nowInQueue *Node
	length     int
}

// CreateQueue constructor
func CreateQueue() *Queue {
	return &Queue{}
}

// showAllQueues show all
func (p *Queue) showAllQueues() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Queue is empty")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

// First ...
func (p *Queue) First() interface{} {
	p.nowInQueue = p.head
	return p.nowInQueue.value
}

// Last ...
func (p *Queue) Last() interface{} {
	p.nowInQueue = p.tail
	return p.nowInQueue.value
}

// Remove ...
func (p *Queue) Remove() interface{} {
	if p.length == 0 {
		return nil
	}

	p.length--

	if p.head.next == nil {
		return p.head
	}

	currentNode := p.head
	previous := p.head
	currentNode = currentNode.next
	previous.next = currentNode.next
	p.head = currentNode
	p.head.previous = nil
	return previous.value
}

// Insert ...
func (p *Queue) Insert(value interface{}, priority string) bool {
	newNode := &Node{
		value:    value,
		priority: priority,
	}
	p.length++

	if p.head == nil {
		p.head = newNode
	} else {
		current := p.tail
		current.next = newNode
		newNode.previous = p.tail
	}

	if newNode.priority == "vip" {
		newNode.next = p.head
		p.head.previous = newNode
		p.head = newNode
		return true
	}

	p.tail = newNode
	return true
}

func main() {}
