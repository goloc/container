// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"errors"
	"fmt"
	"reflect"
)

type LinkedList struct {
	Size     int
	Head     *LinkedListItem
	Ultimate *LinkedListItem
}

func NewLinkedList() *LinkedList {
	list := new(LinkedList)
	return list
}

type LinkedListItem struct {
	Element interface{}
	Next    *LinkedListItem
}

func (list *LinkedList) Add(elements ...interface{}) error {
	errMap := make(map[interface{}]error)
	for _, element := range elements {
		if err := list.add(element); err != nil {
			errMap[element] = err
		}
	}
	if len(errMap) > 0 {
		return errors.New(fmt.Sprintf("Errors has occured: %v", errMap))
	} else {
		return nil
	}
}

func (list *LinkedList) add(element interface{}) error {
	li := new(LinkedListItem)
	li.Element = element
	if list.Head == nil {
		list.Head = li
		list.Ultimate = list.Head
	} else {
		list.Ultimate.Next = li
		list.Ultimate = list.Ultimate.Next
	}
	list.Size++
	return nil
}

func (list *LinkedList) Search(element interface{}) (interface{}, error) {
	item, _, err := list.search(element)
	if err != nil {
		return nil, err
	}
	return item.Element, err
}

func (list *LinkedList) search(element interface{}) (*LinkedListItem, *LinkedListItem, error) {
	var prev *LinkedListItem
	item := list.Head
	for item != nil {
		if item.Element == element {
			return item, prev, nil
		}
		prev = item
		item = item.Next
	}
	return nil, nil, errors.New("Element not found")
}

func (list *LinkedList) Remove(element interface{}) error {
	item, parent, err := list.search(element)
	if err != nil {
		return err
	}
	if item == list.Head {
		list.Head = item.Next
	} else {
		if parent == nil {
			return errors.New("Parent node mandatory for non head node")
		}
		parent.Next = item.Next
	}
	list.Size--
	return nil
}

func (list *LinkedList) GetSize() int {
	return list.Size
}

func (list *LinkedList) ToArray() []interface{} {
	array := make([]interface{}, list.Size)
	list.Visit(func(element interface{}, i int) {
		array[i] = element
	})
	return array
}

func (list *LinkedList) ToArrayOfType(elementType reflect.Type) interface{} {
	var value reflect.Value
	arrayValue := reflect.MakeSlice(reflect.SliceOf(elementType), list.Size, list.Size)
	list.Visit(func(element interface{}, i int) {
		value = reflect.ValueOf(element)
		arrayValue.Index(i).Set(value)
	})
	return arrayValue.Interface()
}

func (list *LinkedList) Visit(trait func(element interface{}, i int)) {
	i := 0
	item := list.Head
	for item != nil {
		trait(item.Element, i)
		i++
		item = item.Next
	}
}
