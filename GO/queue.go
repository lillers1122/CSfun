// GO SLICE QUEUE IMPLEMENTATION
package main
import (
	"fmt"
	"errors"
	)

type queue []int // slice

// 1. Method to enqueue into queue
func (q *queue) Enqueue(v int) *queue {
	*q = append(*q, v)
	return q
}

// 2. Function to initialize new queue with content
func newQueue(items ...int) queue {
	q := queue{}
	for _, num := range items {
		q.Enqueue(num)
	}
	return q
}

// 3. Method to check if queue is empty
func (q *queue) Empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

// 4. Method to print size of queue
func (q *queue) Size() int {
	size := len(*q)
	fmt.Println(size)
	return size
}

// Extra method added to see slice capacity
func (q *queue) Cap() int {
	capacity := cap(*q)
	fmt.Println(capacity)
	return capacity
}

// 5. Method to dequeue from queue
func (q *queue) Dequeue() (int, error) {
	if len(*q) < 1 {
		return -1, errors.New("no content")
	}
	num := (*q)[0]
	*q = (*q)[1:]
	fmt.Println(num)
	return num, nil
}

// 6. Method to return smallest item in queue
func (q *queue) Min() (int, error) {
	if len(*q) < 1 {
		return -1, errors.New("no content")
	}
	min := (*q)[0]

	for _, v := range *q {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// 7. Method to return largest item in queue
func (q *queue) Max() (int, error) {
	if len(*q) < 1 {
		return -1, errors.New("no content")
	}
	max := (*q)[0]

	for _, v := range *q {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func main() {
	var myQueue = newQueue(6)
	fmt.Println("---empty?---")
	fmt.Println(myQueue.Empty())
	fmt.Println("---contents---")
	fmt.Println(myQueue)

	//Check Methods
	fmt.Println("---enqueue 7---")
	myQueue.Enqueue(7)
	fmt.Println("---enqueue 8---")
	myQueue.Enqueue(8)
	fmt.Println("---min-max---")
	fmt.Println(myQueue.Min())
	fmt.Println(myQueue.Max())
	fmt.Println("---size---")
	myQueue.Size()
	myQueue.Cap()
	fmt.Println("---dequeue---")
	myQueue.Dequeue()
	fmt.Println("---new-size---")
	myQueue.Size()
	myQueue.Cap()
	fmt.Println("---dequeue*2---")
	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println("---newer-size---")
	myQueue.Size()
	myQueue.Cap()
	fmt.Println("---now-empty?-----")
	fmt.Println(myQueue.Empty())
	fmt.Println(myQueue)
	myQueue.Size()
	myQueue.Cap()
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Min())
	fmt.Println(myQueue.Max())
}