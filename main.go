package main

import (
	"fmt"
)

func linearSearch(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, value int) int {

	var low int = 0
	var high int = len(arr) - 1

	for low <= high {
		var mid int = (high + low) / 2

		if value == arr[mid] {
			return mid
		} else if value > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func bubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}

	}
	return arr
}

type Node struct {
	next  *Node
    prev *Node
	value int
}

type LinkedList struct {
	head *Node
    length int
}

func (list *LinkedList) AddNode(value int) {
	newNode := &Node{nil, value}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) Traverse() {
	current := list.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (list *LinkedList) DeleteNode(value int) {
	if list.head == nil {
		return
	}

	if list.head.value == value {
		list.head = list.head.next
	}

	prev := list.head
	current := list.head.next

	for current != nil {
		if current.value == value {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}

}

type Queue struct {
	head *Node
	tail *Node
}

func (queue *Queue) Enqueue(val int) {

	newNode := &Node{nil, val}
	if queue.head == nil {
		queue.head = newNode
		queue.tail = newNode
		return
	}
	newNode.next = queue.head
	queue.head = newNode
	return

}

func (queue *Queue) Dequeue() int {

	if queue == nil {
		fmt.Println("Queue is empty")
		return -1
	}

	curr := queue.head
	for curr.next != queue.tail {
		curr = curr.next
	}
	tail := curr.next
	curr.next = nil

	return tail.value
}

func (queue *Queue) Peek() int {
	return queue.head.value
}

func (queue *Queue) Travesrse() {

	curr := queue.head

	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

type Stack struct {
	head *Node
}

func (stck *Stack) Push(val int) {

	newNode := Node{nil, val}

	if stck.head == nil {
		stck.head = &newNode
		return
	}

	prev := stck.head
	stck.head = &newNode
	stck.head.next = prev
	return
}

func (stck *Stack) Pop() int {

	val := stck.head
	stck.head = stck.head.next

	return val.value

}

func main() {
	que := Queue{
		head: nil,
		tail: nil,
	}

	que.Enqueue(2)
	que.Enqueue(3)
	que.Enqueue(5)

	que.Travesrse()
	fmt.Println("=================")
	que.Dequeue()
	que.Travesrse()
	fmt.Println("=================")

	stack := Stack{
		head: nil,
	}

	stack.Push(1)
	stack.Push(2)

	fmt.Println(stack.Pop())
    fmt.Println("Hello from Hussien")
	fmt.Println(stack.Pop())

}
