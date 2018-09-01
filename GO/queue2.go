// GO LINKED LIST QUEUE IMPLEMENTATION
package main
import (
	"fmt"
	"errors"
)

type Node struct {
	data int
	nextNode *Node
}

type QueueList struct {
	head *Node
	tail *Node
	length int
}

// 1. Method to add new nodes to the back of the list
func (q *QueueList) Enqueue(v int) QueueList {
	newNode := &Node{v, nil }
	fmt.Println(&newNode)
	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		old := q.tail
		old.nextNode = newNode
		q.tail = newNode
		fmt.Println("new tail: ", q.tail)

	}
	q.length++
	return *q
	//fmt.Println(q)
	//fmt.Println(q.head)
	//fmt.Println(q.tail)
	//fmt.Println("------item added-------")
}

// 2. Function to initialize new queue with content
func newQueue(items ...int) QueueList {
	q := QueueList{nil, nil, 0}
	for _, num := range items {
		q.Enqueue(num)
	}
	return q
}

// 3. Method that prints out queue contents FIFO
func (q *QueueList) Visit() {
	current := q.head
	for current != nil {
		fmt.Println(current.data)
		current = current.nextNode
	}
}

// 4. Method to check if queue is empty
func (q *QueueList) Empty() bool {
	if q.head == nil {
		return true
	}
	return false
}

// 5. Method to return size of the  queue
func (q *QueueList) Size() int {
	return q.length
}

// 6. Method to return the first node and remove it from list
func (q *QueueList) Dequeue() (int, error) {
	if q.head == nil {
		return -1, errors.New("no contents")
	} else {
		current := q.head
		q.head = current.nextNode
		q.length--
		fmt.Println(q)
		return current.data, nil
	}
}

// 7. Method to return smallest item in queue
func (q *QueueList) Min() (int, error) {
	if q.head == nil {
		return -1, errors.New("no contents")
	}
	min := q.head.data
	current := q.head
	for current != nil {
		if current.data < min {
			min = current.data
		}
		current = current.nextNode
	}
	return min, nil
}

// 8. Moethod to return largest item in queue
func (q *QueueList) Max() (int, error) {
	if q.head == nil {
		return -1, errors.New("no contents")
	}

	max := q.head.data
	current := q.head
	for current != nil {
		if current.data > max {
			max = current.data
		}
		current = current.nextNode
	}
	return max, nil
}

func main() {
	var myQueue = newQueue(1,2)

	fmt.Println("-------empty?-------")
	fmt.Println(myQueue.Empty())

	fmt.Println("----contents----")
	myQueue.Visit()

	fmt.Println("------enqueue 3------")
	myQueue.Enqueue(3)

	fmt.Println("------enqueue 4------")
	myQueue.Enqueue(4)

	fmt.Println("---min-max---")
	fmt.Println(myQueue.Min())
	fmt.Println(myQueue.Max())

	fmt.Println("------size------")
	fmt.Println(myQueue.Size())

	fmt.Println("----dequeue----")
	fmt.Println(myQueue.Dequeue())

	fmt.Println("------new-size------")
	fmt.Println(myQueue.Size())

	fmt.Println("----dequeue*3----")
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())

	fmt.Println("---newer-size---")
	fmt.Println(myQueue.Size())

	fmt.Println("---now-empty?-----")
	fmt.Println(myQueue.Empty())
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Min())
	fmt.Println(myQueue.Max())
}