/*
 * Copyright (c) Ashley Kitson, 2023. All rights reserved.
 * Licensed under the MIT license. See LICENSE file in the project root for details.
 *
 */

package utils_test

import (
	"fmt"
	"github.com/chippyash/logicgate/utils"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeAll(t *testing.T) {
	a := bitmap.Bitmap{}
	for i := uint32(0); i < 64; i += 2 {
		a.Set(i)
	}

	res := bitmap.Bitmap{}
	//rfunc transfers set bits from a to res
	rFunc := func(k uint32, v bool) {
		if v {
			res.Set(k)
		}
	}

	utils.RangeAll(a, rFunc)
	aStr := fmt.Sprintf("%064b", a[0])
	resStr := fmt.Sprintf("%064b", res[0])
	assert.Equal(t, aStr, resStr)
}
