// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package container

import "testing"

func TestSet(t *testing.T) {
	set := Container(NewSet())
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Fail()
	}
	v, err := set.Get(1)
	if v != 1 || err != nil {
		t.Fail()
	}
	v, err = set.Get(2)
	if v != 2 || err != nil {
		t.Fail()
	}
	v, err = set.Get(3)
	if v != 3 || err != nil {
		t.Fail()
	}
	v, err = set.Get(4)
	if err == nil {
		t.Fail()
	}

	if set.Contains(1) != true {
		t.Fail()
	}
	if set.Contains(2) != true {
		t.Fail()
	}
	if set.Contains(3) != true {
		t.Fail()
	}
	if set.Contains(4) != false {
		t.Fail()
	}

	set.Remove(2)

	v, err = set.Get(2)
	if err == nil {
		t.Fail()
	}
	if set.Contains(2) != false {
		t.Fail()
	}

}
