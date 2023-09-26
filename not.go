// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import "github.com/kelindar/bitmap"

// Not is the Inverter Gate
// see https://en.wikipedia.org/wiki/Inverter_(logic_gate)
func Not(a bitmap.Bitmap) bitmap.Bitmap {
	var res bitmap.Bitmap
	a.Clone(&res)
	res.Ones()

	//clears bits in res if they are set in a
	a.Range(func(k uint32) {
		res.Remove(k)
	})

	return res
}
