package main

type Graph[T comparable, V any] struct {
	adjacncylist map[T][]*V
}

func NewGraph[T Comparable, V any]() *Graph[T, *T] {
	return &Graph[T, *T]{adjacncylist: make(map[T][]*T)}
}
