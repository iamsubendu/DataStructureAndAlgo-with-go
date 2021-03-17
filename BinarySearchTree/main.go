// Binary Tree
// A binary tree is a type of tree in which each node has
// at most two children (0,
// 1 or 2) which are referred as left child and right child.

// Binary Tree
// A binary tree is a type of tree in which each node has at
// most two children (0,
// 1 or 2) which are referred as left child and right child.

// A binary search tree (BST) is a binary tree on which nodes
// are ordered in the
// following way:
// 1. The key in the left subtree is less than the key in its
// parent node.
// 2. The key in the right subtree is greater or equal the key
// in its parent node.

// Binary Search Tree ADT Operations
// · Insert(k): Insert an element k into the tree.
// · Delete(k): Delete an element k from the tree.
// · Search(k): Search a particular value k into the tree if
// it is present or not.
// · FindMax(): Find the maximum value stored in the tree.
// · FindMin(): Find the minimum value stored in the tree.

package main

import "fmt"

var count int

// Node is for the binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the binary search tree
// go to right if larger and go to left if smaller
func (n *Node) Insert(k int) {
	if n.Key < k {
		//right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
			//if their is a existing child than recalling the method again on the child
		}
	} else if n.Key > k {
		//left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}

}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	//a node with its address
	fmt.Println(tree)
	//tree.Insert(200)
	//tree.Insert(300)
	//fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(76))
	fmt.Println(count)

}
