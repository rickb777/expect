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

There are four primary categories:

 * Strings - `string` and any subclass
 * Numbers - all ordered types such as `int`, `float32`, plus all the signed/unsigned int and float length variants, plus all their subtypes, plus `string` and any subclass
 * Bools - `bool` and any subclass
 * Errors - `error` only
 * Everything else - structs, maps, arrays, slices as handled by [cmp.Equal](https://pkg.go.dev/github.com/google/go-cmp/cmp)
