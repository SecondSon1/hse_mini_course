package hw1

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	if q.Size() != 0 {
		t.Fatalf("New queue is not empty\n")
	}

	for i := 1; i <= 10; i++ {
		q.Add(i)
	}

	if q.Size() != 10 {
		t.Fatalf("Expected 10 elements, got %d\n", q.Size())
	}

	for i := 1; i <= 5; i++ {
		el, ok := q.Pop()
		if !ok {
			t.Fatalf("Could not pop out %d\n", i)
		}
		if el != i {
			t.Fatalf("Expected to pop out %d, got %d\n", i, el)
		}
	}

	for i := 11; i <= 15; i++ {
		q.Add(i)
	}

	if q.Size() != 10 {
		t.Fatalf("Expected 10 elements, got %d\n", q.Size())
	}

	for i := 6; i <= 15; i++ {
		el, ok := q.Pop()
		if !ok {
			t.Fatalf("Could not pop out %d\n", i)
		}
		if el != i {
			t.Fatalf("Expected to pop out %d, got %d\n", i, el)
		}
	}

	if q.Size() != 0 {
		t.Fatalf("Expected empty queue\n")
	}
}
