// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"errors"
	"reflect"
	"sync"
)

type LinkedList struct {
	Head     *LinkedListItem
	size     int
	ultimate *LinkedListItem
	mutex    sync.RWMutex
}

func NewLinkedList() *LinkedList {
	list := new(LinkedList)
	return list
}

type LinkedListItem struct {
	Element interface{}
	Next    *LinkedListItem
}

func (list *LinkedList) check() {
	if list.ultimate == nil || list.size <= 0 {
		list.ultimate = nil
		list.size = 0
		item := list.Head
		for item != nil {
			list.size++
			list.ultimate = item
			item = item.Next
		}
	}
}

func (list *LinkedList) Contains(elementId interface{}) bool {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	list.check()
	id := elementId.(int)
	if id >= list.size || list.size <= 0 {
		return false
	}
	i := 0
	item := list.Head
	for item != nil {
		if i == id {
			return true
		}
		i++
		item = item.Next
	}
	return false
}

func (list *LinkedList) Add(element interface{}) error {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.check()
	li := new(LinkedListItem)
	li.Element = element
	if list.Head == nil {
		list.Head = li
		list.ultimate = list.Head
	} else {
		list.ultimate.Next = li
		list.ultimate = list.ultimate.Next
	}
	list.size++
	return nil
}

func (list *LinkedList) Get(elementId interface{}) (interface{}, error) {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	list.check()
	id := elementId.(int)
	if id >= list.size || list.size <= 0 {
		return nil, errors.New("Index is out ouf bound")
	}
	i := 0
	item := list.Head
	for item != nil {
		if i == id {
			return item.Element, nil
		}
		i++
		item = item.Next
	}
	return nil, errors.New("Element not found")
}

func (list *LinkedList) Remove(elementId interface{}) error {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.check()
	list.check()
	id := elementId.(int)
	if id >= list.size || list.size <= 0 {
		return errors.New("Index is out ouf bound")
	}
	i := 0
	var prev *LinkedListItem
	item := list.Head
	for item != nil {
		if i == id {
			prev.Next = item.Next
			list.size--
			return nil
		}
		i++
		prev = item
		item = item.Next
	}
	return errors.New("Element not found")
	return nil
}

func (list *LinkedList) Size() int {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	list.check()
	return list.size
}

func (list *LinkedList) ToArray() []interface{} {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	list.check()
	array := make([]interface{}, list.size)
	list.Visit(func(element interface{}, i int) {
		array[i] = element
	})
	return array
}

func (list *LinkedList) ToArrayOfType(elementType reflect.Type) interface{} {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	list.check()
	var value reflect.Value
	arrayValue := reflect.MakeSlice(reflect.SliceOf(elementType), list.size, list.size)
	list.Visit(func(element interface{}, i int) {
		value = reflect.ValueOf(element)
		arrayValue.Index(i).Set(value)
	})
	return arrayValue.Interface()
}

func (list *LinkedList) Visit(trait func(element interface{}, i int)) {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	i := 0
	item := list.Head
	for item != nil {
		trait(item.Element, i)
		i++
		item = item.Next
	}
}
