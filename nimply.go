// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import "github.com/kelindar/bitmap"

// Nimply is the NIMPLY gate
// see https://en.wikipedia.org/wiki/NIMPLY_gate
// Use utils.Trim(Nimply(a,b), 1) to return the truthfulness of the expression
func Nimply(a, b bitmap.Bitmap) bitmap.Bitmap {
	return Not(Imply(a, b))
}
