package main

type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}
