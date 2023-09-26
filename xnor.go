// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import (
	"github.com/kelindar/bitmap"
)

// Xnor is the XNOR gate.
// see https://en.wikipedia.org/wiki/XNOR_gate
func Xnor(a, b bitmap.Bitmap) bitmap.Bitmap {
	var sA, sD bitmap.Bitmap
	sA = Nand(a, b)
	sD = Nand(Nand(sA, b), Nand(a, sA))

	return Nand(sD, sD)
}
