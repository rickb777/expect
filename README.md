# expect

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rickb777/expect)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickb777/expect)](https://goreportcard.com/report/github.com/rickb777/expect)
[![Issues](https://img.shields.io/github/issues/rickb777/expect.svg)](https://github.com/rickb777/expect/issues)

## Simple easy-to-use assertions to use in Go tests.

 * Fluent API
 * Clear error messages
 * Works with Go `testing` API
 * Also works independently
 * Type safety thanks to Go generics
 * No dependencies other than `github.com/google/go-cmp`

## Five Assertion Categories

There are five primary categories:

 * *Strings* - `string` and any subclass
 * *Numbers* - all ordered types such as `int`, `float32`, plus all the signed/unsigned int and float length variants, plus all their subtypes, plus `string` and any subclass
 * *Bools* - `bool` and any subclass
 * *Errors* - `error` only
 * *Everything else* - structs, maps, arrays, slices as handled by [cmp.Equal](https://pkg.go.dev/github.com/google/go-cmp/cmp)

The five kinds of assertion are introduced by the five functions
[Any](https://pkg.go.dev/github.com/rickb777/expect#Any),
[Bool](https://pkg.go.dev/github.com/rickb777/expect#Bool),
[Error](https://pkg.go.dev/github.com/rickb777/expect#Error),
[Number](https://pkg.go.dev/github.com/rickb777/expect#Number), and
[String](https://pkg.go.dev/github.com/rickb777/expect#String).
These functions all take the actual value under test as their input. Other parameters can also be passed in; this allows the input to be a function with a multi-value return, for example. In this case, if any of the other parameters is non-nil (e.g. a non-nil `error`), the assertion will fail and give a corresponding error message. [Error](https://pkg.go.dev/github.com/rickb777/expect#Error) is subtly different - it considers the *last* non-nil argument as its actual input.

All five of these have 

 * a `ToBe(expected, t)` method that tests for equality (except for `Error`, which has `ToBeNil(t)` instead)
 * a `ToEqual(expected,t )` method that also tests for equality, ignoring whether the concrete types match or not (`Error` doesn't have this though)
 * a `Not()` method that inverts the assertion
 * an `Info(...)` method that provides information in any failure message arising. There is a terse synonym `I(...)` too.

All of the assertion methods listed above and below include a `t Tester` (see [Tester](https://pkg.go.dev/github.com/rickb777/expect#Tester)) parameter; normally this will be `*testing.T` but you can use your own type if you need to embed this API in other assertion logic.

There are various other methods too

 * [Bool](https://pkg.go.dev/github.com/rickb777/expect#Bool) has `ToBeTrue(t)` and `ToBeFalse(t)`
 * [Error](https://pkg.go.dev/github.com/rickb777/expect#Error) has `ToBeNil(t)` and `ToHaveOccurred(t)`
 * [Number](https://pkg.go.dev/github.com/rickb777/expect#Number) has `ToBeGreaterThan[OrEqualTo](threshold, t)` and `ToBeLessThan[OrEqualTo](threshold, t)`
 * [String](https://pkg.go.dev/github.com/rickb777/expect#String) has `ToContain(substring, t)`

