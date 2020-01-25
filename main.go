package main

import (
	"fmt"
)

// Вспомните, как работает стратегия ~разделяй и властвуй:
// 1. Определите простейший случай как базовый.
// 2. Придумайте, как свести задачу к базовому случаю.

// Node ..
type Node struct {
	name     string
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

// AddQueue ...
func (p *Queue) AddQueue(name, priority string) {
	s := &Node{
		name:     name,
		priority: priority,
	}

	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.tail
		currentNode.next = s
		s.previous = p.tail
	}
	p.tail = s
	p.length++

}

func (p *Queue) showAllQueues() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Queue is empty")
		return nil
	}
	fmt.Printf("Z%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("Q%+v\n", *currentNode)
	}
	return nil
}

func (p *Queue) startQueue() *Node {
	p.nowInQueue = p.head
	return p.nowInQueue
}

func (p *Queue) nextQueue() *Node {
	p.nowInQueue = p.nowInQueue.next
	return p.nowInQueue
}

func (p *Queue) previousQueue() *Node {
	p.nowInQueue = p.nowInQueue.previous
	return p.nowInQueue
}

// First ...
func (p *Queue) First() {
	p.nowInQueue = p.head
	fmt.Println("First element: ", p.nowInQueue.name)
}

// Last ...
func (p *Queue) Last() {
	p.nowInQueue = p.tail
	fmt.Println("Last element: ", p.nowInQueue.name)
}

func main() {
	myQueue := CreateQueue()
	// fmt.Println("Created queue")
	// fmt.Println()
	myQueue.Insert()
	// fmt.Print("Adding clients to queue...\n\n")
	// myQueue.showAllQueues()
	// myQueue.insertVip("Tyler", "vip")
	myQueue.AddQueue("Marla", "simple")
	myQueue.AddQueue("Bob", "simple")
	myQueue.AddQueue("Cornelius", "simple")
	myQueue.AddQueue("Dod", "simple")
	// myQueue.insertVip("Dan", "vip")

	// need show positon in queue

	// fmt.Println("Show len of queue: ", myQueue.length)
	// fmt.Println()

	// fmt.Println("Showing all clients in queue...")
	// myQueue.showAllQueues()
	// fmt.Println()

	// myQueue.startQueue()
	// fmt.Printf("Now in queue: %s priority %s\n", myQueue.nowInQueue.name, myQueue.nowInQueue.priority)
	// fmt.Println()

	// myQueue.nextQueue()
	// fmt.Println("Call next client...")
	// fmt.Printf("Now in queue: %s priority %s\n", myQueue.nowInQueue.name, myQueue.nowInQueue.priority)
	// fmt.Println()

	// myQueue.previousQueue()
	// fmt.Println("Back to previous client...")
	// fmt.Printf("Now in queue: %s priority %s\n", myQueue.nowInQueue.name, myQueue.nowInQueue.priority)
	// fmt.Println()

	// fmt.Println()
	// myQueue.First()
	// myQueue.Last()

	// fmt.Println()
	// myQueue.Remove()
	// fmt.Println()
	// myQueue.First()
	// myQueue.Last()
	// myQueue.showAllQueues()
}

// Remove ...
func (p *Queue) Remove() {
	if p.length == 0 {
		fmt.Printf("Queue is empty")
	}

	var previousClient *Node
	currentNode := p.head

	previousClient = currentNode
	currentNode = currentNode.next
	previousClient.next = currentNode.next
	p.length--

	p.head = currentNode
	p.head.previous = nil
}

// Insert ...
func (p *Queue) Insert(d ...Node) bool {
	newNode := &d[0]
	if newNode.priority == "vip" {
		if p.head.priority == "vip" {
			newNode.next = p.head.next
			p.head.next = newNode
			newNode.previous = p.head
			p.length++
			return true
		}
		newNode.next = p.head
		p.head.previous = newNode
		p.head = newNode
		p.length++
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
	p.length++
	return true
}
