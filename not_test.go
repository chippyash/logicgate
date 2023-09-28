//go:build unit
// +build unit

// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate_test

import (
	"fmt"
	"github.com/chippyash/logicgate"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNot_Zero_Becomes_One(t *testing.T) {
	a := bitmap.Bitmap{}
	a.Grow(64)

	res := logicgate.Not(a)
	assert.True(t, res.Contains(1))
}

func TestNot_One_Becomes_Zero(t *testing.T) {
	a := bitmap.Bitmap{}
	a.Grow(64)
	a.Ones()

	res := logicgate.Not(a)
	assert.False(t, res.Contains(1))
}

func TestNot_Mixed(t *testing.T) {
	a := bitmap.Bitmap{}
	for i := uint32(0); i < 64; i += 2 {
		a.Set(i)
	}
	res := logicgate.Not(a)
	aStr := fmt.Sprintf("%064b", a[0])
	resStr := fmt.Sprintf("%064b", res[0])
	assert.Equal(t, 32, res.Count())
	assert.NotEqual(t, aStr, resStr)
	assert.Equal(t, "0101010101010101010101010101010101010101010101010101010101010101", aStr)
	assert.Equal(t, "1010101010101010101010101010101010101010101010101010101010101010", resStr)
}
