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

func TestNor_Zero_Zero_is_One(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)

	res := logicgate.Nor(a, b)
	assert.True(t, res.Contains(1))
}

func TestNor_Zero_One_is_Zero(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)
	b.Ones()

	res := logicgate.Nor(a, b)
	assert.False(t, res.Contains(1))
}

func TestNor_One_Zero_is_Zero(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)
	a.Ones()

	res := logicgate.Nor(a, b)
	assert.False(t, res.Contains(1))
}

func TestNor_One_One_is_Zero(t *testing.T) {
	a, b := bitmap.Bitmap{}, bitmap.Bitmap{}
	a.Grow(64)
	b.Grow(64)
	a.Ones()
	b.Ones()

	res := logicgate.Nor(a, b)
	assert.False(t, res.Contains(1))
}
