package main

import (
	"fmt"
)

type BinaryTree[T Comparable] struct {
	Root *TreeNode[T]
}

func (tree *BinaryTree[T]) Insert(value T) {
	if tree.Root == nil {
		tree.Root = &TreeNode[T]{Value: value}
	} else {
		tree.Root.Insert(value)
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

func (root *TreeNode[T]) Insert(val T) {

	if val < root.Value {
		if root.Left == nil {
			root.Left = &TreeNode[T]{Value: val}
		} else {
			root.Left.Insert(val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode[T]{Value: val}
		} else {
			root.Right.Insert(val)
		}
	}

}

func (root *TreeNode[T]) traverseInOrder() {
	if root != nil {
		root.Right.traverseInOrder()
		fmt.Printf("%c\n", root.Value)
		root.Left.traverseInOrder()
	}
}

func (root *TreeNode[T]) breadthFirst() {

	if root != nil {
		fmt.Printf("%c\n", root.Value)
		root.Left.breadthFirst()
		root.Right.breadthFirst()
	}
}

func (root *TreeNode[T]) breadthFirstWithQueue() {

	if root == nil {
		return
	}

	qu := Queue[*TreeNode[T]]{}
	qu.Push(root)
	for qu.length() > 0 {
		curr := qu.Pop()
		fmt.Println(curr.Value)
		if curr.Left != nil {
			qu.Push(curr.Left)
		}
		if curr.Right != nil {
			qu.Push(curr.Right)
		}

	}
}

func (root *TreeNode[T]) depthFirstWithStack() {

	stack := Stack[*TreeNode[T]]{}

	stack.Push(root)

	for stack.length() > 0 {

		curr := stack.Pop()
		fmt.Println(curr.Value)
		if curr.Right != nil {
			stack.Push(curr.Right)
		}
		if curr.Left != nil {
			stack.Push(curr.Left)
		}

	}

}

func (root *TreeNode[T]) contains(val T) bool {
	if root == nil {
		return false
	}
	if root.Value == val {
		return true
	}

	if root.Right != nil {
		return root.Right.contains(val)
	}
	if root.Left != nil {
		return root.Left.contains(val)
	}

	return false
}
func (tree *BinaryTree[T]) contains(val T) bool {

	return tree.Root.contains(val)
}

func (root *TreeNode[T]) sum() T {
	if root == nil {
		return 0
	}
	return root.Value + root.Right.sum() + root.Left.sum()
}

func (tree *BinaryTree[T]) sum() T {
	return tree.Root.sum()
}
