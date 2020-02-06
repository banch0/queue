package main

import (
	"testing"
)

func TestEmptyList(t *testing.T) {
	q := CreateQueue()
	if q.length != 0 {
		t.Error("New queue length must be zero, but got: ", q.length)
	}

	first := q.First()
	if q.First() != first {
		t.Error("First element of empty queue must be nil, got: ", first)
	}

	last := q.Last()
	if q.Last() != last {
		t.Error("Last element of empty queue must be nil, got: ", last)
	}

	remove := q.Remove()
	if q.Remove() != remove {
		t.Error("Remove from empty list must be nil, got: ", remove)
	}

	queueNumber := 0001
	q.Insert(queueNumber, "simple")

	if q.length != 1 {
		t.Error("queue length mus be 1 but, got: ", q.length)
	}
}

func TestRemoveClient(t *testing.T) {
	q := CreateQueue()

	q.Remove()
	if q.Remove() != nil {
		t.Error("Error: expected nil")
	}

	client := struct {
		name    string
		account int
	}{
		name:    "Bob",
		account: 12345678,
	}

	q.Insert(client, "vip")

	remove := q.Remove()

	if remove != client {
		t.Error("Return removing element error")
	}

	if q.length != 0 {
		t.Error("remove function don't work")
	}
}

func TestWithOneItem(t *testing.T) {
	q := CreateQueue()
	c := struct {
		name    string
		account int64
	}{
		name:    "Bob",
		account: 23454543,
	}

	q.Insert(c, "simple")

	if q.length != 1 {
		t.Error("after adding one element size must be 1, got: ", q.length)
	}

	first := q.First()
	last := q.Last()

	if first != last {
		t.Error("first != last")
	}

	q.Remove()

	if q.length != 0 {
		t.Error("After removing elements, size must be 0, but got: ", q.length)
	}
}

func TestMultipleItems(t *testing.T) {
	q := CreateQueue()
	q.Insert(1, "simple")
	q.Insert(2, "vip")
	q.Insert(3, "simple")
	if q.length != 3 {
		t.Error("after adding 3 elements size must be 3, got: ", q.length)
	}
}

func TestVipClient(t *testing.T) {
	q := CreateQueue()
	Marla := struct {
		name    string
		account int64
	}{
		name:    "Marla",
		account: 23453,
	}
	q.Insert(Marla, "simple")
	Bob := struct {
		name    string
		account int64
	}{
		name:    "Bob",
		account: 23454,
	}
	q.Insert(Bob, "simple")
	Tyler := struct {
		name    string
		account int64
	}{
		name:    "Tyler",
		account: 23455,
	}
	q.Insert(Tyler, "vip")

	if q.head.priority != "vip" {
		t.Error("vip client not inserted to start of queue")
	}
}

func TestFirst(t *testing.T) {
	q := CreateQueue()
	value := struct{ name string }{name: "Bob"}

	q.Insert(value, "vip")

	if q.First() != value {
		t.Errorf("First element must be equal to %v, but got %v: ", value, q.First()) //q.nowInQueue.name

	}
}

func TestLast(t *testing.T) {
	q := CreateQueue()
	value := struct{ name string }{name: "Bob"}

	q.Insert(value, "simple")

	if q.Last() != value {
		t.Errorf("Last element must be equal to %v, but got: %v ", value, q.Last())

	}
}
