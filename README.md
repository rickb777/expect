# expect

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rickb777/expect)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickb777/expect)](https://goreportcard.com/report/github.com/rickb777/expect)
[![Build](https://github.com/rickb777/expect/actions/workflows/go.yml/badge.svg)](https://github.com/rickb777/expect/actions)
[![Coverage](https://coveralls.io/repos/github/rickb777/expect/badge.svg?branch=main)](https://coveralls.io/github/rickb777/expect?branch=main)
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

### expect.[Any](https://pkg.go.dev/github.com/rickb777/expect#Any)(actual ...) | expect.[Value](https://pkg.go.dev/github.com/rickb777/expect#Value)(actual ...)

This compares equality for values of any type, but is especially useful for structs, maps, arrays, and slices. It only provides equality tests; the other seven categories below provide a much wider
range.

* If the value under test is of a type with a method `a.Equal(b)` (`a` and `b` having the same type), then the `Equal` method will be used. So this compares types such as `time.Time` correctly.

* Otherwise it behaves like `reflect.DeepEqual`.

The **Any** function is an alias of **Value**; use whichever you prefer.

### expect.[String](https://pkg.go.dev/github.com/rickb777/expect#String)(actual ...)

This compares `string` and any subclass. It is more informative than **Any**, highlighting where the differences start.

### expect.[Number](https://pkg.go.dev/github.com/rickb777/expect#Number)(actual ...)

This compares `int` and all the signed/unsigned int and float length variants, plus all their subtypes. This provides inequality comparisons. It also supports  `string` because that is also is an
ordered type.

However, for near-equality testing of `float32` or `float64`, use **Any** instead because the tolerance [can be specified](#readme-options-for-controlling-how-the-comparisons-work).

### expect.[Bool](https://pkg.go.dev/github.com/rickb777/expect#Bool)(actual ...)

This compares `bool` and any subclass.

### expect.[Map](https://pkg.go.dev/github.com/rickb777/expect#Map)(actual ...)

This compares `map[K]V` where the map key `K` is a comparable type.

**Map** provides more methods than **Any**, but is otherwise very similar.

### expect.[Slice](https://pkg.go.dev/github.com/rickb777/expect#Slice)(actual ...)

This compares `[]T` where `T` is any type.

**Slice** provides more methods than **Any**, but is otherwise very similar.

### expect.[Error](https://pkg.go.dev/github.com/rickb777/expect#Error)(... actual)

This compares `error` only.

### expect.[Func](https://pkg.go.dev/github.com/rickb777/expect#Func)(func)

This runs some function and checks whether it panicked.

## Application

The eight primary functions above all take the **actual value** under test as their input.

Other parameters can also be passed in. If any of these other parameters is a non-nil `error`, the assertion will fail and give a corresponding error message. Any other parameters are ignored; this
includes any nil `error`.

**Error** is slightly different - it considers the *last* non-nil `error` as its actual input. Any other parameters are ignored; this includes any nil `error`.

In particular, this allows the input to be a function with a multi-value return.

## Assertions

The assertions are all infinitive verbs, i.e. methods such as `ToBe`. Typical use is of the form

```go
    expect.Value(myValue).ToBe(t, expectedValue)
expect.String(myValue).ToBe(t, expectedValue)
expect.Number(myValue).ToBe(t, expectedValue)
expect.Bool(myValue).ToBe(t, expectedValue)
expect.Map(myValue).ToBe(t, expectedValue)
expect.Slice(myValue).ToBe(t, expectedValue)
expect.Error(myValue).ToBeNil(t)
expect.Func(myFunc).ToPanic(t)
```

All of the assertions require a `t Tester` parameter (see [Tester](https://pkg.go.dev/github.com/rickb777/expect#Tester)). Normally this will be `*testing.T` but you can use your own type if you need
to embed this API in other assertion logic.

The assertions available are intended to be obvious to use because they relate to the type being tested. They are as follows.

|                          | Value | String | Number | Bool | Map | Slice | Error | Func |
|--------------------------|-------|--------|--------|------|-----|-------|-------|------|
| `ToBe`                   | Yes   | Yes    | Yes    | Yes  | Yes | Yes   | -     | -    |
| `ToEqual`                | Yes   | Yes    | Yes    | Yes  | -   | -     | -     | -    |
| `ToBeNil`                | Yes   | -      | -      | -    | Yes | Yes   | Yes   | -    |
| `ToBeEmpty`              | -     | Yes    | -      | -    | Yes | Yes   | -     | -    |
| `ToHaveLength`           | -     | Yes    | -      | -    | Yes | Yes   | -     | -    |
| `ToContain`              | -     | Yes    | -      | -    | Yes | Yes   | Yes   | -    |
| `ToContainAll`           | -     | -      | -      | -    | Yes | Yes   | -     | -    |
| `ToContainAny`           | -     | -      | -      | -    | Yes | Yes   | -     | -    |
| `ToMatch`                | -     | Yes    | -      | -    | -   | -     | Yes   | -    |
| `ToBeTrue`               | -     | -      | -      | Yes  | -   | -     | -     | -    |
| `ToBeFalse`              | -     | -      | -      | Yes  | -   | -     | -     | -    |
| `ToBeGreaterThan`        | -     | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeGreaterThanOrEqual` | -     | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeLessThan`           | -     | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeLessThanOrEqual`    | -     | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeBetween`            | -     | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeBetweenOrEqual`     | -     | -      | Yes    | -    | -   | -     | -     | -    |
| `ToHaveOccurred`         | -     | -      | -      | -    | -   | -     | Yes   | -    |
| `ToPanic`                | -     | -      | -      | -    | -   | -     | -     | Yes  |
| `ToPanicWithMessage`     | -     | -      | -      | -    | -   | -     | -     | Yes  |

Many categories have

* `ToBe(t, expected)` **tests for equality**, whereas
* `ToEqual(t, expected)` tests for equality ignoring whether the concrete types match or not - `ToBe` should be preferred but `ToEqual` is sometimes more convenient

Another group of related assertions is

* `ToBeNil(t)` verifies that the actual value is a nil pointer
* `ToBeEmpty(t)` verifies that the actual value has zero size
* `ToHaveLength(t, length)` verifies that the length of the actual value is as specified.

Containment tests are achieved using

* `ToContain(t, substring)` verifies that the substring is included within the actual string or error message. Maps are special: they can have an optional value, so `ToContain(t, key, [value])` tests
  that the key is present and that the the value, if present, must match what is held in the map.
* `ToContainAll(t, ...)` verifies that all the values are present, in any order
* `ToContainAny(t, ...)` verifies that any of the values is present
* `ToMatch(t, ...)` verifies that the string matches a regular expression

Boolean shorthands are

* `ToBeTrue(t)` is the same as `ToBe(t, true)`
* `ToBeFalse(t)` is the same as `ToBe(t, false)`

Numeric inequalities are

* `ToBeGreaterThan(t, threshold)` i.e. actual > threshold
* `ToBeGreaterThanOrEqual(t, threshold)` i.e. actual >= threshold
* `ToBeLessThan(t, threshold)` i.e. actual < threshold
* `ToBeLessThanOrEqual(t, threshold)` i.e. actual <= threshold
* `ToBeBetween(t, min, max)` i.e. min < actual < max
* `ToBeBetweenOrEqual(t, min, max)` i.e. min <= actual <= max

Note that these inequality assertions actually apply to all *ordered types*, which includes all int/uint types, float32/float64 and also string. All subtypes of ordered types are also included.

Errors are handled with `ToHaveOccurred(t)`, or more typically `Not().ToHaveOccurred(t)` (`Not()` is described below). These are equivalent to `Not().ToBeNil(t)` and `ToBeNil(t)`, respectively.

Functions that panic can be tested with a zero-argument function that calls the code under test and then uses `ToPanic()`. If `panic(value)` value is a string, `ToPanicWithMessage(t, substring)` can
check the actual message.

### Synonyms

For **Map**, `ToHaveSize(t, expected)` is a synonym for `ToHaveLength(t, expected)`.

## Negation Method

All categories include the general method `Not()`. This inverts the assertion defined by the `ToXxxx` method that follows it.

```go
    // `Not()` simply negates what follows.
expect.Number(v).Not().ToBe(t, 321)
```

## Conjunction Method

**Number** and **String** have `Or()` that allows multiple alternatives to be accepted.

* If any of them succeed, the test will pass.
* If all of them fail, the error message will list all the possibilities in the expected outcome.

There is no need for 'and' conjunctions because you simply add more assertions.

## Extra Information Methods

All categories include these general methods

* `Info(...)` provides information in the failure message, if there is one.
* `I(...)` is a terse synonym for `Info(...)`.

```go
    // The `Info` method can be helpful when testing inside a loop, for example.
// If `Not()` is also used, the natural order is to put `Not()` directly before the assertion:
var i int // some loop counter
expect.Number(v).Info("loop %d", i).Not().ToBe(t, 321)
```

**String** also has `Trim(n)` that truncates message strings if they exceed the specified length. When an expectation fails, the actual and expected strings are chopped at the front and/or back so
that the difference is visible in the failure message and the visible parts are not longer than the trim value.

```go
    expect.String(s).Trim(100).ToContain(t, " a very very long string ")
```

Both the actual and expected strings are truncated if their length is too long. If there is a mis-match, the error message scrolls the truncated string to ensure that the first difference is in view.

## Options for Controlling How The Comparisons Work

**Any**, **Map**, and **Slice** use [cmp.Equal](https://pkg.go.dev/github.com/google/go-cmp/cmp) under the hood. This is flexible, allowing for options to control how the comparison proceeds - for
example when considering how close floating point numbers need to be to be considered equal. There is a `Using(...)` method to specify what options it should use.

By default, the three options used are

* All fields in structs are compared, regardless of whether they exported or unexported; all structs in maps and slices are treated likewise.
* Floating point numbers are compared within the tolerance set by `ApproximateFloatFraction`.
* Maps/slices that are empty are treated the same as those that are nil.

## Status

This has been quite stable for some time and is available as a beta release.

## Origins

This API was mostly inspired by these

* [Gomega](https://github.com/onsi/gomega) had some great ideas but is overly complex to use.
* [Testify](https://github.com/stretchr/testify) was essentially simple to use but lacked flexibility, possibly because the API is not fluent, and had various gotchas.
