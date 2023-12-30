package queue

import "slices"

type Queue[T any] struct {
	elems []T
}

func (q *Queue[T]) size() int {
	return len(q.elems)
}

func (q *Queue[T]) empty() bool {
	return len(q.elems) == 0
}

func (q *Queue[T]) put(elem T) {
	q.elems = append(q.elems, elem)
}

func (q *Queue[T]) get() T {
	elem := q.elems[0]
	q.elems = slices.Delete(q.elems, 0, 1)
	return elem
}
