/*
 * Copyright (c) Ashley Kitson, 2023. All rights reserved.
 * Licensed under the MIT license. See LICENSE file in the project root for details.
 *
 */

package utils

import (
	"github.com/kelindar/bitmap"
)

// Compare compares two bitmaps (a, b) and returns 0 if they are numerically equal, 1 if a is greater than b and -1 if a is less than b.
// Comparison is left to right, i.e. MSB to LSB.
func Compare(a, b bitmap.Bitmap) int {
	//nil bitmaps
	if a == nil && b == nil {
		return 0
	}
	if a == nil && b != nil {
		return -1
	}
	if a != nil && b == nil {
		return 1
	}
	//empty bitmaps that contain no 1s
	if a.Count() == 0 && b.Count() == 0 {
		return 0
	}
	//empty bitmaps that contain no length
	if len(a) == 0 && len(b) == 0 {
		return 0
	}
	//one empty bitmap
	if len(b) == 0 && len(a) != 0 && a.Count() > 1 {
		return 1
	}
	if len(a) == 0 && len(b) != 0 && b.Count() > 1 {
		return -1
	}

	//inequality
	l := len(a)
	if len(b) > l {
		l = len(b)
	}
	for i := uint32(l * 64); i > 0; i-- {
		if a.Contains(i-1) && !b.Contains(i-1) {
			return 1
		}
		if !a.Contains(i-1) && b.Contains(i-1) {
			return -1
		}
	}
	//equality
	return 0
}
