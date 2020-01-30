package main

import (
	"fmt"
)

// Node ..
type Node struct {
	name     interface{}
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
	return p.nowInQueue.name
}

// Last ...
func (p *Queue) Last() interface{} {
	p.nowInQueue = p.tail
	return p.nowInQueue.name
}

// Remove ...
func (p *Queue) Remove() error {
	if p.length == 0 {
		return nil
	}

	p.length--

	if p.head.next == nil {
		p.head = nil
		return nil
	}

	currentNode := p.head
	previous := p.head
	currentNode = currentNode.next
	previous.next = currentNode.next
	p.head = currentNode
	p.head.previous = nil
	return nil
}

// Insert ...
func (p *Queue) Insert(name, priority string) bool {
	newNode := &Node{name: name, priority: priority}
	p.length++
	if newNode.priority == "vip" {
		if p.head.priority == "vip" {
			newNode.next = p.head.next
			p.head.next = newNode
			newNode.previous = p.head
			return true
		}
		newNode.next = p.head
		p.head.previous = newNode
		p.head = newNode
		return true
	}

	if p.head != nil {
		current := p.tail
		current.next = newNode
		newNode.previous = p.tail
	} else {
		p.head = newNode
	}

	p.tail = newNode
	return true
}

// Main ...
func main() {}
