// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"errors"
	"reflect"
	"sync"
)

type Set struct {
	Map   map[interface{}]bool
	mutex sync.RWMutex
}

func NewSet() *Set {
	s := new(Set)
	s.Map = make(map[interface{}]bool)
	return s
}

func (s *Set) Contains(element interface{}) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if len(s.Map) <= 0 {
		return false
	}
	for k, _ := range s.Map {
		if reflect.DeepEqual(k, element) {
			return true
		}
	}
	return false
}

func (s *Set) Add(element interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Map[element] = true
	return nil
}

func (s *Set) Get(element interface{}) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v := s.Map[element]
	if v == false {
		return nil, errors.New("Element not found")
	}
	return element, nil
}

func (s *Set) Remove(element interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	v := s.Map[element]
	if v == false {
		return errors.New("Element not found")
	}
	delete(s.Map, element)
	return nil
}

func (s *Set) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.Map)
}

func (s *Set) ToArray() []interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	array := make([]interface{}, len(s.Map))
	i := 0
	for k, v := range s.Map {
		array[i] = &KeyValue{Key: k, Value: v}
		i++
	}
	return array
}

func (s *Set) ToArrayOfType(elementType reflect.Type) interface{} {
	return s.ToArray()
}

func (s *Set) Visit(trait func(element interface{}, i int)) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	i := 0
	for k, _ := range s.Map {
		trait(k, i)
		i++
	}
}
