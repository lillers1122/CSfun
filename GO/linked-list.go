// GO LINKED LIST IMPLEMENTATION
package main
import ("fmt"
"errors")

type Node struct {
	data int
	nextNode *Node
}

type LinkedList struct {
	head *Node
}

// 1. Add a Node
func (l *LinkedList) add(inputData int) {
	newNode := Node{data: inputData, nextNode: l.head}
	l.head = &newNode
}

// 2. Search
func (l *LinkedList) search(input int) bool {
	current := l.head
	for current != nil {
		if current.data == input {
			return true
		}
		current = current.nextNode
	}
	return false
}

// 3. Find Max - USES ERROR FOR NIL
func (l *LinkedList) findMax() (int, error) {
	if l.head == nil {
		return -1, errors.New("empty list")
	}

	max := l.head.data
	current := l.head.nextNode

	for current != nil {
		if current.data > max {
			max = current.data
		}
		current = current.nextNode
	}
	return max, nil
}

// 4. Find Min
func (l *LinkedList) findMin() int {
	min := l.head.data
	current := l.head.nextNode
	for current != nil {
		if current.data < min {
			min = current.data
		}
		current = current.nextNode
	}
	return min
}

// 5. Length
func (l *LinkedList) length() int {
	count := 0
	current := l.head

	for current != nil {
		count++
		current = current.nextNode
	}
	return count
}

// 6. Find Nth from Beginning
func (l *LinkedList) findN(value int) int {
	count := 0
	current := l.head
	for current != nil {
		if count == (value - 1) {
			return current.data
		}
		count++
		current = current.nextNode
	}
	return current.data
}

// 7. Insert Ascending, assumes LL is sorted - Thank you Karinna!
func (l *LinkedList) ascInsert(inputData int) {
	newNode := Node{data: inputData}

	// if list is empty, just insert node as head
	if l.head == nil {
		l.head = &newNode
	} else if l.head.data > inputData {
		// if the head is bigger than the input data, insert node as head's nextNode
		newNode.nextNode = l.head
		l.head = &newNode
	} else {
		current := l.head

		//if list has one, smaller node: point to new node last
		if current.nextNode == nil {
			current.nextNode = &newNode
		} else {
			// otherwise, iterate through list: stop when current's next node is bigger than the input data, or it's the end of the list.
			for current.nextNode != nil && current.nextNode.data < inputData {
				current = current.nextNode
			}
			// assign newNode's nextNode if current.next's data > inputData or nil as the; reassign current's next node as newNode
			newNode.nextNode = current.nextNode
			current.nextNode = &newNode

		}
	}
}

// 8. Visit, print each node
func (l *LinkedList) visit() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.nextNode
	}
}

// 9. Delete first found of specified value
func (l *LinkedList) delete(inputData int) {
	current := l.head
	if current.data == inputData { // head is value
		l.head = current.nextNode
		return
	}
	for current.nextNode != nil && current.nextNode.data != inputData {
		current = current.nextNode
	}
	if current.nextNode != nil {
		current.nextNode = current.nextNode.nextNode
	}
}

// 10. Reverse list
func (l *LinkedList) reverseList() {
	//declare variables
	var previous *Node = nil
	current := l.head
	var next *Node = nil

	//reverse list
	for current != nil {
		next = current.nextNode
		current.nextNode = previous
		previous = current
		current = next
	}

	//reset head
	l.head = previous
}

func main() {
	numList := &LinkedList{nil }

	//Add more nodes
	numList.add(1)
	numList.add(2)
	numList.add(3)
	numList.add(4)
	numList.visit()
	fmt.Println("-------------")

	//reverse and print list
	//numList.reverseList()
	//numList.visit()
	//fmt.Println("-------------")

	//insert ascending
	numList.reverseList()
	numList.visit()
	fmt.Println("-------------")
	numList.ascInsert(7)
	numList.ascInsert(5)
	numList.visit()
	fmt.Println("-------------")

	//delete and print list
	//numList.delete(3)
	//numList.delete(4)
	//numList.delete(1)
	//numList.visit()
	//numList.delete(2)
}
