//go:build unit
// +build unit

// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package utils_test

import (
	"github.com/chippyash/logicgate/utils"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare_Two_Nil_Bitmaps_Are_Equal(t *testing.T) {
	var a, b bitmap.Bitmap
	res := utils.Compare(a, b)
	assert.Equal(t, 0, res, "a: %016b b: %016b", a, b)
}

func TestCompare_Two_Empty_Bitmaps_Are_Equal(t *testing.T) {
	//two empty bitmaps
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	res := utils.Compare(a, b)
	assert.Equal(t, 0, res, "a: %016b b: %016b", a, b)

	//one empty bitmap, one bitmap == 0
	a.Grow(16)
	res = utils.Compare(a, b)
	assert.Equal(t, 0, res, "a: %016b b: %016b", a, b)
}

func TestCompare_One_Empty_Bitmaps_Are_Not_Equal(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(16)
	a.Set(1)
	res := utils.Compare(a, b)
	assert.Equal(t, 1, res, "a: %016b b: %016b", a, b)
}

func TestCompare_A_Is_Bigger_Than_B_Simple_Case(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(16)
	b.Grow(16)
	a.Set(2)
	b.Set(1)
	res := utils.Compare(a, b)
	assert.Equal(t, 1, res, "a: %016b b: %016b", a, b)
}

func TestCompare_A_Is_Smaller_Than_B_Simple_Case(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(16)
	b.Grow(16)
	a.Set(1)
	b.Set(2)
	res := utils.Compare(a, b)
	assert.Equal(t, -1, res, "a: %016b b: %016b", a, b)
}

func TestCompare_A_Is_Bigger_Than_B_MultiBits(t *testing.T) {
	//a == 11110100
	//b == 11110010
	a, b := bitmap.Bitmap{0xF4}, bitmap.Bitmap{0xF2}

	res := utils.Compare(a, b)
	assert.Equal(t, 1, res, "a: %016b b: %016b", a, b)
}

func TestCompare_B_Is_Bigger_Than_A_MultiBits(t *testing.T) {
	//a == 11110010
	//b == 11110100
	a, b := bitmap.Bitmap{0xF2}, bitmap.Bitmap{0xF4}

	res := utils.Compare(a, b)
	assert.Equal(t, -1, res, "a: %016b b: %016b", a, b)
}

func FuzzCompare(f *testing.F) {
	f.Add(uint64(0xFF), uint64(0xFF))
	f.Fuzz(func(t *testing.T, a, b uint64) {
		out := utils.Compare(bitmap.Bitmap{a}, bitmap.Bitmap{b})
		if a == b && out != 0 {
			t.Errorf("Expected output to be 0, input was %d: %016b, %d: %016b output was: %d", a, a, b, b, out)
		}
		if a < b && out != -1 {
			t.Errorf("Expected output to be -1, input was %d: %016b, %d: %016b output was: %d", a, a, b, b, out)
		}
		if a > b && out != 1 {
			t.Errorf("Expected output to be 1, input was %d: %016b, %d: %016b output was: %d", a, a, b, b, out)
		}
	})
}
