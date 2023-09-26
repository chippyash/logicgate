// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import "github.com/kelindar/bitmap"

// And is the AND gate
// see https://en.wikipedia.org/wiki/AND_gate
func And(a, b bitmap.Bitmap) bitmap.Bitmap {
	var sA bitmap.Bitmap
	a.Clone(&sA)
	sA.And(b)

	return sA
}
