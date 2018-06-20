package design

import (
	"reflect"
	"testing"
)

func TestItemCollection_Iterator(t *testing.T) {
	items := []*Item{
		{
			ID:   1,
			Name: "1Good",
		},
		{
			ID:   6,
			Name: "6Bad",
		},
		{
			ID:   1024,
			Name: "1024Clu",
		},
		{
			ID:   2048,
			Name: "2048Bin",
		},
		{
			ID:   3,
			Name: "2048Bin",
		},
	}

	collection := &ItemCollection{itemList: items}
	evenFilter := func(item *Item) bool {
		return item.ID%2 == 0
	}

	oddFilter := func(item *Item) bool {
		return item.ID%2 != 0
	}

	var result []*Item
	iter := collection.Iterator(evenFilter)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}

	expect := []*Item{
		{
			ID:   6,
			Name: "6Bad",
		},
		{
			ID:   1024,
			Name: "1024Clu",
		},
		{
			ID:   2048,
			Name: "2048Bin",
		},
	}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("expect: %v, get: %v\n", expect, result)
	}

	result = nil
	iter = collection.Iterator(oddFilter)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}

	expect = []*Item{
		{
			ID:   1,
			Name: "1Good",
		},
		{
			ID:   3,
			Name: "2048Bin",
		},
	}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("expect: %v, get: %v\n", expect, result)
	}
}
