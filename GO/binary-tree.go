// GO BINARY SEARCH TREE
package main

import "fmt"

type Node struct {
	data int
	lChild *Node
	rChild *Node
}

type BinaryTree struct {
	root *Node
}

// Initialize Tree with content
func newBinaryTree(items ...int) BinaryTree {
	n := BinaryTree{nil }
	for _, num := range items {
		n.add(num)
	}
	return n
}

// 1. Search - return bool
func (n *BinaryTree) search(v int) bool {
	current := n.root
	for current != nil {
		//root is int
		if v == current.data {
			return true
		} else if v < current.data {
			current = current.lChild
		} else {
			current = current.rChild
		}
	}
	return false
}

// 2. Insert given int value
func (n *BinaryTree) add(inputData int) {
	newNode := Node{inputData, nil, nil }
	//empty tree
	if n.root == nil {
		n.root = &newNode
	} else {
		//not empty
		insertHelper(n.root, &newNode)
	}
}

func insertHelper(current, newOne *Node) {
	if newOne.data <= current.data {
		//new data smaller than current node data
		if current.lChild != nil {
			//haven't reached location; keep going
			insertHelper(current.lChild, newOne)
		} else {
			//location found
			current.lChild = newOne
			return
		}
	} else {
		//new data larger than current node data
		if current.rChild != nil {
			//haven't reached location; keep going
			insertHelper(current.rChild, newOne)
		} else {
			//location found
			current.rChild = newOne
			return
		}
	}
}

// 3. Compute height of tree
func findHeightR(node *Node) int { // RECURSIVE
	if node == nil {
		return 0
	}
	heightLeft := findHeightR(node.lChild)
	heightRight := findHeightR(node.rChild)
	if heightLeft < heightRight {
		return 1 + heightRight
	}
	return 1 + heightLeft
}

func findHeight(node *Node) int {
	if node == nil {
		return 0
	}
	queue1 := []*Node{node}
	height := 0
	nodeCount := 1

	for nodeCount > 0 {
		height++
		newNodeCount := 0
		current := queue1[0]
		nodeCount--
		queue1 = queue1[1:]

		if current.lChild != nil {
			queue1 = append(queue1, current.lChild)
			newNodeCount++
		}
		if current.rChild != nil {
			queue1 = append(queue1, current.rChild)
			newNodeCount++
		}
		nodeCount = newNodeCount
	}
	return height
}

// 4. Print values - pre-order; Root-Left-Right (me before my children)
func preOrderR(node *Node) { // RECURSIVE
	if node == nil {
		return
	} else {
		fmt.Println(node.data)
		preOrder(node.lChild)
		preOrder(node.rChild)
	}
}

func preOrder(node *Node) { // ITERATIVE
	if node == nil {
		return
	}
	stack1 := []*Node{}
	stack1 = append(stack1, node)
	for len(stack1) > 0 {
		current := stack1[len(stack1)-1]
		fmt.Println(current.data)
		stack1 = stack1[:len(stack1)-1]
		if current.rChild != nil {
			stack1 = append(stack1, current.rChild)
		}
		if current.lChild != nil {
			stack1 = append(stack1, current.lChild)
		}
	}
}

// 5. Print values - in-order; Left-Root-Right, ascending
func inOrderR(node *Node) { // RECURSIVE
	if node == nil {
		return
	} else {
		inOrder(node.lChild)
		fmt.Println(node.data)
		inOrder(node.rChild)
	}
}

func inOrder(node *Node) { // ITERATIVE
	if node == nil {
		return
	}
	stack1 := []*Node{}
	current := node
	for current != nil || len(stack1) > 0 {
		if current != nil {
			stack1 = append(stack1, current)
			current = current.lChild
		} else if len(stack1) > 0 {
			current = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			fmt.Println(current.data)
			current = current.rChild
		}
	}
}

// 6. Print values - post-order; Left-Right-Root (my children before me)
func postOrderR(node *Node) { // RECURSIVE
	if node == nil {
		return
	} else {
		postOrder(node.lChild)
		postOrder(node.rChild)
		fmt.Println(node.data)
	}
}

func postOrder(node *Node) { // ITERATIVE
	if node == nil {
		return
	}
	stack1 := []*Node{}
	stack2 := []*Node{}
	stack1 = append(stack1, node)
	for len(stack1) > 0 {
		temp := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, temp)
		if temp.lChild != nil {
			stack1 = append(stack1, temp.lChild)
		}
		if temp.rChild != nil {
			stack1 = append(stack1, temp.rChild)
		}
	}

	for len(stack2) > 0 {
		fmt.Println(stack2[len(stack2)-1].data)
		stack2 = stack2[:len(stack2)-1]
	}
}

// 7. Print values - level-order
func levelOrder(node *Node) {
	if node == nil {
		return
	}

	queue1 := []*Node{node}
	for len(queue1) > 0 {
		current := queue1[0]
		fmt.Println(current.data)
		queue1 = queue1[1:]

		if current.lChild != nil {
			queue1 = append(queue1, current.lChild)
		}
		if current.rChild != nil {
			queue1 = append(queue1, current.rChild)
		}
	}

}

// 8. Delete a given value from tree
//func (n *BinaryTree) delete(v int) {
//}

func main() {
	//Initialize tree with items
	fmt.Println("-----start: 8,3,10,1,6,14,4,7,13-----")
	var myTree = newBinaryTree(8,3,10,1,6,14,4,7,13)
	fmt.Println(myTree.root)
	fmt.Println("-------search 2-------")
	fmt.Println(myTree.search(2))
	fmt.Println("-------search 10-------")
	fmt.Println(myTree.search(10))
	fmt.Println("-------height-------")
	fmt.Println(findHeight(myTree.root))
	fmt.Println("-------RecursiveHeight-------")
	fmt.Println(findHeightR(myTree.root))
	fmt.Println("-------preOrder-------")
	preOrder(myTree.root)
	fmt.Println("-------inOrder-------")
	inOrder(myTree.root)
	fmt.Println("-------postOrder-------")
	postOrder(myTree.root)
	fmt.Println("-------levelOrder-------")
	levelOrder(myTree.root)
}