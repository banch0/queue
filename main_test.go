package main

import "testing"

// func Test_empty_list(t *testing.T) {
// 	q := CreateQueue()
// 	if q.length != 0 {
// 		t.Error("new list length must be zero, got: ", q.length)
// 	}
// }

// func Test_list_with_one_item(t *testing.T) {
// 	q := CreateQueue()
// 	q.AddQueue("Bob", "simple")
// 	if q.length != 1 {
// 		t.Error("after adding one element size must be 1, got: ", q.length)
// 	}
// }

// func Test_list_with_multiple_items(t *testing.T) {
// 	q := CreateQueue()
// 	q.AddQueue("Tyler", "smple")
// 	q.AddQueue("Marla", "smple")
// 	q.AddQueue("Cornelius", "smple")
// 	if q.length != 3 {
// 		t.Error("after adding 3 elements size must be 3, got: ", q.length)
// 	}
// }

func TestVipClient(t *testing.T) {
	q := CreateQueue()
	a := &Node{
		name:     "Marla",
		priority: "simple",
	}
	q.Insert(*a)
	q.AddQueue("Roger", "simple")
	q.AddQueue("Roger2", "simple")
	ab := &Node{
		name:     "Bob",
		priority: "simple",
	}
	q.Insert(*ab)
	ac := &Node{
		name:     "Tyler",
		priority: "vip",
	}
	q.Insert(*ac)
	ad := &Node{
		name:     "Cornelius",
		priority: "vip",
	}
	q.Insert(*ad)
	q.showAllQueues()
	if q.head.priority != "vip" {
		t.Error("vip client not inserted to start of queue")
	}
	// q.First() // show first element
	// log.Println("length of queue: ", q.length)
}
