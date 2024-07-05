package hw1

type Queue[T any] struct {
  in, out []T
}

func NewQueue[T any]() Queue[T] {
  return Queue[T] {
    in:  []T{},
    out: []T{},
  }
}

func (q *Queue[T]) Size() int {
  return len(q.in) + len(q.out)
}

func (q *Queue[T]) Add(x T) {
  q.in = append(q.in, x)  
}

func (q *Queue[T]) Pop() (T, bool) {
  if len(q.out) == 0 {
    if len(q.in) == 0 {
      var zero T
      return zero, false
    }
    for i := len(q.in) - 1; i >= 0; i-- {
      q.out = append(q.out, q.in[i])
    }
    q.in = nil
  }
  x := q.out[len(q.out) - 1]
  q.out = q.out[:len(q.out)-1]
  return x, true
}
