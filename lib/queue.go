package lib

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Push(value T) {
	q.elements = append(q.elements, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	var zero T

	if len(q.elements) == 0 {
		return zero, false
	}

	element := q.elements[0]

	q.elements = q.elements[1:]

	return element, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
