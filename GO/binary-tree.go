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
	// empty tree
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
func findHeight(node *Node) int {
	if node == nil {
		return 0
	}
	heightLeft := findHeight(node.lChild)
	heightRight := findHeight(node.rChild)
	if heightLeft < heightRight {
		return 1 + heightRight
	}
	return 1 + heightLeft
}

// 4. Print values - pre-order; Root-Left-Right (me before my children)
func preOrder(node *Node) {
	if node == nil {
		return
	} else {
		fmt.Println(node.data)
		preOrder(node.lChild)
		preOrder(node.rChild)
	}
}

// 5. Print values - in-order; Left-Root-Right, ascending
func inOrder(node *Node) {
	if node == nil {
		return
	} else {
		inOrder(node.lChild)
		fmt.Println(node.data)
		inOrder(node.rChild)
	}
}
// 6. Print values - post-order; Left-Right-Root (my children before me)
func postOrder(node *Node) {
	if node == nil {
		return
	} else {
		postOrder(node.lChild)
		postOrder(node.rChild)
		fmt.Println(node.data)
	}
}

// 7. Print values - level-order

// 8. Delete a given value from tree

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
	fmt.Println("-------preOrder-------")
	preOrder(myTree.root)
	fmt.Println("-------inOrder-------")
	inOrder(myTree.root)
	fmt.Println("-------postOrder-------")
	postOrder(myTree.root)
}