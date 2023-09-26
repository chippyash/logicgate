// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"fmt"
	lg "github.com/chippyash/logicgate"
	"github.com/kelindar/bitmap"
)

// _adder is a helper function to perform the full adder
// there is no earthly reason for this to be a function, but it is. It is just to demonstrate
// that  the logic gates can be used to perform the full adder using bitmaps
func _adder(a, b, c bool) (bool, bool) {
	mA, mB, mC := bitmap.Bitmap{0}, bitmap.Bitmap{0}, bitmap.Bitmap{0}
	var sum, carry bool
	// whilst we are dealing with an entire  bitmap, we can use the logic gates to perform the necessary operations
	// one bit at a time
	if a {
		mA.Set(0)
	}
	if b {
		mB.Set(0)
	}
	if c {
		mC.Set(0)
	}
	//sum = a xor b xor carry
	sum = lg.Xor(lg.Xor(mA, mB), mC).Contains(0)
	//carry = (a and b) or (carry and (a xor b))
	carry = lg.Or(lg.And(mA, mB), lg.And(mC, lg.Xor(mA, mB))).Contains(0)

	return sum, carry
}

// Add demonstrates a full adder as described at https://en.wikipedia.org/wiki/Adder_(electronics)
func Add(a, b bitmap.Bitmap) bitmap.Bitmap {
	var sum bitmap.Bitmap
	//set new sum (s) and carry to false for initialisation
	var s, carry bool
	//make the sum is a bitmap the same length as the largest incoming bitmap
	l := len(a)
	if len(b) > l {
		l = len(b)
	}
	ll := uint32(64 ^ l)
	sum.Grow(ll)

	//this is a bit simplistic as it only takes into account the first 64 bits of the bitmaps
	fmt.Printf("       a: %016b b: %016b\n", a[0], b[0])
	//process for each bit in the bitmap. As this is operating on bitmaps as numbers, we assume MSB order
	//and therefore process from LSB to MSB
	for i := uint32(0); i < ll; i++ {
		//side effect notice. Whist s is the sum of the last addition, carry is affected by the result of the previous addition
		s, carry = _adder(a.Contains(i), b.Contains(i), carry)
		if s {
			sum.Set(i)
		}
		fmt.Printf("i: %02d mA: %01t mB: %01t carry: %01t sum: %016b\n", i, a.Contains(i), b.Contains(i), carry, sum[0])
	}
	return sum
}

func main() {
	a, b := 25, 9197
	s := Add(bitmap.Bitmap{uint64(a)}, bitmap.Bitmap{uint64(b)})
	fmt.Printf("%016b + %016b = %016b\n", a, b, s[0])
	fmt.Printf("Check %016b\n", a+b)
}
