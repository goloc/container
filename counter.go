// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import (
	"sync"
)

type Counter struct {
	counts map[string]uint32
	max    uint32
	mutex  sync.RWMutex
}

func NewCounter() *Counter {
	c := new(Counter)
	c.counts = make(map[string]uint32)
	return c
}

func (c *Counter) Incr(key string) {
	c.mutex.Lock()
	c.counts[key]++
	count := c.counts[key]
	if count > c.max {
		c.max = count
	}
	c.mutex.Unlock()
}

func (c *Counter) GetMax() uint32 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.max
}

func (c *Counter) Visit(trait func(key string, count uint32, i int)) {
	c.mutex.RLock()
	i := 0
	for key, count := range c.counts {
		trait(key, count, i)
		i++
	}
	c.mutex.RUnlock()
}
