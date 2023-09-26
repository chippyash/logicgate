// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package utils_test

import (
	"errors"
	"github.com/chippyash/logicgate/utils"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare_Two_Nil_Bitmaps_Are_Equal(t *testing.T) {
	var a, b bitmap.Bitmap
	res, err := utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)
}

func TestCompare_One_Nil_Bitmap_Throws_Error(t *testing.T) {
	var a bitmap.Bitmap
	b := bitmap.Bitmap{}

	_, err := utils.Compare(a, b)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, utils.ErrIncomparableBitmaps))
}

func TestCompare_Two_Empty_Bitmaps_Are_Equal(t *testing.T) {
	//two empty bitmaps
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	res, err := utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)

	//one empty bitmap, one bitmap == 0
	a.Grow(16)
	res, err = utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)
}

func TestCompare_One_Empty_Bitmaps_Are_Not_Equal(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(16)
	a.Set(1)
	res, err := utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, 1, res)
}

func TestCompare_A_Is_Bigger_Than_B_Simple_Case(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(16)
	b.Grow(16)
	a.Set(2)
	b.Set(1)
	res, err := utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, 1, res)
}

func TestCompare_A_Is_Smaller_Than_B_Simple_Case(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(16)
	b.Grow(16)
	a.Set(1)
	b.Set(2)
	res, err := utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, -1, res)
}

func TestCompare_A_Is_Bigger_Than_B_MultiBits(t *testing.T) {
	//a == 11110100
	//b == 11110010
	a, b := bitmap.Bitmap{0xF4}, bitmap.Bitmap{0xF2}

	res, err := utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, 1, res)
}

func TestCompare_B_Is_Bigger_Than_A_MultiBits(t *testing.T) {
	//a == 11110010
	//b == 11110100
	a, b := bitmap.Bitmap{0xF2}, bitmap.Bitmap{0xF4}

	res, err := utils.Compare(a, b)
	assert.NoError(t, err)
	assert.Equal(t, -1, res)
}
