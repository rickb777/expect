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

## Assertion Categories

There are **seven primary categories**, each introduce by a function:

### [Any](https://pkg.go.dev/github.com/rickb777/expect#Any)
This compares any types, but is especially structs, maps, arrays, slices. Although this will compare anything, it only provides equality tests and the error messages may be less informative than the other categories below

### [Bool](https://pkg.go.dev/github.com/rickb777/expect#Bool)
This compares `bool` and any subclass.

### [Error](https://pkg.go.dev/github.com/rickb777/expect#Error)
This compares `error` only.

### [Number](https://pkg.go.dev/github.com/rickb777/expect#Number)
This compares `int` and all the signed/unsigned int and float length variants, plus all their subtypes. This provides inequality comparisons. It also supports  `string` because that is also is an ordered type.

### [Map](https://pkg.go.dev/github.com/rickb777/expect#Map)
This compares `map[K]V` where the map key `K` is a comparable type.

### [Slice](https://pkg.go.dev/github.com/rickb777/expect#Slice)
This compares `[]T` where `T` is a comparable type.

### [String](https://pkg.go.dev/github.com/rickb777/expect#String)
This compares `string` and any subclass. It is more informative than `Any`.

These functions all take the actual value under test as their input. Other parameters can also be passed in; this allows the input to be a function with a multi-value return, for example. In this case, if any of the other parameters is non-nil (e.g. a non-nil `error`), the assertion will fail and give a corresponding error message. [Error](https://pkg.go.dev/github.com/rickb777/expect#Error) is subtly different - it considers the *last* non-nil argument as its actual input.

All of these include these methods

 * `ToBe(t, expected)` method that tests for equality (except for `Error`, which has `ToBeNil(t)` instead)
 * `Not()` method that inverts the assertion
 * `Info(...)` method that provides information in any failure message arising. There is a terse synonym `I(...)` too.

Most of them have

 * a `ToEqual(t, expected)` method that also tests for equality *ignoring* whether the concrete types match or not (`Error`, `Number` and `Slice` don't have this though)

All of the assertion methods `ToXxxx` listed above and below include a `t Tester` (see [Tester](https://pkg.go.dev/github.com/rickb777/expect#Tester)) parameter; normally this will be `*testing.T` but you can use your own type if you need to embed this API in other assertion logic.

There are various other methods too

 * [Bool](https://pkg.go.dev/github.com/rickb777/expect#Bool) has `ToBeTrue(t)` and `ToBeFalse(t)`
 * [Error](https://pkg.go.dev/github.com/rickb777/expect#Error) has `ToBeNil(t)` and `ToHaveOccurred(t)`
 * [Number](https://pkg.go.dev/github.com/rickb777/expect#Number) has `ToBeGreaterThan[OrEqualTo](t, threshold)` and `ToBeLessThan[OrEqualTo](t, threshold)`
 * [Map](https://pkg.go.dev/github.com/rickb777/expect#Map) has `ToContain(t, key, [value])`
 * [Slice](https://pkg.go.dev/github.com/rickb777/expect#Slice) has `ToContain{All|Any}(t, substring)`
 * [String](https://pkg.go.dev/github.com/rickb777/expect#String) has `ToContain(t, substring)`

`Any`, `Map`, and `Slice` use [cmp.Equal](https://pkg.go.dev/github.com/google/go-cmp/cmp) under the hood. This is flexible, allowing for options to control how the comparison proceeds - for example when considering how close floating point numbers need to be to be considered equal. There is a `Using(...)` method to specify what options it should use.

## Status

This is not yet production-ready.

## History

This API was mostly inspired by [Gomega](https://github.com/onsi/gomega), which had some great ideas but is overly complex to use.
