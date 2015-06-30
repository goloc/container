// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

type LimitedBinaryTree struct {
	BinaryTree
	Limit       int
	PreserveMin bool
}

func NewLimitedBinaryTree(compare Compare, limit int, preserveMin bool) *LimitedBinaryTree {
	tree := new(LimitedBinaryTree)
	tree.BinaryTree = *NewBinaryTree(compare)
	tree.Limit = limit
	tree.PreserveMin = preserveMin
	return tree
}

func (tree *LimitedBinaryTree) Add(element interface{}) error {
	tree.mutex.Lock()
	defer tree.mutex.Unlock()
	tree.check()
	if tree.size >= tree.Limit {
		if tree.PreserveMin {
			max, parent := tree.Head.right()
			if tree.CompareFunc(element, max.Element) >= 0 {
				return nil
			} else {
				tree.remove(max, parent)
			}
		} else {
			min, parent := tree.Head.left()
			if tree.CompareFunc(element, min.Element) <= 0 {
				return nil
			} else {
				tree.remove(min, parent)
			}
		}
		return tree.BinaryTree.add(element)
	} else {
		return tree.BinaryTree.add(element)
	}
}
