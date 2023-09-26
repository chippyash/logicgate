/*
 * Copyright (c) Ashley Kitson, 2023. All rights reserved.
 * Licensed under the MIT license. See LICENSE file in the project root for details.
 *
 */

package utils

import "github.com/kelindar/bitmap"

// RangeAll iterates over all the bits in the bitmap.
// For the rare instances where bitmap.Range is insufficient
func RangeAll(src bitmap.Bitmap, fn func(k uint32, v bool)) {
	l := uint32(len(src))
	for blkAt := uint32(0); blkAt < l; blkAt++ {
		blk := src[blkAt]
		bit := uint64(0x1)
		for bitAt := uint32(0); bitAt < 64; bitAt++ {
			k := blkAt*64 + bitAt
			v := (blk & bit) != 0
			fn(k, v)
			bit <<= 1
		}
	}
}
