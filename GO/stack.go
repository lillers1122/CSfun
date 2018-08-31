// GO SLICE STACK IMPLEMENTATION
package main
import (
	"fmt"
	"errors")

type stack []int // slice

// 1. Function to initialize new stack with content
func newStack(items ...int) stack {
	s := stack{}
	for _, num := range items {
		s.Push(num)
	}
	return s
}

// 2. Method to push int into stack
func (s *stack) Push(v int) *stack {
	*s = append(*s, v)
	return s
}

// 3. Method to check if stack is empty
func (s *stack) Empty() bool {
	if len(*s) < 1 {
		return true
	}
	return false
}

// 4. Method to print size of slice
func (s *stack) Size()  {
	fmt.Println(len(*s))
}

// 5. Method to remove top item from stack
func (s *stack) Pop() (int, error) {
	if len(*s) == 0 {
		return -1, errors.New("no content!")
	}
	var x = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x, nil
}

// 6. Method to return smallest item in stack
func (s *stack) Min() (int, error) {
	if len(*s) == 0 {
		return -1, errors.New("no content!")
	}
	var min = (*s)[0]

	for _, v := range *s {
		if min > v {
			min = v
		}
	}
	return min, nil
}

// 7. Method to return largest item in stack
func (s *stack) Max() (int, error) {
	if len(*s) == 0 {
		return -1, errors.New("no content!")
	}

	var max = (*s)[0]
	for _, v := range *s {
		if max < v {
			max = v
		}
	}
	return max, nil
}

func main() {
	var myStack = newStack(1,2,3,4)
	fmt.Println("---empty?---")
	fmt.Println(myStack.Empty())
	fmt.Println("---contents---")
	fmt.Println(myStack)

	//Check Methods
	fmt.Println("---Push 5---")
	myStack.Push(5)
	fmt.Println(myStack)
	fmt.Println("---min-max---")
	fmt.Println(myStack.Min())
	fmt.Println(myStack.Max())
	fmt.Println("---size---")
	myStack.Size()
	fmt.Println("---pop---")
	fmt.Println(myStack.Pop())
	fmt.Println("---new-size---")
	myStack.Size()
	fmt.Println("---pop*4---")
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println("---newer-size---")
	myStack.Size()
	fmt.Println("---now-empty?-----")
	fmt.Println(myStack.Empty())
	fmt.Println(myStack)
	myStack.Size()
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Min())
	fmt.Println(myStack.Max())
}