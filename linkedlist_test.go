// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"reflect"
	"testing"
)

func TestLinkedlist(t *testing.T) {
	list := Container(NewLinkedList())
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.GetSize() != 3 {
		t.Fail()
	}

	array := list.ToArray()
	if len(array) != 3 {
		t.Fail()
	}
	for i, v := range array {
		if v != i+1 {
			t.Fail()
		}
	}

	arrayInt := list.ToArrayOfType(reflect.TypeOf(0)).([]int)
	if len(arrayInt) != 3 {
		t.Fail()
	}
	for i, v := range arrayInt {
		if v != i+1 {
			t.Fail()
		}
	}

	v, err := list.Search(2)
	if v != 2 || err != nil {
		t.Fail()
	}
	v, err = list.Search(60)
	if v != nil || err == nil {
		t.Fail()
	}

	list.Add(4)
	list.Add(5)
	if list.GetSize() != 5 {
		t.Fail()
	}

	list.Remove(2)

	v, err = list.Search(2)
	if v != nil || err == nil {
		t.Fail()
	}
	arrayInt = list.ToArrayOfType(reflect.TypeOf(0)).([]int)
	if len(arrayInt) != 4 {
		t.Fail()
	}
}
