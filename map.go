// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"errors"
	"reflect"
	"sync"
)

type Map struct {
	Map   map[interface{}]interface{}
	mutex sync.RWMutex
}

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

func NewMap() *Map {
	m := new(Map)
	m.Map = make(map[interface{}]interface{})
	return m
}

func (m *Map) Contains(key interface{}) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	_, ok := m.Map[key]
	return ok
}

func (m *Map) Add(keyValue interface{}) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	kv, ok := keyValue.(*KeyValue)
	if ok {
		m.Map[kv.Key] = kv.Value
		return nil
	}
	return errors.New("Parameter must be of type KeyValue or *KeyValue")
}

func (m *Map) Get(key interface{}) (interface{}, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	v, ok := m.Map[key]
	if ok == false {
		return nil, errors.New("Element not found")
	}
	return &KeyValue{Key: key, Value: v}, nil
}

func (m *Map) Remove(key interface{}) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.Map[key]
	if ok == false {
		return errors.New("Element not found")
	}
	delete(m.Map, key)
	return nil
}

func (m *Map) Size() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.Map)
}

func (m *Map) ToArray() []interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	array := make([]interface{}, len(m.Map))
	i := 0
	for k, v := range m.Map {
		array[i] = &KeyValue{Key: k, Value: v}
		i++
	}
	return array
}

func (m *Map) ToArrayOfType(elementType reflect.Type) interface{} {
	return m.ToArray()
}

func (m *Map) Visit(trait func(element interface{}, i int)) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	i := 0
	for k, v := range m.Map {
		trait(&KeyValue{Key: k, Value: v}, i)
		i++
	}
}
