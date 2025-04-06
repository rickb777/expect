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

If the value under test is of a type with a method `a.Equal(b)` (`a` and `b` having the same type), then the `Equal` method will be used. So this compares types such as `time.Time` correctly.

Otherwise it behaves like `reflect.DeepEqual`.

### expect.[String](https://pkg.go.dev/github.com/rickb777/expect#String)(actual ...)
This compares `string` and any subclass. It is more informative than **Any**, highlighting where the differences start.

### expect.[Number](https://pkg.go.dev/github.com/rickb777/expect#Number)(actual ...)
This compares `int` and all the signed/unsigned int and float length variants, plus all their subtypes. This provides inequality comparisons. It also supports  `string` because that is also is an ordered type. 

However, for near-equality testing of `float32` or `float64`, use **Any** instead because the tolerance can be specified.

### expect.[Bool](https://pkg.go.dev/github.com/rickb777/expect#Bool)(actual ...)
This compares `bool` and any subclass.

### expect.[Map](https://pkg.go.dev/github.com/rickb777/expect#Map)(actual ...)
This compares `map[K]V` where the map key `K` is a comparable type. **Map** provides more methods than **Any** above, but is otherwise very similar. 

### expect.[Slice](https://pkg.go.dev/github.com/rickb777/expect#Slice)(actual ...)
This compares `[]T` but only where `T` is a comparable type. Use **Any** for other slices.

**Slice** provides more methods than **Any**, but is otherwise very similar.

### expect.[Error](https://pkg.go.dev/github.com/rickb777/expect#Error)(... actual)
This compares `error` only.

### expect.[Func](https://pkg.go.dev/github.com/rickb777/expect#Func)(func)
This runs some function and checks whether it panicked.

## Application

The eight primary functions above all take the actual value under test as their input.

Other parameters can also be passed in. If any of these other parameters is non-nil (e.g. a non-nil `error`), the assertion will fail and give a corresponding error message. This allows, for example, the input to be a function with a multi-value return. 

Note that **Error** is different - it considers the *last* non-nil argument as its actual input. Any preceding arguments are ignored.

## Assertions

The assertions are all infinitive verbs, i.e. methods such as `ToBe`. Typical use is of the form

```go
    expect.Any(myValue).ToBe(t, expectedValue)
    expect.String(myValue).ToBe(t, expectedValue)
    expect.Number(myValue).ToBe(t, expectedValue)
    expect.Bool(myValue).ToBe(t, expectedValue)
    expect.Map(myValue).ToBe(t, expectedValue)
    expect.Slice(myValue).ToBe(t, expectedValue)
    expect.Error(myValue).ToBeNil(t)
    expect.Func(myValue).ToPanic(t)
```

All of the assertions require a `t Tester` parameter (see [Tester](https://pkg.go.dev/github.com/rickb777/expect#Tester)). Normally this will be `*testing.T` but you can use your own type if you need to embed this API in other assertion logic.

The assertions available are as follows.

|                          | Any | String | Number | Bool | Map | Slice | Error | Func |
|--------------------------|-----|--------|--------|------|-----|-------|-------|------|
| `ToBe`                   | Yes | Yes    | Yes    | Yes  | Yes | Yes   | -     | -    |
| `ToEqual`                | Yes | Yes    | -      | Yes  | -   | -     | -     | -    |
| `ToBeNil`                | Yes | -      | -      | -    | Yes | Yes   | Yes   | -    |
| `ToBeEmpty`              | -   | Yes    | -      | -    | Yes | Yes   | -     | -    |
| `ToHaveLength`           | -   | Yes    | -      | -    | Yes | Yes   | -     | -    |
| `ToContain`              | -   | Yes    | -      | -    | Yes | -     | Yes   | -    |
| `ToContainAll`           | -   | -      | -      | -    | Yes | Yes   | -     | -    |
| `ToContainAny`           | -   | -      | -      | -    | Yes | Yes   | -     | -    |
| `ToMatch`                | -   | Yes    | -      | -    | -   | -     | -     | -    |
| `ToBeTrue`               | -   | -      | -      | Yes  | -   | -     | -     | -    |
| `ToBeFalse`              | -   | -      | -      | Yes  | -   | -     | -     | -    |
| `ToBeGreaterThan`        | -   | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeGreaterThanOrEqual` | -   | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeLessThan`           | -   | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeLessThanOrEqual`    | -   | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeBetween`            | -   | -      | Yes    | -    | -   | -     | -     | -    |
| `ToBeBetweenOrEqual`     | -   | -      | Yes    | -    | -   | -     | -     | -    |
| `ToHaveOccurred`         | -   | -      | -      | -    | -   | -     | Yes   | -    |
| `ToPanic`                | -   | -      | -      | -    | -   | -     | -     | Yes  |
| `ToPanicWithMessage`     | -   | -      | -      | -    | -   | -     | -     | Yes  |

Many categories have

* `ToBe(t, expected)` **tests for equality**, whereas
* `ToEqual(t, expected)` tests for equality ignoring whether the concrete types match or not

Another group of related assertions is

* `ToBeNil(t)` verifies that the actual value is a nil pointer
* `ToBeEmpty(t)` verifies that the actual value has zero size
* `ToHaveLength(t, length)` verifies that the length of the actual value is as specified.

Containment tests are achieved using

* `ToContain(t, substring)` verifies that the substring is included within the actual string or error message. Maps are special: they can have an optional value, so `ToContain(t, key, [value])` tests that the key is present and that the the value, if present, must match what is held in the map.
* `ToContainAll(t, ...)` verifies that all the values are present, in any order
* `ToContainAny(t, ...)` verifies that any of the values are present, in any order
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

Note that these inequality assertions actually apply to all *ordered types*, which includes all int/uint types, float32/float64 and also string.

Errors are handled with `ToHaveOccurred(t)`, or more typically `Not().ToHaveOccurred(t)` (`Not()` is described below). These are equivalent to `Not().ToBeNil(t)` and `ToBeNil(t)`, respectively.

Functions that panic can be tested with a zero-argument function that calls the code under test and then uses `ToPanic()`. If `panic(value)` value is a string, `ToPanicWithMessage(t, substring)` can check the actual message. 

### Synonyms

For **Map**, `ToHaveSize(t, expected)` is a synonym for `ToHaveLength(t, expected)`.

## 'Or' Conjunction Method

**Number** and **String** have `Or()` that allows multiple alternatives to be accepted.

* If any of them succeed, the test will pass.
* If they all fail, the error message will list all the possibilities in the expected outcome.

There is no need for 'and' conjunctions because you simply add more assertions.

## Other Basic Methods

All categories include these general methods

* `Info(...)` provides information in the failure message, if there is one. There is a terse synonym `I(...)` too.
* `Not()` inverts the assertion defined by the `ToXxxx` method that follows it (these assertions are described above)

**String** also has `Trim(n)` that truncates message strings if they exceed the specified length.

## Options for Controlling How The Comparisons Work

**Any**, **Map**, and **Slice** use [cmp.Equal](https://pkg.go.dev/github.com/google/go-cmp/cmp) under the hood. This is flexible, allowing for options to control how the comparison proceeds - for example when considering how close floating point numbers need to be to be considered equal. There is a `Using(...)` method to specify what options it should use.

By default, the three options used are

 * All fields in structs are compared, regardless of whether they exported or unexported; all structs in maps and slices are treated likewise. 
 * Floating point numbers are compared within the tolerance set by `ApproximateFloatFraction`.
 * Maps/slices that are empty are treated the same as those that are nil.

## Status

This is now ready for beta testing.

## History

This API was mostly inspired by [Gomega](https://github.com/onsi/gomega), which had some great ideas but is overly complex to use.
