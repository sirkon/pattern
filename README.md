# pattern
[![Build Status](https://travis-ci.org/sirkon/pattern.svg?branch=master)](https://travis-ci.org/sirkon/pattern)

This library is aimed to deal with simple patterns, i.e. patterns of certain length with some positions being wildcards

## Samples of patterns

| Pattern text      | Match example     | Will not match example |
|-------------------|-------------------|------------------------|
| a..a              | abba              | abbba                  |
| ..:..:..:..:..:.. | 00:11:22:33:44:55 | 0:11:22:33:44:5        |
| ..\\...           | 11.22             | 12345                  |

## Sample rules

* As you probably noticed, `.` plays the role of `.` in regular expressions, i.e. wildcard character
* There are escape sequences, those are available
    ```
    \n
    \r
    \t
    \\
    \`
    \.
    ``` 

## Code usage sample

```go
package main

import (
	"log"
	
	"github.com/sirkon/pattern"
)

func main() {
	p, err := pattern.NewPattern("a..a")
	if err != nil {
		log.Fatal(err)
	}
	
	if !p.Match([]byte("abba")) {
		log.Fatalf("pattern %s must match against abba", p)
	}
	log.Println(string(p.Lookup([]byte("  abba-cda")))) // will output "abba-cda"
}
```

## Benchmarking (Core i5-4670k)

| Test name                    | Iterations | Operation cost in ns      |
|------------------------------|------------|---------------------------|
| BenchmarkPattern_Match-4     | 10000000   | 127 ns/op                 |
| BenchmarkRegexp_Match-4      | 500000     | 3235 ns/op                |
| BenchmarkRagel_Match-4       | 20000000   | 112 ns/op                 |
| BenchmarkPattern_Lookup-4    | 30000      | 43941 ns/op               |
| BenchmarkRegexp_Lookup-4     | 300        | 4562271 ns/op             |
| BenchmarkRagel_Lookup-4      | 10000      | 121227 ns/op              |

As you see, regular expression are much-much slower (there's a room for optimization I believe though). And pattern 
vs ragel comparsion is bit tricky: ragel is a tad faster at regular match but loses at lookup. Lookup
comparison was done with this sample set.  

| Prefix length      | Text                   |
|--------------------|------------------------|
| 0                  | `00:00:00:00:00:00`    |
| 1                  | `_00:00:00:00:00:00`   |
| 2                  | `__00:00:00:00:00:00`  |
| …                  | …                      |
| 128                | `_…_00:00:00:00:00:00` | 


But Ragel is as fast as current implementation when limited to 16 character in the head and faster with shorter 
prefixes. Obviously, other patterns and source strings may lead to different results. 