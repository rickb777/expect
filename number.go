package expect

import "cmp"

// OrderedType is used for assertions about numbers and other ordered types.
type OrderedType[O cmp.Ordered] struct {
	actual O
	assertion
}

// Number creates an ordering assertion. It accepts all numbers, and also coincidentally accepts strings.
// Its methods are the full set of ordering comparisons, i.e. >, >=, <, <=, ==, and !=.
//
// If more than one argument is passed, all subsequent arguments will be required to be nil/zero.
// This is convenient if you want to make an assertion on a method/function that returns a value and an error,
// a common pattern in Go.
func Number[O cmp.Ordered](value O, other ...any) OrderedType[O] {
	return OrderedType[O]{actual: value, assertion: assertion{otherActual: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a OrderedType[O]) Info(info any, other ...any) OrderedType[O] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a OrderedType[O]) I(info any, other ...any) OrderedType[O] {
	return a.Info(info, other...)
}

// Not inverts the assertion.
func (a OrderedType[O]) Not() OrderedType[O] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected numbers have the same values and types.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBe(t Tester, expected O) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if a.not {
		if a.actual == expected {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be ―――\n  %+v\n", expected)
		} else {
			a.passes++
		}
	} else {
		if a.actual != expected {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be ―――\n  %+v\n", expected)
		} else {
			a.passes++
		}
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeGreaterThan asserts that the actual values is greater than the threshold value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeGreaterThan(t Tester, threshold O) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if a.not {
		if a.actual > threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be greater than ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	} else {
		if a.actual <= threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be greater than ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeLessThan asserts that the actual values is less than the threshold value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeLessThan(t Tester, threshold O) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if a.not {
		if a.actual < threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be less than ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	} else {
		if a.actual >= threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be less than ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeLessThanOrEqual asserts that the actual values is less than or equal to the threshold value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeLessThanOrEqual(t Tester, threshold O) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if a.not {
		if a.actual <= threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be less than or equal to ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	} else {
		if a.actual > threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be less than or equal to ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeGreaterThanOrEqual asserts that the actual values is greater than or equal to the threshold value.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeGreaterThanOrEqual(t Tester, threshold O) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if a.not {
		if a.actual >= threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be greater than or equal to ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	} else {
		if a.actual < threshold {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be greater than or equal to ―――\n  %+v\n", threshold)
		} else {
			a.passes++
		}
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeBetweenOrEqual asserts that the actual values is between two threshold values.
// The assertion succeeds if minimum <= value <= maximum.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeBetweenOrEqual(t Tester, minimum, maximum O) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if minimum > maximum {
		a.describeActual1line("Impossible test%s %T: minimum %v > maximum %v.\n",
			preS(a.info), a.actual, minimum, maximum)
	} else if a.not {
		if minimum <= a.actual && a.actual <= maximum {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be between ―――\n  %+v … %v (inclusive)\n", minimum, maximum)
		} else {
			a.passes++
		}
	} else {
		if a.actual < minimum || a.actual > maximum {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be between ―――\n  %+v … %v (inclusive)\n", minimum, maximum)
		} else {
			a.passes++
		}
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeBetween asserts that the actual values is between two threshold values.
// The assertion succeeds if minimum < value < maximum.
// The tester is normally [*testing.T].
func (a OrderedType[O]) ToBeBetween(t Tester, minimum, maximum O) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if minimum >= maximum {
		a.describeActual1line("Impossible test%s %T: minimum %v >= maximum %v.\n",
			preS(a.info), a.actual, minimum, maximum)
	} else if a.not {
		if minimum < a.actual && a.actual < maximum {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be between ―――\n  %+v … %v (exclusive)\n", minimum, maximum)
		} else {
			a.passes++
		}
	} else {
		if a.actual <= minimum || a.actual >= maximum {
			a.describeActualMulti("Expected%s %T ―――\n  %+v\n", preS(a.info), a.actual, a.actual)
			a.addExpectation("to be between ―――\n  %+v … %v (exclusive)\n", minimum, maximum)
		} else {
			a.passes++
		}
	}

	a.applyAll(t)
}
