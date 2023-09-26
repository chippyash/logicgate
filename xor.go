// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import "github.com/kelindar/bitmap"

// Xor is the XOR Gate
// see https://en.wikipedia.org/wiki/XOR_gate
func Xor(a, b bitmap.Bitmap) bitmap.Bitmap {
	var sA bitmap.Bitmap
	a.Clone(&sA)
	sA.Xor(b)

	return sA
}
