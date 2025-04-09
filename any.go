package expect

import (
	"fmt"
	gocmp "github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"reflect"
	"strings"
)

// AnyType is used for equality assertions for any type.
type AnyType[T any] struct {
	opts   gocmp.Options
	actual any
	assertion
}

// ApproximateFloatFraction provides an option that compares any (a, b float32) or (a, b float64)
// pair. Change this if needed. See [cmpopts.EquateApprox] and [DefaultOptions].
//
// If more than one argument is passed, all subsequent arguments will be required to be nil/zero.
// This is convenient if you want to make an assertion on a method/function that returns a value and an error,
// a common pattern in Go.
var ApproximateFloatFraction = 1e-4

// DefaultOptions returns options used by [gocmp.Equal] for comparing values.
// The default options
//
//   - sets the threshold for float comparison to [ApproximateFloatFraction]
//   - sets empty and nil maps or slices to be treated the same
//
// You can also use [AnyType.Using] instead.
var DefaultOptions = func() gocmp.Options {
	return gocmp.Options{cmpopts.EquateApprox(ApproximateFloatFraction, 0), cmpopts.EquateEmpty()}
}

// Any creates an assertion for deep value comparison of any type. This is very flexible but only
// provides methods to determine whether a value is equal (or not equal) to what's expected
// (see [AnyType.ToBe], [AnyType.ToBeNil] and [AnyType.ToEqual]).
//
// For alternative comparisons, see the more-specialized [String], [Number], [Bool], [Slice],
// [Map], [Error] and [Func] functions.
//
// Any uses [gocmp.Equal] so the manner of comparison can be tweaked using that API - see also [AnyType.Using]
//
//   - If the values have an Equal method of the form "(T) Equal(T) bool" or
//     "(T) Equal(I) bool" where T is assignable to I, then it uses the result of
//     x.Equal(y) even if x or y is nil.
//
//   - Lastly, it tries to compare x and y based on their basic kinds.
//     Simple kinds like booleans, integers, floats, complex numbers, strings,
//     and channels are compared using the equivalent of the == operator in Go.
//     Functions are only equal if they are both nil, otherwise they are unequal.
//
// Structs are equal if recursively calling Equal on all fields report equal. All
// struct fields are compared and this is repeated recursively. Unless the compare
// options are changed, it does not matter whether fields exported on unexported.
//
// Slices are equal if they are both nil or both non-nil, where recursively
// calling Equal on all non-ignored slice or array elements report equal.
// Unless the compare options are changed, empty non-nil slices and nil slices
// are equal.
//
// Maps are equal if they are both nil or both non-nil, where recursively
// calling Equal on all non-ignored map entries report equal.
// Map keys are equal according to the == operator.
// To use custom comparisons for map keys, consider using
// [github.com/google/go-cmp/cmp/cmpopts.SortMaps].
// Unless the compare options are changed, empty non-nil maps and nil maps
// are equal.
//
// Pointers and interfaces are equal if they are both nil or both non-nil,
// where they have the same underlying concrete type and recursively
// calling Equal on the underlying values reports equal.
//
// Before recursing into a pointer, slice element, or map, the current path
// is checked to detect whether the address has already been visited.
// If there is a cycle, then the pointed at values are considered equal
// only if both addresses were previously visited in the same path step.
func Any[T any](value T, other ...any) AnyType[T] {
	return AnyType[T]{actual: value, opts: DefaultOptions(), assertion: assertion{otherActual: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a AnyType[T]) Info(info any, other ...any) AnyType[T] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a AnyType[T]) I(info any, other ...any) AnyType[T] {
	return a.Info(info, other...)
}

// Using replaces the default comparison options with those specified here.
// You can also set [DefaultOptions] instead.
func (a AnyType[T]) Using(opt ...gocmp.Option) AnyType[T] {
	a.opts = opt
	return a
}

// Not inverts the assertion.
func (a AnyType[T]) Not() AnyType[T] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

func isNilish(val any) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return false
}

// ToBeNil asserts that the actual value is nil / is not nil.
// The tester is normally [*testing.T].
func (a AnyType[T]) ToBeNil(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if !a.not && !isNilish(a.actual) {
		a.describeActualExpected1("%T ―――\n%s――― to be nil.\n", a.actual, verbatim(a.actual))
	} else if a.not && isNilish(a.actual) {
		a.describeActualExpected1("%T not to be nil.\n", a.actual)
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected data have the same values and types.
// The tester is normally [*testing.T].
func (a AnyType[T]) ToBe(t Tester, expected T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toEqual(t, "to be", a.actual, expected, false)
}

//-------------------------------------------------------------------------------------------------

// ToEqual asserts that the actual and expected data have the same values and similar types.
// The tester is normally [*testing.T].
func (a AnyType[T]) ToEqual(t Tester, expected any) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	convertedActual := a.actual
	differentType := false

	if a.actual != nil && expected != nil &&
		reflect.TypeOf(a.actual).ConvertibleTo(reflect.TypeOf(expected)) {
		convertedActual = reflect.ValueOf(a.actual).Convert(reflect.TypeOf(expected)).Interface()
		differentType = true
	}

	a.toEqual(t, "to equal", convertedActual, expected, differentType)
}

//-------------------------------------------------------------------------------------------------

func (a AnyType[T]) toEqual(t Tester, what string, actual, expected any, differentType bool) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	isStruct := actual != nil && reflect.TypeOf(actual).Kind() == reflect.Struct

	opts := append(a.opts, allowUnexported(gatherTypes(nil, actual, expected)))

	diffs := gocmp.Diff(expected, actual, opts)

	expectedType := ""
	if differentType {
		expectedType = fmt.Sprintf(" %T", expected)
	}

	if !a.not && diffs != "" {
		if isStruct {
			a.describeActualExpected1("struct %s as shown (-want, +got) ―――\n", what)
			a.addExpectation("%s", strings.ReplaceAll(diffs, " ", " "))
		} else {
			a.describeActualExpectedM("%T ―――\n%s", a.actual, verbatim(a.actual))
			a.addExpectation("%s%s ―――\n%s", what, expectedType, verbatim(expected))
		}
	} else if a.not && diffs == "" {
		a.describeActualExpected1("%T ", a.actual)
		a.addExpectation("%s%s ―――\n%s", what, expectedType, verbatim(expected))
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

type typeSet map[reflect.Type]bool

func gatherTypes(m typeSet, types ...interface{}) typeSet {
	if m == nil {
		m = make(typeSet)
	}
	for _, typ := range types {
		discoverTypes(reflect.TypeOf(typ), m)
	}
	return m
}

// allowUnexported returns an [Option] that allows [Equal] to forcibly introspect
// unexported fields of the specified struct types.
func allowUnexported(m typeSet) gocmp.Option {
	return gocmp.Exporter(func(t reflect.Type) bool { return m[t] })
}

func discoverTypes(t reflect.Type, m map[reflect.Type]bool) {
	if t != nil {
		switch t.Kind() {
		case reflect.Struct:
			if _, exists := m[t]; !exists {
				m[t] = true
				for i := 0; i < t.NumField(); i++ {
					discoverTypes(t.Field(i).Type, m)
				}
			}

		case reflect.Slice, reflect.Pointer:
			discoverTypes(t.Elem(), m)

		case reflect.Map:
			discoverTypes(t.Key(), m)
			discoverTypes(t.Elem(), m)
		}
	}
}
