// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"errors"
	"reflect"
	"sync"
)

type Compare func(interface{}, interface{}) int

type BinaryTree struct {
	compareFunc Compare
	Head        *BinaryTreeNode
	size        int
	mutex       sync.RWMutex
}

func NewBinaryTree(compare Compare) *BinaryTree {
	tree := new(BinaryTree)
	tree.compareFunc = compare
	return tree
}

type BinaryTreeNode struct {
	Element     interface{}
	Left, Right *BinaryTreeNode
}

func NewBinaryTreeNode(element interface{}) *BinaryTreeNode {
	node := new(BinaryTreeNode)
	node.Element = element
	return node
}

func (tree *BinaryTree) SetCompareFunc(compare Compare) {
	tree.compareFunc = compare
}

func (tree *BinaryTree) near(element interface{}) (*BinaryTreeNode, *BinaryTreeNode, int, error) {
	var dif int
	var parent *BinaryTreeNode
	current := tree.Head
	if current == nil {
		return nil, nil, 1, errors.New("No head")
	}
	for {
		dif = tree.compareFunc(element, current.Element)
		if dif == 0 {
			return current, parent, dif, nil
		} else if dif > 0 {
			if current.Right == nil {
				parent = current
				return current, parent, dif, nil
			} else {
				parent = current
				current = current.Right
			}
		} else {
			if current.Left == nil {
				return current, parent, dif, nil
			} else {
				parent = current
				current = current.Left
			}
		}
	}
}

func (tree *BinaryTree) check() {
	if tree.size <= 0 {
		if tree.Head == nil {
			tree.size = 0
		} else {
			tree.size = tree.count(tree.Head)
		}
	}
}

func (tree *BinaryTree) count(node *BinaryTreeNode) int {
	i := 1
	if node.Left != nil {
		i += tree.count(node.Left)
	}
	if node.Right != nil {
		i += tree.count(node.Right)
	}
	return i
}

func (tree *BinaryTree) Contains(element interface{}) bool {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	_, _, dif, err := tree.near(element)
	if err == nil && dif == 0 {
		return true
	}
	return false
}

func (tree *BinaryTree) Add(element interface{}) error {
	tree.mutex.Lock()
	defer tree.mutex.Unlock()
	tree.check()
	return tree.add(element)
}

func (tree *BinaryTree) add(element interface{}) error {
	if tree.Head == nil {
		tree.Head = NewBinaryTreeNode(element)
		tree.size++
		return nil
	} else {
		currentNode, _, dif, err := tree.near(element)
		if err != nil {
			return err
		}
		if dif == 0 {
			currentNode.Element = element
			return nil
		} else if dif > 0 {
			currentNode.Right = NewBinaryTreeNode(element)
			tree.size++
			return nil
		} else {
			currentNode.Left = NewBinaryTreeNode(element)
			tree.size++
			return nil
		}
	}
}

func (tree *BinaryTree) join(node1 *BinaryTreeNode, node2 *BinaryTreeNode) *BinaryTreeNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	node1.Right = tree.join(node1.Right, node2)
	return node1
}

func (tree *BinaryTree) Get(element interface{}) (interface{}, error) {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	node, _, dif, err := tree.near(element)
	if err != nil {
		return nil, err
	}
	if dif == 0 {
		return node.Element, nil
	} else {
		return nil, errors.New("Element not found")
	}
}

func (tree *BinaryTree) Remove(element interface{}) error {
	tree.mutex.Lock()
	defer tree.mutex.Unlock()
	tree.check()
	node, parent, dif, err := tree.near(element)
	if err != nil {
		return err
	}
	if dif == 0 {
		return tree.remove(node, parent)
	} else {
		return errors.New("Element not found")
	}
}

func (tree *BinaryTree) remove(node *BinaryTreeNode, parent *BinaryTreeNode) error {
	if node == nil {
		return errors.New("No node to delete")
	}
	newNode := tree.join(node.Left, node.Right)
	if node == tree.Head {
		tree.Head = newNode
	} else {
		if parent == nil {
			return errors.New("Parent node mandatory for non head node")
		}
		if parent.Left == node {
			parent.Left = newNode
		} else if parent.Right == node {
			parent.Right = newNode
		} else {
			return errors.New("No parent relation on input parameters")
		}
	}
	tree.size--
	return nil
}

func (tree *BinaryTree) Size() int {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	tree.check()
	return tree.size
}

func (tree *BinaryTree) ToArray() []interface{} {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	tree.check()
	array := make([]interface{}, tree.size)
	tree.Visit(func(element interface{}, i int) {
		array[i] = element
	})
	return array
}

func (tree *BinaryTree) ToArrayOfType(elementType reflect.Type) interface{} {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	tree.check()
	var value reflect.Value
	arrayValue := reflect.MakeSlice(reflect.SliceOf(elementType), tree.size, tree.size)
	tree.Visit(func(element interface{}, i int) {
		value = reflect.ValueOf(element)
		arrayValue.Index(i).Set(value)
	})
	return arrayValue.Interface()
}

func (tree *BinaryTree) Visit(trait func(element interface{}, i int)) {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	if tree.Head != nil {
		tree.visit(tree.Head, 0, trait)
	}
}

func (tree *BinaryTree) visit(node *BinaryTreeNode, i int, trait func(element interface{}, i int)) int {
	if node.Left != nil {
		i = tree.visit(node.Left, i, trait)
	}
	trait(node.Element, i)
	i++
	if node.Right != nil {
		i = tree.visit(node.Right, i, trait)
	}
	return i
}

func (tree *BinaryTree) Left() (interface{}, error) {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	current := tree.Head
	if current == nil {
		return nil, errors.New("No head")
	}
	element, _ := current.left()
	return element, nil
}

func (node *BinaryTreeNode) left() (*BinaryTreeNode, *BinaryTreeNode) {
	var parent *BinaryTreeNode
	if node != nil {
		for node.Left != nil {
			parent = node
			node = node.Left
		}
	}
	return node, parent
}

func (tree *BinaryTree) Right() (interface{}, error) {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	current := tree.Head
	if current == nil {
		return nil, errors.New("No head")
	}
	element, _ := current.right()
	return element, nil
}

func (node *BinaryTreeNode) right() (*BinaryTreeNode, *BinaryTreeNode) {
	var parent *BinaryTreeNode
	if node != nil {
		for node.Right != nil {
			parent = node
			node = node.Right
		}
	}
	return node, parent
}
