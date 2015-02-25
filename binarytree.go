// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"errors"
)

type Compare func(interface{}, interface{}) int

type BinaryTree struct {
	CompareFunc Compare
	Head        *BinaryTreeNode
	Size        int
}

func NewBinaryTree(compare Compare) *BinaryTree {
	tree := new(BinaryTree)
	tree.CompareFunc = compare
	return tree
}

type BinaryTreeNode struct {
	Element     interface{}
	Left, Right *BinaryTreeNode
	Size        int
}

func NewBinaryTreeNode(element interface{}) *BinaryTreeNode {
	node := new(BinaryTreeNode)
	node.Element = element
	node.Size = 1
	return node
}

func (tree *BinaryTree) near(element interface{}) (*BinaryTreeNode, *BinaryTreeNode, int) {
	var dif int
	var parent *BinaryTreeNode
	current := tree.Head
	if current == nil {
		return nil, nil, 0
	}
	for {
		dif = tree.CompareFunc(element, current.Element)
		if dif == 0 {
			return current, parent, dif
		} else if dif > 0 {
			if current.Right == nil {
				parent = current
				return current, parent, dif
			} else {
				parent = current
				current = current.Right
			}
		} else {
			if current.Left == nil {
				return current, parent, dif
			} else {
				parent = current
				current = current.Left
			}
		}
	}
}

func (tree *BinaryTree) Add(element interface{}) (*BinaryTreeNode, error) {
	if tree.Head == nil {
		tree.Head = NewBinaryTreeNode(element)
		tree.Size++
		return tree.Head, nil
	} else {
		currentNode, _, dif := tree.near(element)
		if dif == 0 {
			return currentNode, errors.New("The input element is equal to an element in tree")
		} else if dif > 0 {
			currentNode.Right = NewBinaryTreeNode(element)
			tree.Size++
			return currentNode.Right, nil
		} else {
			currentNode.Left = NewBinaryTreeNode(element)
			tree.Size++
			return currentNode.Left, nil
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

func (tree *BinaryTree) Search(element interface{}) (interface{}, error) {
	node, _, dif := tree.near(element)
	if dif == 0 {
		return node.Element, nil
	} else {
		return nil, errors.New("Element not found")
	}
}

func (tree *BinaryTree) Remove(element interface{}) error {
	node, parent, dif := tree.near(element)
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
	tree.Size--
	return nil
}

func (tree *BinaryTree) ToArray() []interface{} {
	array := make([]interface{}, tree.Size)
	tree.Visit(func(element interface{}, i int) {
		array[i] = element
	})
	return array
}

func (tree *BinaryTree) Visit(trait func(element interface{}, i int)) {
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

func (node *BinaryTreeNode) min() (*BinaryTreeNode, *BinaryTreeNode) {
	var parent *BinaryTreeNode
	if node != nil {
		for node.Left != nil {
			parent = node
			node = node.Left
		}
	}
	return node, parent
}

func (node *BinaryTreeNode) max() (*BinaryTreeNode, *BinaryTreeNode) {
	var parent *BinaryTreeNode
	if node != nil {
		for node.Right != nil {
			parent = node
			node = node.Right
		}
	}
	return node, parent
}
