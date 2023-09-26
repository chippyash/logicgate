/*
 * Copyright (c) Ashley Kitson, 2023. All rights reserved.
 * Licensed under the MIT license. See LICENSE file in the project root for details.
 *
 */

package utils

import (
	"errors"
	"github.com/chippyash/logicgate"
	"github.com/kelindar/bitmap"
)

var ErrIncomparableBitmaps = errors.New("incomparable bitmaps")

func Compare(a, b bitmap.Bitmap) (int, error) {
	//nil bitmaps
	if a == nil && b == nil {
		return 0, nil
	}
	if a == nil || b == nil {
		return 0, ErrIncomparableBitmaps
	}
	//empty bitmaps
	// empty bitmaps are bitmaps that contain no 1s or have no length
	if a.Count() == 0 && b.Count() == 0 {
		return 0, nil
	}

	//equality
	res := logicgate.Xnor(a, b)
	if res.Contains(1) {
		return 0, nil
	}

	//inequality
	mA, okA := a.Max()
	mB, okB := b.Max()
	if okA && !okB {
		//b is empty
		return 1, nil
	}
	if !okA && okB {
		//a is empty
		return -1, nil
	}

	if mA == mB {
		var aa, bb bitmap.Bitmap
		a.Clone(&aa)
		b.Clone(&bb)
		aa.Remove(mA)
		bb.Remove(mA)
		return Compare(aa, bb)
	}
	if mA > mB {
		return 1, nil
	}

	return -1, nil
}
