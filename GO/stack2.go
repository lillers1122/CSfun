// GO LINKED LIST STACK IMPLEMENTATION
package main
import ( "fmt"
"errors")

type Node struct {
	data int
	nextNode *Node
}

type StackList struct {
	head *Node
}

// 1. Function to initialize new stack with content
func newStack(items ...int) StackList {
	n := StackList{nil}
	for _, num := range items {
		n.Push(num)
	}
	return n
}

// 2. Method to push int to stack (to the head of the list)
func (n *StackList) Push(inputData int) {
	newNode := Node{inputData, n.head}
	n.head = &newNode
}

// 3. Method to print contents of stack (head forward)
func (l *StackList) Visit() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.nextNode
	}
}

// 4. Method to check if stack is empty
func (n *StackList) Empty() bool {
	if n.head == nil {
		return true
	}
	return false
}

// 5. Method to return size of the stack
func (n *StackList) Size() int {
	count := 0
	current := n.head
	for current != nil {
		count += 1
		current = current.nextNode
	}
	return count
}

// 6. Method to remove top item from the stack
func (n *StackList) Pop() (int, error) {
	if n.head == nil {
		return -1, errors.New("no contents")
	}
	current := n.head
	n.head = current.nextNode
	return current.data, nil
}

// 7. Method to return smallest item in stack
func (n *StackList) Min() (int, error) {
	if n.head == nil {
		return -1, errors.New("no contents")
	}
	current := n.head
	min := current.data
	for current != nil {
		if current.data < min {
			min = current.data
		}
		current = current.nextNode
	}
	return min, nil
}

// 8. Method to return largest item in stack
func (n *StackList) Max() (int, error) {
	if n.head == nil {
		return -1, errors.New("no contents")
	}
	current := n.head
	max := current.data
	for current != nil {
		if current.data > max {
			max = current.data
		}
		current = current.nextNode
	}
	return max, nil
}

func main() {
	//Initialize stack with items
	var numStack = newStack(4,3)
	fmt.Println("---empty?---")
	fmt.Println(numStack.Empty())
	fmt.Println("---contents---")
	numStack.Visit()

	//Check Methods
	fmt.Println("---Push 2---")
	numStack.Push(2)
	numStack.Visit()
	fmt.Println("---min-max---")
	fmt.Println(numStack.Min())
	fmt.Println(numStack.Max())
	fmt.Println("---size---")
	fmt.Println(numStack.Size())
	fmt.Println("---pop---")
	fmt.Println(numStack.Pop())
	fmt.Println("---new-size---")
	fmt.Println(numStack.Size())
	fmt.Println("---pop*2---")
	fmt.Println(numStack.Pop())
	fmt.Println(numStack.Pop())
	fmt.Println("---newer-size---")
	fmt.Println(numStack.Size())
	fmt.Println("---now-empty?---")
	fmt.Println(numStack.Empty())
	fmt.Println(numStack.Pop())
	fmt.Println(numStack.Size())
	fmt.Println(numStack.Max())
	fmt.Println(numStack.Min())
}
