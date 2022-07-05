package collections

type Stack[T any] struct {
	items []T
}

func NewStack() StackCollection[any] {
	return &Stack[any]{}
}

func (b *Stack[any]) Push(item any) bool {
	b.items = append(b.items, item)
	return true
}

func (b *Stack[any]) Pop() *any {

	if b.Size() < 1 {
		return nil
	}

	i := len(b.items) - 1
	popped := b.items[i]
	b.items = b.items[:i]
	return &popped
}

func (b *Stack[any]) CreateIterator() Iterator[any] {
	return &CollectionIterator[any]{
		items: b.items,
	}
}

func (b Stack[any]) IsEmpty() bool {
	return len(b.items) < 1
}

func (b Stack[any]) Size() int {
	return len(b.items)
}
