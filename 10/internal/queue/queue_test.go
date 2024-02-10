package queue

import "testing"

func TestQueueEmpty(t *testing.T) {
	var q Queue[int]

	r := q.Empty()

	if r != true {
		t.Fatalf(`q.Empty() = %t, want %t`, r, true)
	}
}

func TestQueueNotEmpty(t *testing.T) {
	var q Queue[int]

	q.Put(1)
	r := q.Empty()

	if r != false {
		t.Fatalf(`q.Empty() = %t, want %t`, r, false)
	}
}

func TestQueuePutAndGet(t *testing.T) {
	var q Queue[int]

	q.Put(2)
	x := q.Get()

	if x != 2 {
		t.Fatalf(`q.Get() = %d, want %d`, x, 2)
	}
}

func TestQueuePutAndGetMany(t *testing.T) {
	var q Queue[int]

	q.Put(1)
	q.Put(2)
	q.Put(3)
	x1 := q.Get()
	x2 := q.Get()
	x3 := q.Get()

	if x1 != 1 {
		t.Fatalf(`q.Get() = %d, want %d`, x1, 1)
	}
	if x2 != 2 {
		t.Fatalf(`q.Get() = %d, want %d`, x2, 2)
	}
	if x3 != 4 {
		t.Fatalf(`q.Get() = %d, want %d`, x3, 4)
	}
}
