// Licensed under the MIT license. See LICENSE file in the project root for details.
// Copyright (c) Ashley Kitson, 2023. All rights reserved.

package logicgate

import "github.com/kelindar/bitmap"

// Nor is the NOR gate
// see https://en.wikipedia.org/wiki/NOR_gate
func Nor(a, b bitmap.Bitmap) bitmap.Bitmap {
	var sA bitmap.Bitmap
	sA = Nand(Nand(a, a), Nand(b, b))

	return Nand(sA, sA)
}
