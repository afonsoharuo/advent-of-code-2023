// Package queue provides a simple queue.
package queue

import (
	"slices"
)

// Queue provides a queue for elements of type T.
type Queue[T any] struct {
	elems []T
}

// Empty returns true if there are no elements in the queue.
func (q *Queue[T]) Empty() bool {
	return len(q.elems) == 0
}

// Put adds the element to the end of the queue.
func (q *Queue[T]) Put(elem T) {
	q.elems = append(q.elems, elem)
}

// Get removes and returns the first element of the queue.
func (q *Queue[T]) Get() T {
	elem := q.elems[0]
	q.elems = slices.Delete(q.elems, 0, 1)
	return elem
}
