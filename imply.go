// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package logicgate

import "github.com/kelindar/bitmap"

// Imply is the IMPLY gate
// see https://en.wikipedia.org/wiki/IMPLY_gate
// Use utils.Trim(Imply(a,b), 1) to return the truthfulness of the expression
func Imply(a, b bitmap.Bitmap) bitmap.Bitmap {
	return Or(Not(a), b)
}
