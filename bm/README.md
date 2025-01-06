# A simple concurrent-safe bitmap implementation

## Usage
```go
package main

import (
	"github.com/althk/collections/bm"
)

func main() {
	bmap := bm.New(16)  // initialize a bitmap of 16 bits size
	
	// set a few bits
	_ = bmap.Set(2)
	_ = bmap.Set(3)
	
	// it is an error when trying to set a -ve bit or one larger
	// than the size.
	err := bmap.Set(100)
	// errors.Is(err, bm.ErrOutOfBounds)
	
	// check if a bit is set
	if bmap.IsSet(2) {
		// do something when bit 2 is set
    }
	
	// Clear the bit after work
	_ = bmap.Clear(2)
}
```
### Benchmark

```shell
goos: linux
goarch: amd64
pkg: github.com/althk/bm
cpu: AMD Ryzen 7 7700X 8-Core Processor             
BenchmarkBitmap_Set/size_10000-16         	283787299	         4.216 ns/op
BenchmarkBitmap_Set/size_100000-16        	284957320	         4.238 ns/op
BenchmarkBitmap_Set/size_1000000-16       	284969374	         4.201 ns/op
PASS
ok  	github.com/althk/bm	4.884s
```