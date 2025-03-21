package expect

import "cmp"

// Number creates an ordering assertion. It also accepts strings.
//
// If present, the third parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func Number[T cmp.Ordered](t Tester, value T, info ...any) OrderedType[T] {
	return OrderedType[T]{t: t, actual: value, info: makeInfo(info...)}
}

// Not inverts the assertion.
func (a OrderedType[T]) Not() OrderedType[T] {
	a.not = !a.not
	return a
}

// ToBe asserts that the actual values is equal to the expected value.
func (a OrderedType[T]) ToBe(expected T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual == expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual != expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― to be ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}
}

// ToBeGreaterThan asserts that the actual values is greater than the expected value.
func (a OrderedType[T]) ToBeGreaterThan(expected T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual > expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be greater than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual <= expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― to be greater than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}
}

// ToBeLessThan asserts that the actual values is less than the expected value.
func (a OrderedType[T]) ToBeLessThan(expected T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual < expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be less than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual >= expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― to be less than ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}
}

// ToBeLessThanOrEqualTo asserts that the actual values is less than or equal to the expected value.
func (a OrderedType[T]) ToBeLessThanOrEqualTo(expected T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual <= expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be less than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual > expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― to be less than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}
}

// ToBeGreaterThanOrEqualTo asserts that the actual values is greater than or equal to the expected value.
func (a OrderedType[T]) ToBeGreaterThanOrEqualTo(expected T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if a.not {
		if a.actual >= expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― not to be greater than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	} else {
		if a.actual < expected {
			a.t.Errorf("Expected%s %T ―――\n  %+v\n――― to be greater than or equal to ―――\n  %+v\n", preS(a.info), a.actual, a.actual, expected)
		}
	}
}
