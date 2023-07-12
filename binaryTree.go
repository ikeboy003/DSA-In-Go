package main

type BinaryTree[T Comparable] struct {
	Root *TreeNode[T]
}

func newBinaryTree[T Comparable]() BinaryTree[T] {
	return BinaryTree[T]{}
}

func (b *BinaryTree[T]) Insert(value T) {

	newNode := &TreeNode[T]{Value: value}

	if b.Root == nil {
		b.Root = newNode
	} else {
		b.insertNode(value, b.Root)
	}

}

func (b *BinaryTree[T]) insertNode(value T, node *TreeNode[T]) {

}
