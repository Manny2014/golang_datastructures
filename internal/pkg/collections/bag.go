package collections

type Bag[T any] struct {
	items []T
}

func NewBag() BagCollection[any] {
	return &Bag[any]{}
}

func (b *Bag[any]) Add(item any) bool {
	b.items = append(b.items, item)
	return true
}

func (b *Bag[any]) CreateIterator() Iterator[any] {
	return &CollectionIterator[any]{
		items: b.items,
	}
}

func (b Bag[any]) IsEmpty() bool {
	return len(b.items) < 1
}

func (b Bag[any]) Size() int {
	return len(b.items)
}
