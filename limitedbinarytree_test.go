// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"testing"
)

func TestLimitedbinarytreePreserveMin(t *testing.T) {
	tree := NewLimitedBinaryTree(func(e1 interface{}, e2 interface{}) int {
		return e1.(int) - e2.(int)
	}, 5, true)
	tree.Add(6)
	tree.Add(8)
	tree.Add(9)
	tree.Add(3)
	tree.Add(2)
	tree.Add(4)
	tree.Add(1)
	tree.Add(5)
	tree.Add(7)
	if tree.Size != 5 {
		t.Fail()
	}
	array := tree.ToArray()
	if len(array) != 5 {
		t.Fail()
	}
	for i, e := range array {
		v := e.(int)
		if v != i+1 {
			t.Fail()
		}
	}
}

func TestLimitedbinarytreePreserveMax(t *testing.T) {
	tree := NewLimitedBinaryTree(func(e1 interface{}, e2 interface{}) int {
		return e1.(int) - e2.(int)
	}, 5, false)
	tree.Add(6)
	tree.Add(8)
	tree.Add(9)
	tree.Add(3)
	tree.Add(2)
	tree.Add(4)
	tree.Add(1)
	tree.Add(5)
	tree.Add(7)
	if tree.Size != 5 {
		t.Fail()
	}
	array := tree.ToArray()
	if len(array) != 5 {
		t.Fail()
	}
	for i, e := range array {
		v := e.(int)
		if v != i+5 {
			t.Fail()
		}
	}
}
