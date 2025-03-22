package expect

import "cmp"

// Number creates an ordering assertion. It also accepts strings.
//
// If more than one argument is passed, all subsequent arguments will be required to be nil/zero.
// This is convenient if you want to make an assertion on a method/function that returns a value and an error,
// a common pattern in Go.
func Number[O cmp.Ordered](value O, other ...any) OrderedType[O] {
	return OrderedType[O]{actual: value, assertion: assertion{other: other}}
}

// Info adds a description of the assertion to be included in any error message.
// If present, the third parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a OrderedType[O]) Info(info ...any) OrderedType[O] {
	a.info = makeInfo(info...)
	return a
}

// I is a synonym for [Info].
func (a OrderedType[O]) I(info ...any) OrderedType[O] {
	return a.Info(info...)
}

// Not inverts the assertion.
func (a OrderedType[O]) Not() OrderedType[O] {
	a.not = !a.not
	return a
}

// ToBe asserts that the actual values is equal to the expected value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBe(expected O, t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual == expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual != expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― to be ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

// ToBeGreaterThan asserts that the actual values is greater than the expected value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeGreaterThan(expected O, t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual > expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be greater than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual <= expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― to be greater than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

// ToBeLessThan asserts that the actual values is less than the expected value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeLessThan(expected O, t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual < expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be less than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual >= expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― to be less than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

// ToBeLessThanOrEqualTo asserts that the actual values is less than or equal to the expected value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeLessThanOrEqualTo(expected O, t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual <= expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be less than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual > expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― to be less than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

// ToBeGreaterThanOrEqualTo asserts that the actual values is greater than or equal to the expected value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeGreaterThanOrEqualTo(expected O, t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual >= expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be greater than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual < expected {
			t.Errorf("Expected%s %T ―――\n  %+v\n――― to be greater than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}
