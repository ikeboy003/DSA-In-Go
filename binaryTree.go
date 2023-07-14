package main

import "fmt"

type BinaryTree[T Comparable] struct {
	Root *TreeNode[T]
}

func (tree *BinaryTree[T]) Insert(value T) {
	if tree.Root == nil {
		tree.Root = &TreeNode[T]{Value: value}
	} else {
		InsertNode(tree.Root, value)
	}
}

func InsertNode[T Comparable](node *TreeNode[T], value T) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &TreeNode[T]{Value: value}
		} else {
			InsertNode[T](node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode[T]{Value: value}
		} else {
			InsertNode[T](node.Right, value)
		}
	}

}

func newBinaryTree[T Comparable]() *BinaryTree[T] {

	return &BinaryTree[T]{}
}

func (tree BinaryTree[T]) traverseInOrder() {
	if tree.Root != nil {
		traverseInOrderByNode[T](tree.Root)
	}
}

func traverseInOrderByNode[T Comparable](node *TreeNode[T]) {
	if node != nil {
		traverseInOrderByNode(node.Left)
		fmt.Printf("%d ", node.Value)
		traverseInOrderByNode(node.Right)
	}
}

func (tree BinaryTree[T]) traversePostOrder() {
	if tree.Root != nil {
		traversePostOrderByNode[T](tree.Root)
	}
}

func traversePostOrderByNode[T Comparable](node *TreeNode[T]) {
	if node != nil {
		traverseInOrderByNode(node.Left)
		traverseInOrderByNode(node.Right)
		fmt.Printf("%d ", node.Value)

	}
}

func (tree BinaryTree[T]) traversePreOrder() {
	if tree.Root != nil {
		traversePreeOrderByNode[T](tree.Root)
	}
}

func traversePreeOrderByNode[T Comparable](node *TreeNode[T]) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		traverseInOrderByNode(node.Left)
		traverseInOrderByNode(node.Right)

	}
}
