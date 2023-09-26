// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import "github.com/kelindar/bitmap"

// Or is the OR gate
// see https://en.wikipedia.org/wiki/OR_gate
func Or(a, b bitmap.Bitmap) bitmap.Bitmap {
	var sA bitmap.Bitmap
	a.Clone(&sA)
	sA.Or(b)

	return sA
}
