// Copyright (c) Ashley Kitson, 2023. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"fmt"
	"github.com/chippyash/logicgate"
	"github.com/chippyash/logicgate/utils"
	"github.com/kelindar/bitmap"
)

var a, b, c bitmap.Bitmap

func InitFE() {
	a = bitmap.Bitmap{0x00}
	b = bitmap.Bitmap{0x01}
	fmt.Printf("\na=%08b, b=%08b\n\n", a[0], b[0])
}

func ExAnd() {
	println("AND")
	println("a AND a == a")
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], a[0], logicgate.And(a, a)[0])
	println("a AND b == a")
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], b[0], logicgate.And(a, b)[0])
	println("b AND a == a")
	fmt.Printf("%08b ^ %08b = %08b\n", b[0], a[0], logicgate.And(b, a)[0])
	println("b AND b == b")
	fmt.Printf("%08b ^ %08b = %08b\n\n\n", b[0], b[0], logicgate.And(b, b)[0])
}

func ExOr() {
	println("OR")
	println("a OR a == a")
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], a[0], logicgate.Or(a, a)[0])
	println("a OR b == b")
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], b[0], logicgate.Or(a, b)[0])
	println("b OR a == b")
	fmt.Printf("%08b ^ %08b = %08b\n", b[0], a[0], logicgate.Or(b, a)[0])
	println("b OR b == b")
	fmt.Printf("%08b ^ %08b = %08b\n\n\n", b[0], b[0], logicgate.Or(b, b)[0])
}

func ExXor() {
	println("XOR")
	println("a XOR a == a")
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], a[0], logicgate.Xor(a, a)[0])
	println("a XOR b == b")
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], b[0], logicgate.Xor(a, b)[0])
	println("b XOR a == b")
	fmt.Printf("%08b ^ %08b = %08b\n", b[0], a[0], logicgate.Xor(b, a)[0])
	println("b XOR b == a")
	fmt.Printf("%08b ^ %08b = %08b\n\n\n", b[0], b[0], logicgate.Xor(b, b)[0])
}

func ExImply() {
	println("IMPLY")
	println("a --> a == true")
	var res bitmap.Bitmap
	res = utils.Trim(logicgate.Imply(a, a), 1)
	fmt.Printf("%01b ^ %01b = %01b\n", a[0], a[0], res[0])
	println("a --> b == true")
	res = utils.Trim(logicgate.Imply(a, b), 1)
	fmt.Printf("%01b ^ %01b = %01b\n", a[0], b[0], res[0])
	println("b --> a == false")
	res = utils.Trim(logicgate.Imply(b, a), 1)
	fmt.Printf("%01b ^ %01b = %01b\n", b[0], a[0], res[0])
	println("b --> b == true")
	res = utils.Trim(logicgate.Imply(b, b), 1)
	fmt.Printf("%01b ^ %01b = %01b\n\n\n", b[0], b[0], res[0])
}

func ExNimply() {
	println("NIMPLY")
	println("a -/-> a == false")
	var res bitmap.Bitmap
	res = utils.Trim(logicgate.Nimply(a, a), 1)
	fmt.Printf("%01b ^ %01b = %01b\n", a[0], a[0], res[0])
	println("a -/-> b == false")
	res = utils.Trim(logicgate.Nimply(a, b), 1)
	fmt.Printf("%01b ^ %01b = %01b\n", a[0], b[0], res[0])
	println("b -/-> a == true")
	res = utils.Trim(logicgate.Nimply(b, a), 1)
	fmt.Printf("%01b ^ %01b = %01b\n", b[0], a[0], res[0])
	println("b -/-> b == false")
	res = utils.Trim(logicgate.Nimply(b, b), 1)
	fmt.Printf("%01b ^ %01b = %01b\n\n\n", b[0], b[0], res[0])
}

func ExNand() {
	println("NAND")
	var res bitmap.Bitmap
	println("a NAND a")
	res = utils.Trim(logicgate.Nand(a, a), 8)
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], a[0], res[0])
	println("a NAND b")
	res = utils.Trim(logicgate.Nand(a, b), 8)
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], b[0], res[0])
	println("b NAND a")
	res = utils.Trim(logicgate.Nand(b, a), 8)
	fmt.Printf("%08b ^ %08b = %08b\n", b[0], a[0], res[0])
	println("b NAND b")
	res = utils.Trim(logicgate.Nand(b, b), 8)
	fmt.Printf("%08b ^ %08b = %08b\n\n\n", b[0], b[0], res[0])
}

func ExXnor() {
	println("XNOR")
	var res bitmap.Bitmap
	println("a XNOR a")
	res = utils.Trim(logicgate.Xnor(a, a), 8)
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], a[0], res[0])
	println("a XNOR b")
	res = utils.Trim(logicgate.Xnor(a, b), 8)
	fmt.Printf("%08b ^ %08b = %08b\n", a[0], b[0], res[0])
	println("b XNOR a")
	res = utils.Trim(logicgate.Xnor(b, a), 8)
	fmt.Printf("%08b ^ %08b = %08b\n", b[0], a[0], res[0])
	println("b XNOR b")
	res = utils.Trim(logicgate.Xnor(b, b), 8)
	fmt.Printf("%08b ^ %08b = %08b\n\n\n", b[0], b[0], res[0])
}

func ExNot() {
	println("NOT")
	println("NOT a")
	fmt.Printf("NOT %08b => %08b\n", a[0], logicgate.Not(a)[0])
	println("NOT b")
	fmt.Printf("NOT %08b => %08b\n\n\n", b[0], logicgate.Not(b)[0])
}

func ExTrim() {
	println("Trim(bitmap, bitLen) => bitmap")
	bm := bitmap.Bitmap{0xF0F0F0F0F0F0F0F0}
	fmt.Printf("bitmap            = %64b\n", bm[0])
	b2 := utils.Trim(bm, 64)
	fmt.Printf("Trim(bitmap, 64) => %64b\n", b2[0])
	b2 = utils.Trim(bm, 32)
	fmt.Printf("Trim(bitmap, 32) => %64b\n", b2[0])
	b2 = utils.Trim(bm, 16)
	fmt.Printf("Trim(bitmap, 16) => %64b\n", b2[0])
	b2 = utils.Trim(bm, 8)
	fmt.Printf("Trim(bitmap,  8) => %64b\n", b2[0])
}

func ExRangeall() {
	println("RangeAll(bitmap, func(k uint32, v bool))")
	bm := bitmap.Bitmap{0xFFFFFFFFFFFFFFFF}
	rFunc := func(k uint32, v bool) {
		if k%2 == 0 {
			bm.Remove(k)
		}
	}
	fmt.Printf("bitmap           = %64b\n", bm[0])
	fmt.Println("rFunc := func(k uint32, v bool) {\n\t\tif k%2 == 0 {\n\t\t\tbm.Remove(k)\n\t\t}\n\t}")
	utils.RangeAll(bm, rFunc)
	fmt.Printf("RangeAll(bitmap, rFunc) => %64b\n\n\n", bm[0])
}

func ExCompare() {
	println("Compare(bitmap, bitmap) => int, error")
	a = bitmap.Bitmap{0x00}
	b = bitmap.Bitmap{0xff}
	fmt.Printf("a=%08b, b=%08b\n", a[0], b[0])
	r, _ := utils.Compare(a, a)
	fmt.Printf("Compare(%08b, %08b) => %d\n", a[0], a[0], r)
	r, _ = utils.Compare(a, b)
	fmt.Printf("Compare(%08b, %08b) => %d\n", a[0], b[0], r)
	r, _ = utils.Compare(b, a)
	fmt.Printf("Compare(%08b, %08b) => %d\n", b[0], a[0], r)
	r, _ = utils.Compare(b, b)
	fmt.Printf("Compare(%08b, %08b) => %d\n", b[0], b[0], r)
}

func main() {
	InitFE()
	println("Functional Equivalence\n")
	ExAnd()
	ExOr()
	ExXor()
	println("Additional Functions\n")
	ExImply()
	ExNimply()
	ExNand()
	ExNot()
	ExXnor()
	println("Utility functions\n")
	ExTrim()
	ExRangeall()
	ExCompare()
}
