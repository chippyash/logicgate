// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import (
	"github.com/kelindar/bitmap"
)

// Nand is the NAND gate
// see https://en.wikipedia.org/wiki/NAND_gate
func Nand(a, b bitmap.Bitmap) bitmap.Bitmap {
	var a1 bitmap.Bitmap
	a.Clone(&a1)
	a1.And(b)

	return Not(a1)
}
