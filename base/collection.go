package base

type EmptyType struct{}

type Collection[T comparable] struct {
	Elements map[T]EmptyType
}

type Slicer[T any] interface {
	Size() int
	Slice() []T
}

type Container[T comparable] interface {
	Contains(element T) bool
}

func NewCollection[T comparable]() *Collection[T] {
	return &Collection[T]{
		Elements: make(map[T]EmptyType),
	}
}

func (c *Collection[T]) Size() int {
	return len(c.Elements)
}

func (c *Collection[T]) Slice() []T {
	slice := make([]T, 0, len(c.Elements))
	for element := range c.Elements {
		slice = append(slice, element)
	}
	return slice
}

func (c *Collection[T]) Contains(element T) bool {
	_, contains := c.Elements[element]
	return contains
}
