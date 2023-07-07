package main

type HashMap[K comparable, V any] struct {
	data map[K]V
}

func (h *HashMap[K, V]) Put(k K, v V) {
	h.data[k] = v
}

func (h *HashMap[K, V]) Get(k K) (v V, ok bool) {
	v, ok = h.data[k]
	return v, ok
}
