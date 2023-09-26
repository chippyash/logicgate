// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package utils

import (
	"github.com/chippyash/logicgate"
	"github.com/kelindar/bitmap"
)

// Trim returns a bitmap with the source bitmap ANDed with a bitmap of bitLen length
func Trim(a bitmap.Bitmap, bitLen uint32) bitmap.Bitmap {
	mask := bitmap.Bitmap{}
	mask.Grow(bitLen)
	for i := uint32(0); i < bitLen; i++ {
		mask.Set(i)
	}

	return logicgate.And(a, mask)
}
