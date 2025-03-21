package expect

import (
	"fmt"
	gocmp "github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"reflect"
)

// ApproximateFloatFraction provides an option that compares any (a, b float32) or (a, b float64)
// pair. Change this if needed. See [cmpopts.EquateApprox] and [DefaultOptions].
var ApproximateFloatFraction = 1e-4

// DefaultOptions returns options used by [gocmp.Equal] for comparing values.
// You can also use [AnyType.Using] instead.
var DefaultOptions = func() gocmp.Options {
	return gocmp.Options{cmpopts.EquateApprox(ApproximateFloatFraction, 0), cmpopts.EquateEmpty()}
}

// Any creates an assertion for deep value comparison. This uses [cmp.Equal] so the manner of
// comparison can be tweaked using that API - see also [AnyType.Using]
//
// If present, the third parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func Any(t Tester, value any, info ...any) AnyType {
	return AnyType{t: t, actual: value, info: makeInfo(info...), opts: DefaultOptions()}
}

// Using replaces the default comparison options with those specified here.
// You can also set [DefaultOptions] instead.
func (a AnyType) Using(opt ...gocmp.Option) AnyType {
	a.opts = opt
	return a
}

// Not inverts the assertion.
func (a AnyType) Not() AnyType {
	a.not = !a.not
	return a
}

// ToBe asserts that the actual and expected data have the same values and types.
func (a AnyType) ToBe(expected any) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	match := gocmp.Equal(a.actual, expected, a.opts)

	if (!a.not && !match) || (a.not && match) {
		a.t.Errorf("Expected%s %T ―――\n%s――― %sto equal ―――\n%s", preS(a.info), a.actual, verbatim(a.actual), notS(a.not), verbatim(expected))
	}
}

// ToEqual asserts that the actual and expected data have the same underlying values, but the
// concrete types may differ. For example, an int and a uint with the same numeric value are
// considered equal.
func (a AnyType) ToEqual(expected any) {
	if h, ok := a.t.(helper); ok {
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
		a.t.Errorf("Expected%s %T ―――\n%s――― %sto be equivalent to %T ―――\n%s", preS(a.info),
			a.actual, verbatim(a.actual), notS(a.not), expected, verbatim(expected))
	}
}

//=================================================================================================

func verbatim(v any) string {
	a := fmt.Sprintf("  %+v\n", v)
	b := fmt.Sprintf("  %#v\n", v)
	if a == b {
		return blank(a)
	}
	return a + b
}
