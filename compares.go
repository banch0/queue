package main

import (
	"container/list"
	"fmt"
	"strconv"
	"time"
)

// ContainerList ...
func ContainerList() {
	// New list.
	values := list.New()
	// Add 3 elements to the list.
	values.PushBack("bird")
	values.PushBack("cat")
	values.PushFront("snake")
	// Add 100 elements at the front.
	for i := 0; i < 5; i++ {
		// Convert ints to strings.
		values.PushFront(strconv.Itoa(i))
	}
	values.PushBack("dog")
	// Loop over container list.
	for temp := values.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
}

// CompareListAndSlice ...
func CompareListAndSlice() {
	t0 := time.Now()

	// Version 1: use container list.
	for i := 0; i < 10000; i++ {
		// New list.
		values := list.New()
		// Add 2 elements to the list.
		values.PushBack("bird")
		values.PushBack("cat")
		// Add 20 elements at the front.
		for i := 0; i < 20; i++ {
			// Convert ints to strings.
			values.PushFront(strconv.Itoa(i))
		}
	}

	t1 := time.Now()

	// Version 2: use slice.
	for i := 0; i < 10000; i++ {
		// New empty slice.
		values := []string{}
		// Add 2 elements to the slice.
		values = append(values, "bird")
		values = append(values, "cat")
		// Add 20 elements at the front.
		for i := 0; i < 20; i++ {
			// Create a new slice and put the string at its start.
			// ... This inserts as the front.
			tempSlice := []string{}
			tempSlice = append(tempSlice, strconv.Itoa(i))
			// Now append all previous strings after the first one.
			for x := range values {
				tempSlice = append(tempSlice, values[x])
			}
			// Use the new slice.
			values = tempSlice
		}
	}

	t2 := time.Now()
	// Results.
	fmt.Println(t1.Sub(t0))
	fmt.Println(t2.Sub(t1))
}

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []int) []int {
	element := queue[0]
	fmt.Println("Dequeued:", element)
	return queue[1:]
}

func sliceQueue() {
	var queue []int // Make a queue of ints.

	queue = enqueue(queue, 10)
	queue = enqueue(queue, 20)
	queue = enqueue(queue, 30)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 40)
	fmt.Println("Queue:", queue)
}
