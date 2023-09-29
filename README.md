<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/chippyash/logicgate" alt="Go Version">
<a href="https://pkg.go.dev/github.com/chippyash/logicgate"><img src="https://pkg.go.dev/badge/github.com/chippyash/logicgate" alt="PkgGoDev"></a>
<a href="https://goreportcard.com/report/github.com/chippyash/logicgate"><img src="https://goreportcard.com/badge/github.com/chippyash/logicgate" alt="Go Report Card"></a>
<a href="https://coveralls.io/github/chippyash/logicgate?branch=master"><img src="https://coveralls.io/repos/github/chippyash/logicgate/badge.svg?branch=master" alt="Coverage Status"></a>
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
</p>

# Logic Gates
## github.com/chippyash/logicgate

## What
Builds on [Kelindar's Bitmap library](github.com/kelindar/bitmap) to provide:
 - functional equivalents to the bitmap declarative methods
 - additional functions for some 'missing' methods
 - some utility functions that wouldn't make it into the Bitmap library

## Why
To provide the ability to write code functionally using bitmaps. Parameters
to the functions are not changed, maintaining immutability.

To show that it is reasonably possible to take someone else's code and provide a different experience without harming
the originator.

Kelindar's Bitmap library allows for very large bitmaps. His primary use case is categorisation, for which bit maps
are a good candidate, particularly if you are trying to slim down your data storage. My use case was slightly different,
but Kelindar's Bitmap serves the purpose as a highly efficient implementation.

The `bit` that was missing was the ability to compare bitmaps.

## How

 - Go 1.18+  The base version is dictated by the Bitmap library. This may change in the future.
 - `import "github.com/chippyash/logicgate"`
 - Read the unit tests to see how functions are used
 - Run the examples/main.go program for example usage

### Bitmap functional equivalents

 - func And(a, b bitmap.Bitmap) bitmap.Bitmap
 - func Or(a, b bitmap.Bitmap) bitmap.Bitmap
 - func Xor(a, b bitmap.Bitmap) bitmap.Bitmap

### Additional functions

 - func Imply(a, b bitmap.Bitmap) bitmap.Bitmap
 - func Nimply(a, b bitmap.Bitmap) bitmap.Bitmap
 - func Nand(a, b bitmap.Bitmap) bitmap.Bitmap
 - func Not(a bitmap.Bitmap) bitmap.Bitmap
 - func Xnor(a, b bitmap.Bitmap) bitmap.Bitmap

### Utility functions

 - func Compare(a, b bitmap.Bitmap) int
 - func RangeAll(src bitmap.Bitmap, fn func(k uint32, v bool))
 - func Trim(a bitmap.Bitmap, bitLen uint32) bitmap.Bitmap

### An example of what you can do
My first use for Bitmap was to categorise team games by the players in the team, the objective being to ensure that I 
hadn't gotten the same players playing together too often.  That's where I hit a roadblock with Bitmap, it had no comparator.
In a sense, this library is a response to that. But once down the rabbit hole, off we must go!

`examples/full_adder/main.go` is a purely contrived attempt to produce a bitmap oriented electronic circuit as described in
https://en.wikipedia.org/wiki/Adder_(electronics)  It has no functional reason to exist, as there are far better ways to do
these things. It exists only as a demonstration of the art of the possible.


## Contributing

Feel free to submit a pull request to make changes or add functionality.  Please ensure that
changes are accompanied by relevant unit tests. I am particularly interested in efficiency improvements
if you can make them, or edge cases if you find them, and of course `the art of the possible` as examples that others may 
learn from.

I cannot fix issues in the dependency chain. Please refer to the authors of a dependency.

### Developing

 - clone the repo
 - `go get -u ./...`

## References
[Logic Gates](https://en.wikipedia.org/wiki/Logic_gate)

## License
The library is licensed under the [MIT License](LICENSE)