// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"reflect"
	"testing"
)

func TestBinarytree(t *testing.T) {
	tree := Container(NewBinaryTree(func(e1 interface{}, e2 interface{}) int {
		return e1.(int) - e2.(int)
	}))
	tree.Add(6)
	tree.Add(8)
	tree.Add(9)
	tree.Add(3)
	tree.Add(2)
	tree.Add(4)
	tree.Add(1)
	tree.Add(5)
	tree.Add(7)
	if tree.GetSize() != 9 {
		t.Fail()
	}

	array := tree.ToArray()
	if len(array) != 9 {
		t.Fail()
	}
	for i, v := range array {
		if v != i+1 {
			t.Fail()
		}
	}

	arrayInt := tree.ToArrayOfType(reflect.TypeOf(0)).([]int)
	if len(arrayInt) != 9 {
		t.Fail()
	}
	for i, v := range arrayInt {
		if v != i+1 {
			t.Fail()
		}
	}

	v, err := tree.Search(5)
	if v != 5 || err != nil {
		t.Fail()
	}
	v, err = tree.Search(60)
	if v != nil || err == nil {
		t.Fail()
	}

	tree.Add(15)
	tree.Add(11)
	tree.Add(10)
	tree.Add(14)
	tree.Add(12)
	tree.Add(13)
	if tree.GetSize() != 15 {
		t.Fail()
	}

	err = tree.Add(8)
	if err == nil {
		t.Fail()
	}
	if tree.GetSize() != 15 {
		t.Fail()
	}

	v, err = tree.Search(8)
	if v != 8 || err != nil {
		t.Fail()
	}
	v, err = tree.Search(6)
	if v != 6 || err != nil {
		t.Fail()
	}
	v, err = tree.Search(1)
	if v != 1 || err != nil {
		t.Fail()
	}
	v, err = tree.Search(15)
	if v != 15 || err != nil {
		t.Fail()
	}

	tree.Remove(8)
	tree.Remove(6)
	tree.Remove(1)
	tree.Remove(15)

	v, err = tree.Search(8)
	if v != nil || err == nil {
		t.Fail()
	}
	v, err = tree.Search(6)
	if v != nil || err == nil {
		t.Fail()
	}
	v, err = tree.Search(1)
	if v != nil || err == nil {
		t.Fail()
	}
	v, err = tree.Search(15)
	if v != nil || err == nil {
		t.Fail()
	}
	if tree.GetSize() != 11 {
		t.Fail()
	}
	arrayInt = tree.ToArrayOfType(reflect.TypeOf(0)).([]int)
	if len(arrayInt) != 11 {
		t.Fail()
	}
}
