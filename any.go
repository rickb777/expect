package expect

import (
	gocmp "github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"reflect"
)

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
// (see [AnyType.ToBe] and [AnyType.ToEqual]).
//
// This uses [gocmp.Equal] so the manner of comparison can be tweaked using that API - see also [AnyType.Using]
func Any[T any](value T, other ...any) AnyType[T] {
	return AnyType[T]{actual: value, opts: DefaultOptions(), assertion: assertion{other: other}}
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

// ToBe asserts that the actual and expected data have the same values and types.
// The tester is normally [*testing.T].
func (a AnyType[T]) ToBe(t Tester, expected T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toEqual(t, "be", expected)
}

//-------------------------------------------------------------------------------------------------

// ToEqual asserts that the actual and expected data have the same values and similar types.
// The tester is normally [*testing.T].
func (a AnyType[T]) ToEqual(t Tester, expected any) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toEqual(t, "equal", expected)
}

//-------------------------------------------------------------------------------------------------

func (a AnyType[T]) toEqual(t Tester, what string, expected any) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	convertedActual := a.actual

	if a.actual != nil &&
		expected != nil &&
		reflect.TypeOf(a.actual).ConvertibleTo(reflect.TypeOf(expected)) {
		convertedActual = reflect.ValueOf(a.actual).Convert(reflect.TypeOf(expected)).Interface()
	}

	match := gocmp.Equal(convertedActual, expected, a.opts)

	if (!a.not && !match) || (a.not && match) {
		t.Errorf("Expected%s %T ―――\n%s――― %sto %s %T ―――\n%s", preS(a.info),
			a.actual, verbatim(a.actual), notS(a.not), what, expected, verbatim(expected))
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}
