package main

type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

type TreeNode[T Comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type Comparable interface {
	int | float32 | rune
}
