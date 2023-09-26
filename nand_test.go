// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate_test

import (
	"github.com/chippyash/logicgate"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNand_Zero_Zero_Is_One(t *testing.T) {
	var a, b bitmap.Bitmap
	a.Grow(64)
	b.Grow(64)
	res := logicgate.Nand(a, b)
	assert.True(t, res.Contains(1))
}

func TestNand_Zero_One_Is_One(t *testing.T) {
	var a, b bitmap.Bitmap
	a.Grow(64)
	b.Grow(64)
	b.Ones()
	res := logicgate.Nand(a, b)
	assert.True(t, res.Contains(1))
}

func TestNand_One_Zero_Is_One(t *testing.T) {
	var a, b bitmap.Bitmap
	a.Grow(64)
	b.Grow(64)
	a.Ones()
	res := logicgate.Nand(a, b)
	assert.True(t, res.Contains(1))
}

func TestNand_One_One_Is_Zero(t *testing.T) {
	var a, b bitmap.Bitmap
	a.Grow(64)
	b.Grow(64)
	a.Ones()
	b.Ones()
	res := logicgate.Nand(a, b)
	assert.False(t, res.Contains(1))
}
