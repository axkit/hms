# hms package

[![Build Status](https://github.com/axkit/bitset/actions/workflows/go.yml/badge.svg)](https://github.com/axkit/hms/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axkit/hms)](https://goreportcard.com/report/github.com/axkit/hms)
[![GoDoc](https://pkg.go.dev/badge/github.com/axkit/hms)](https://pkg.go.dev/github.com/axkit/hms)
[![Coverage Status](https://coveralls.io/repos/github/axkit/hms/badge.svg?branch=master)](https://coveralls.io/github/axkit/hms?branch=master)

The `hms` package provides a simple time handling within a day.  

## Installation

To install the package, run:

```sh
go get github.com/axkit/hms
```

## Usage

### Creating a new HMS instance

```go
package main

import (
	"fmt"
	"time"
	"github.com/axkit/hms"
)

func main() {
	h := hms.New(60*60 * time.Second)
	fmt.Println(h) // Output: 01:00:00
}
```

### Adding durations

```go
h := hms.New(23 * time.Hour)
newH, overMidnight := h.Add(time.Hour + time.Minute)
fmt.Println(newH) // Output: 00:01:00
fmt.Println(overMidnight) // Output: true
```

### Parsing from string

```go
h, err := hms.Parse("01:02:03")
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println(h) // Output: 01:02:03
}
```

### Converting to duration

```go
h := hms.New(3661 * time.Second)
d := h.ToDuration()
fmt.Println(d) // Output: 1h1m1s
```

## Errors

`ErrParseFailed`

Returned by parsing functions when an invalid format/character is encountered.

## Running Tests

To run the tests, use:

```sh
go test ./...
```

## License
This package is open-source and distributed under the MIT License. Contributions and feedback are welcome!

