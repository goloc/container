// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"reflect"
)

type Container interface {
	Contains(interface{}) bool
	Add(interface{}) error
	Get(interface{}) (interface{}, error)
	Remove(interface{}) error
	Size() int
	ToArray() []interface{}
	ToArrayOfType(reflect.Type) interface{}
	Visit(func(interface{}, int))
}
