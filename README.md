[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Knuth-Morris-Pratt Algorithm
<strong>[EN]</strong>  
Implementation of the Knuth-Morris-Pratt algorithm or KMP algorithm in GO.

Knuth-Morris-Pratt is a substring search algorithm.

<strong>[FR]</strong>  
Implémentation de l'algorithme de Knuth-Morris-Pratt ou l'algorithme KMP en Go.

Knuth-Morris-Pratt est un algorithme de recherche de sous-chaîne.

## Technology
* [go](https://golang.org/)
* [ginkgo](https://github.com/onsi/ginkgo)

## Usage
CLI:
```bash
go test -v test/
```

CODE:
```go
import kmp "github.com/fbonhomm/knuth-morris-pratt/source"

var index, length int
var buffer = []byte("abc abcdab abcdabcdabde")
var pattern = []byte("abcdabd") // 7
var pattern1 = []byte("abcdabdr") // 8

index = kmp.Search(buffer, pattern) // 15
index = kmp.Search(buffer, pattern1) // -1
```

## Explanations
 - [English](documentation/explanation.en.md)
 - [Français](documentation/explanation.fr.md)

## Links
https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm 
https://dev.to/girish3/string-matching-kmp-algorithm-cie  
https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/ 
