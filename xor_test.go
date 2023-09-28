//go:build unit
// +build unit

// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate_test

import (
	"github.com/chippyash/logicgate"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXor_Zero_Zero_Is_Zero(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)
	res := logicgate.Xor(a, b)
	assert.False(t, res.Contains(1))
}

func TestXor_One_One_Is_Zero(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)
	a.Ones()
	b.Ones()
	res := logicgate.Xor(a, b)
	assert.False(t, res.Contains(1))
}

func TestXor_One_Zero_Is_One(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)
	a.Ones()
	res := logicgate.Xor(a, b)
	assert.True(t, res.Contains(1))
}

func TestXor_Zero_One_Is_One(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)
	b.Ones()
	res := logicgate.Xor(a, b)
	assert.True(t, res.Contains(1))
}
