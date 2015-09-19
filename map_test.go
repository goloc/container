// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import "testing"

func TestMap(t *testing.T) {
	m := Container(NewMap())
	m.Add(&KeyValue{1, "1"})
	m.Add(&KeyValue{2, "2"})
	m.Add(&KeyValue{3, "3"})

	if m.Size() != 3 {
		t.Fail()
	}
	v, err := m.Get(1)
	if v.(*KeyValue).Value != "1" || err != nil {
		t.Fail()
	}
	v, err = m.Get(2)
	if v.(*KeyValue).Value != "2" || err != nil {
		t.Fail()
	}
	v, err = m.Get(3)
	if v.(*KeyValue).Value != "3" || err != nil {
		t.Fail()
	}
	v, err = m.Get(4)
	if err == nil {
		t.Fail()
	}

	if m.Contains(1) != true {
		t.Fail()
	}
	if m.Contains(2) != true {
		t.Fail()
	}
	if m.Contains(3) != true {
		t.Fail()
	}
	if m.Contains(4) != false {
		t.Fail()
	}

	m.Remove(2)

	v, err = m.Get(2)
	if err == nil {
		t.Fail()
	}
	if m.Contains(2) != false {
		t.Fail()
	}

}
