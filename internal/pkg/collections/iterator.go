package collections

type CollectionIterator[T any] struct {
	index int
	items []T
}

func (c *CollectionIterator[any]) HasNext() bool {
	if c.index < len(c.items) {
		return true
	}
	return false

}
func (c *CollectionIterator[any]) GetNext() *any {
	if c.HasNext() {
		items := c.items[c.index]
		c.index++
		return &items
	}
	return nil
}
