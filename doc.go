// Package goreturn provides a set of functions that run a given function in a goroutine
// and return its result(s) in a channel. This makes it easier to work with goroutines
// and channels in Go, especially when dealing with functions that return multiple values.
//
// The Return0 function takes a function with no return values as an argument,
// runs it in a goroutine, and returns a channel that will receive a single `nil` value
// when the function completes. This is useful for running a function asynchronously
// and being notified when it has finished executing.
//
// The Return1 function takes a function with one return value as an argument,
// runs it in a goroutine, and returns a channel that will receive the return value.
// This allows you to run a function asynchronously and receive its result in a non-blocking manner.
//
// The Return2 function takes a function with two return values as an argument,
// runs it in a goroutine, and returns a channel that will receive a tuple containing the return values.
// This is useful when you want to run a function that returns two values asynchronously,
// and receive both results in a non-blocking manner.
//
// The Return3 function takes a function with three return values as an argument,
// runs it in a goroutine, and returns a channel that will receive a tuple containing the return values.
// This allows you to run a function that returns three values asynchronously,
// and receive all three results in a non-blocking manner.
//
// The Tuple and Tuple3 types are used to hold the return values from the functions passed to Return2 and Return3 respectively.
// They are used to send multiple return values through a single channel.
package goreturn
