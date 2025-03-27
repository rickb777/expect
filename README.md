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

There are **eight primary categories**, each introduce by a function:

### expect.[Any](https://pkg.go.dev/github.com/rickb777/expect#Any)(actual ...)
This compares any types, but is especially useful for structs, maps, arrays, slices. Although this will compare anything, it only provides equality tests and the error messages may be less informative than the other categories below.

### expect.[String](https://pkg.go.dev/github.com/rickb777/expect#String)(actual ...)
This compares `string` and any subclass. It is more informative than **Any**, highlighting where the differences start.

### expect.[Number](https://pkg.go.dev/github.com/rickb777/expect#Number)(actual ...)
This compares `int` and all the signed/unsigned int and float length variants, plus all their subtypes. This provides inequality comparisons. It also supports  `string` because that is also is an ordered type.

### expect.[Bool](https://pkg.go.dev/github.com/rickb777/expect#Bool)(actual ...)
This compares `bool` and any subclass.

### expect.[Map](https://pkg.go.dev/github.com/rickb777/expect#Map)(actual ...)
This compares `map[K]V` where the map key `K` is a comparable type.

### expect.[Slice](https://pkg.go.dev/github.com/rickb777/expect#Slice)(actual ...)
This compares `[]T` but only where `T` is a comparable type. Use **Any** for other slices.

### expect.[Error](https://pkg.go.dev/github.com/rickb777/expect#Error)(... actual)
This compares `error` only.

### expect.[Func](https://pkg.go.dev/github.com/rickb777/expect#Func)(func)
This runs some function and checks whether it panicked.

## Application

The eight primary functions above all take the actual value under test as their input.

Other parameters can also be passed in. If any of these other parameters is non-nil (e.g. a non-nil `error`), the assertion will fail and give a corresponding error message. This allows, for example, the input to be a function with a multi-value return. 

Note that **Error** is different - it considers the /last/ non-nil argument as its actual input. Any preceding arguments are ignored.

## Basic Methods

All categories include these general methods

 * `Info(...)` provides information in the failure message, if there is one. There is a terse synonym `I(...)` too.
 * `Not()` inverts the assertion defined by the `ToXxxx` method that follows it (these assertions are described below)

**String** also has `Trim(n)` that truncates message strings if they exceed the specified length.

## Assertions

The assertions are all infinitive verbs, i.e. methods such as `ToBe`. All of them require a `t Tester` parameter (see [Tester](https://pkg.go.dev/github.com/rickb777/expect#Tester)). Normally this will be `*testing.T` but you can use your own type if you need to embed this API in other assertion logic.

Many categories have

 * `ToBe(t, expected)` **tests for equality**.
 * `ToEqual(t, expected)` tests for equality ignoring whether the concrete types match or not
 * `ToBeNil(t)` tests for nil value.

| Category | `ToBe` | `ToEqual` | `ToBeNil` |
|----------|--------|-----------|-----------|
| Any      | Yes    | -         | Yes       |
| String   | Yes    | Yes       | -         |
| Number   | Yes    | -         | -         |
| Bool     | Yes    | Yes       | -         |
| Map      | Yes    | -         | Yes       |
| Slice    | Yes    | -         | Yes       |
| Error    | -      | -         | Yes       |
| Func     | -      | -         | -         |

### Size / Length

**String**, **Slice** and **Map** have a `ToHaveLength(t, expected)` method, plus `ToBeEmpty(t)`.

`ToHaveSize(t, expected)` is a synonym for `ToHaveLength(t, expected)`.

### Others

There are various other methods too

* **String** has `ToContain(t, substring)` 
* **Slice** has `ToContainAll(t, ...)`, `ToContainAny(t, ...)`
* **Map** has `ToContain(t, key, [value])` (the value, if present, must match what is held in the map)
* **Number** has four `ToBeGreaterThan[OrEqualTo](t, threshold)` and `ToBeLessThan[OrEqualTo](t, threshold)` methods
* **Bool** has `ToBeTrue(t)` and `ToBeFalse(t)`
 * **Error** has `ToBeNil(t)` and `ToHaveOccurred(t)`
 * **Func** has `ToPanic(t)` and `ToPanicWithMessage(t, string)`

## Compare Options

**Any**, **Map**, and **Slice** use [cmp.Equal](https://pkg.go.dev/github.com/google/go-cmp/cmp) under the hood. This is flexible, allowing for options to control how the comparison proceeds - for example when considering how close floating point numbers need to be to be considered equal. There is a `Using(...)` method to specify what options it should use. By default, the three options used are

 * All fields in structs are compared, regardless of whether they exported or unexported; all structs in maps and slices are treated likewise. 
 * Floating point numbers are compared within the tolerance set by `ApproximateFloatFraction`.
 * Maps/slices that are empty are treated the same as those that are nil.

## Status

This is now ready for beta testing.

## History

This API was mostly inspired by [Gomega](https://github.com/onsi/gomega), which had some great ideas but is overly complex to use.
