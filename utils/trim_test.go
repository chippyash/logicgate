// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package utils_test

import (
	"github.com/chippyash/logicgate/utils"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrim(t *testing.T) {
	a := bitmap.Bitmap{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF}
	for i := uint32(0); i < 256; i++ {
		res := utils.Trim(a, i)
		c := uint32(res.Count())
		assert.Equal(t, i, c, "bitLen=%d, count=%d", i, c)
	}
}
