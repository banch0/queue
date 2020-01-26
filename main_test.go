package main

import (
	"testing"
)

func Test_empty_list(t *testing.T) {
	q := CreateQueue()
	if q.length != 0 {
		t.Error("new list length must be zero, got: ", q.length)
	}
}

func Test_remove_client(t *testing.T) {
	q := CreateQueue()

	q.Insert(Node{
		name:     "Bob",
		priority: "simple",
	})
	q.Remove()
	if q.length != 0 {
		t.Error("remove function don't work")
	}
}

func Test_list_with_one_item(t *testing.T) {
	q := CreateQueue()
	sec := &Node{
		name:     "Bob",
		priority: "simple",
	}
	q.Insert(*sec)
	if q.length != 1 {
		t.Error("after adding one element size must be 1, got: ", q.length)
	}
}

func Test_list_with_multiple_items(t *testing.T) {
	q := CreateQueue()
	fir := &Node{
		name:     "Marla",
		priority: "simple",
	}
	q.Insert(*fir)
	sec := &Node{
		name:     "Bob",
		priority: "simple",
	}
	q.Insert(*sec)
	thi := &Node{
		name:     "Tyler",
		priority: "vip",
	}
	q.Insert(*thi)
	if q.length != 3 {
		t.Error("after adding 3 elements size must be 3, got: ", q.length)
	}
}

func TestVipClient(t *testing.T) {
	q := CreateQueue()
	Marla := &Node{
		name:     "Marla",
		priority: "simple",
	}
	q.Insert(*Marla)
	Bob := &Node{
		name:     "Bob",
		priority: "simple",
	}
	q.Insert(*Bob)
	Tyler := &Node{
		name:     "Tyler",
		priority: "vip",
	}
	q.Insert(*Tyler)
	Tyler2 := &Node{
		name:     "Tyler2",
		priority: "vip",
	}
	q.Insert(*Tyler2)
	Tyler3 := &Node{
		name:     "Tyler3",
		priority: "vip",
	}
	q.Insert(*Tyler3)
	// q.showAllQueues()
	if q.head.priority != "vip" {
		t.Error("vip client not inserted to start of queue")
	}
}

func TestFirst(t *testing.T) {
	q := CreateQueue()
	q.Insert(Node{name: "Bob", priority: "simple"})
	if q.First() != "Bob" {
		t.Error("First element name must be Bob, got: ", q.nowInQueue.name)
	}
}

func TestLast(t *testing.T) {
	q := CreateQueue()
	q.Insert(Node{name: "Bob", priority: "simple"})
	if q.Last() != "Bob" {
		t.Error("Last element name must be Bob, got: ", q.nowInQueue.name)
	}
}
