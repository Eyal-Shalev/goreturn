# goreturn

[![Tests](https://github.com/Eyal-Shalev/goreturn/actions/workflows/test.yml/badge.svg)](https://github.com/Eyal-Shalev/goreturn/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Eyal-Shalev/goreturn)](https://goreportcard.com/report/github.com/Eyal-Shalev/goreturn)
[![GoDoc](https://godoc.org/github.com/Eyal-Shalev/goreturn?status.svg)](https://godoc.org/github.com/Eyal-Shalev/goreturn)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`goreturn` is a Go library that provides a set of functions to run a function in a goroutine and return its result(s) in a channel. This library is designed to make it easier to work with goroutines and channels in Go, especially when dealing with functions that return multiple values.

## Features

- `Return0(fn func()) <-chan any`: Runs a function with no return values in a goroutine and returns a channel that will receive a single `nil` value when the function completes.
- `Return1[T1 any](fn func() T1) <-chan T1`: Runs a function with one return value in a goroutine and returns a channel that will receive the return value.
- `Return2[T1 any, T2 any](fn func() (T1, T2)) <-chan *Tuple[T1, T2]`: Runs a function with two return values in a goroutine and returns a channel that will receive a tuple containing the return values.
- `Return3[T1 any, T2 any, T3 any](fn func() (T1, T2, T3)) <-chan *Tuple3[T1, T2, T3]`: Runs a function with three return values in a goroutine and returns a channel that will receive a tuple containing the return values.

## Installation

To install `goreturn`, use the `go get` command:

```shell
go get github.com/Eyal-Shalev/goreturn
```

## Usage

Here is an example of how to use `goreturn`:

```go
package main

import (
	"fmt"
	"github.com/Eyal-Shalev/goreturn"
)

func main() {
	c := goreturn.Return1(func() int {
		return 42
	})
	v := <-c
	fmt.Println("Return1() returned", v)
}
```

In this example, the function `func() int { return 42 }` is run in a goroutine, and its return value is received from the channel returned by `Return1`.

## Contributing

Contributions to `goreturn` are welcome! Please submit a pull request on GitHub.

## License

`goreturn` is licensed under the MIT License. See the `LICENSE` file for details.