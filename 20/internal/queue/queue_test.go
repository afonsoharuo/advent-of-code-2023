package queue

import "testing"

func TestNewQueue(t *testing.T) {
	queue := Queue[int]{}
	if !queue.empty() {
		t.Fatalf(`new Queue not empty`)
	}
}

func TestQueueSize(t *testing.T) {
	queue := Queue[int]{}

	var size, expSize int

	expSize = 0
	size = queue.size()
	if size != expSize {
		t.Fatalf(`queue.Size() = %d, want %d`, size, expSize)
	}

	queue.put(1)
	queue.put(2)
	queue.put(3)

	expSize = 3
	size = queue.size()
	if size != expSize {
		t.Fatalf(`queue.Size() = %d, want %d`, size, expSize)
	}

	queue.get()

	expSize = 2
	size = queue.size()
	if size != expSize {
		t.Fatalf(`queue.Size() = %d, want %d`, size, expSize)
	}
}

func TestQueueValues(t *testing.T) {
	queue := Queue[int]{}

	var result, expResult int

	expResult = 9
	queue.put(9)
	queue.put(8)
	queue.put(7)
	result = queue.get()

	if result != expResult {
		t.Fatalf(`queue.Size() = %d, want %d`, result, expResult)
	}

	expResult = 8
	result = queue.get()

	if result != expResult {
		t.Fatalf(`queue.Size() = %d, want %d`, result, expResult)
	}

	expResult = 7
	queue.put(6)
	result = queue.get()
	if result != expResult {
		t.Fatalf(`queue.Size() = %d, want %d`, result, expResult)
	}
}
