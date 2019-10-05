# Knuth-Morris-Pratt-Algorithm
Implementation of the Knuth-Morris-Pratt algorithm or KMP algorithm in GO

## Technology
* [go](https://golang.org/)
* [ginkgo](https://github.com/onsi/ginkgo)

## Usage
```go
import "github.com/fbonhomm/Knuth-Morris-Pratt-Algorithm/kmp"

var index, length int
var buffer = []byte("abc abcdab abcdabcdabde")
var pattern = []byte("abcdabd") // 7
var pattern1 = []byte("abcdabdr") // 8

index = kmp.SearchFirstOccurrence(buffer, pattern) // 15
index = kmp.SearchFirstOccurrence(buffer, pattern1) // -1

index, length = kmp.SearchGreaterOccurrence(buffer, pattern) // 15, 7
index, length = kmp.SearchGreaterOccurrence(buffer, pattern1) // 15, 7
```

## LZSS Explications
 - [Explications en francais](doc/FR-EXPLICATION.md)
 - [English Explications](doc/EN-EXPLICATION.md) (coming soon)

## Links
https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm