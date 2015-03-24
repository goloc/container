// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import ()

type Counter struct {
	BinaryTree
}

type Count struct {
	Key string
	Val int
}

func NewCounter() *Counter {
	counter := new(Counter)
	counter.BinaryTree.CompareFunc = func(r1, r2 interface{}) int {
		count1 := r1.(*Count)
		count2 := r2.(*Count)
		if count1.Key == count2.Key {
			return 0
		}
		dif := count1.Val - count2.Val
		if dif == 0 {
			if count1.Key > count2.Key {
				return 1
			} else {
				return -1
			}
		} else {
			return dif
		}
	}
	return counter
}

func (counter *Counter) Incr(key string, val int) int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	newcount := &Count{Key: key, Val: val}
	node, parent, dif, err := counter.near(newcount)
	if err == nil && dif == 0 {
		newcount.Val += node.Element.(*Count).Val
		counter.remove(node, parent)
	}
	counter.add(newcount)
	return newcount.Val
}
