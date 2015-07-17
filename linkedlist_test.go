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
	if list.Size() != 3 {
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

	v, err := list.Get(0)
	if v != 1 || err != nil {
		t.Fail()
	}
	v, err = list.Get(1)
	if v != 2 || err != nil {
		t.Fail()
	}
	v, err = list.Get(2)
	if v != 3 || err != nil {
		t.Fail()
	}
	v, err = list.Get(3)
	if v != nil || err == nil {
		t.Fail()
	}

	list.Add(4)
	list.Add(5)
	if list.Size() != 5 {
		t.Fail()
	}

	list.Remove(1)
	if list.Size() != 4 {
		t.Fail()
	}

	v, err = list.Get(1)
	if v != 3 || err != nil {
		t.Fail()
	}
	v, err = list.Get(2)
	if v != 4 || err != nil {
		t.Fail()
	}
	v, err = list.Get(3)
	if v != 5 || err != nil {
		t.Fail()
	}
	arrayInt = list.ToArrayOfType(reflect.TypeOf(0)).([]int)
	if len(arrayInt) != 4 {
		t.Fail()
	}
}
