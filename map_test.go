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
}
