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
}
