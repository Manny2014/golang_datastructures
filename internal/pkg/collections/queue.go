package collections

type Queue[T any] struct {
	items []T
}

func NewQueue() QueueCollection[any] {
	return &Queue[any]{}
}

func (b *Queue[any]) Enqueue(item any) bool {
	b.items = append(b.items, item)
	return true
}

func (b *Queue[any]) Dequeue() *any {

	if b.Size() < 1 {
		return nil
	}

	popped := b.items[0]
	b.items = append(b.items[:0], b.items[1:]...)
	return &popped
}

func (b *Queue[any]) CreateIterator() Iterator[any] {
	return &CollectionIterator[any]{
		items: b.items,
	}
}

func (b Queue[any]) IsEmpty() bool {
	return len(b.items) < 1
}

func (b Queue[any]) Size() int {
	return len(b.items)
}
