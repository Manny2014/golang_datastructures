package collections

type Iterator[T any] interface {
	HasNext() bool
	GetNext() *T
}

type Collection[T any] interface {
	CreateIterator() Iterator[T]
	IsEmpty() bool
	Size() int
}

type BagCollection[T any] interface {
	Collection[T]
	Add(item T) bool
}

type StackCollection[T any] interface {
	Collection[T]
	Push(item T) bool
	Pop() *T
}

type QueueCollection[T any] interface {
	Collection[T]
	Enqueue(item T) bool
	Dequeue() *T
}
