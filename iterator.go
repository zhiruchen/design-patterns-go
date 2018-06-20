package design

type Item struct {
	ID   int64
	Name string
}

type Iterator interface {
	HasNext() bool
	Next() *Item
}

type ItemFilter func(item *Item) bool

type ItemCollection struct {
	itemList []*Item
}

func (c *ItemCollection) items() []*Item {
	return c.itemList
}

func (c *ItemCollection) Iterator(filter ItemFilter) Iterator {
	return &ItemCollectionIterator{collection: c, itemFilter: filter, currentPos: -1}
}

type ItemCollectionIterator struct {
	collection *ItemCollection
	itemFilter ItemFilter
	currentPos int
}

func (iter *ItemCollectionIterator) HasNext() bool {
	return iter.findNextPos() != -1
}

func (iter *ItemCollectionIterator) Next() *Item {
	iter.currentPos = iter.findNextPos()

	if iter.currentPos != -1 {
		return iter.collection.items()[iter.currentPos]
	}

	return nil
}

func (iter *ItemCollectionIterator) findNextPos() int {
	items := iter.collection.items()
	curPos := iter.currentPos
	itemLen := len(items)

	for {
		curPos++
		if curPos >= itemLen {
			return -1
		}

		if iter.itemFilter(items[curPos]) {
			return curPos
		}
	}

	return curPos
}
