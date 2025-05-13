package expect

import (
	"cmp"
	"reflect"
)

// OrderedType is used for assertions about numbers and other ordered types.
type OrderedType[O cmp.Ordered] struct {
	actual O
	assertion
}

// OrderedOr is only used for conjunction concatenation (see [OrderedOr.Or]).
type OrderedOr[O cmp.Ordered] struct {
	main           *OrderedType[O]
	passes         int
	unwantedTester Tester
}

// Number creates an ordering assertion. It accepts all numbers, and also coincidentally accepts strings.
// Its methods are the full set of ordering comparisons, i.e. >, >=, <, <=, ==, and !=.
//
// If more than one argument is passed, all subsequent arguments will be required to be nil/zero.
// This is convenient if you want to make an assertion on a method/function that returns a value and an error,
// a common pattern in Go.
func Number[O cmp.Ordered](value O, other ...any) *OrderedType[O] {
	return &OrderedType[O]{actual: value, assertion: assertion{otherActual: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a *OrderedType[O]) Info(info any, other ...any) *OrderedType[O] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a *OrderedType[O]) I(info any, other ...any) *OrderedType[O] {
	return a.Info(info, other...)
}

// Not inverts the assertion.
func (a *OrderedType[O]) Not() *OrderedType[O] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected numbers have the same values and types.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToBe(t Tester, expected O) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if a.not {
		if a.actual == expected {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be ―――\n%+v\n", expected)
			return a.conjunction(t, false)
		}
	} else {
		if a.actual != expected {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be ―――\n%+v\n", expected)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToEqual asserts that the actual and expected numbers have the same values and similar types.
// The expected value must be some numeric type.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToEqual(t Tester, expected any) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	match := false

	expectedType := reflect.TypeOf(expected)
	switch expectedType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		// ok

	default:
		a.describeActualExpectedM("%T ―――\n%+v\n", expected, expected)
		a.addExpectation("type must be int, uint, or float (of any length) ―――\n")
		return a.conjunction(t, false)
	}

	if expected != nil &&
		reflect.TypeFor[O]().ConvertibleTo(expectedType) {
		convertedActual := reflect.ValueOf(a.actual).Convert(expectedType).Interface()
		match = convertedActual == expected
	}

	if a.not {
		if match {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be ―――\n%+v\n", expected)
			return a.conjunction(t, false)
		}
	} else {
		if !match {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be ―――\n%+v\n", expected)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBeGreaterThan asserts that the actual values is greater than the threshold value.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToBeGreaterThan(t Tester, threshold O) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if a.not {
		if a.actual > threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be greater than ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	} else {
		if a.actual <= threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be greater than ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBeLessThan asserts that the actual values is less than the threshold value.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToBeLessThan(t Tester, threshold O) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if a.not {
		if a.actual < threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be less than ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	} else {
		if a.actual >= threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be less than ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBeLessThanOrEqual asserts that the actual values is less than or equal to the threshold value.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToBeLessThanOrEqual(t Tester, threshold O) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if a.not {
		if a.actual <= threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be less than or equal to ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	} else {
		if a.actual > threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be less than or equal to ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBeGreaterThanOrEqual asserts that the actual values is greater than or equal to the threshold value.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToBeGreaterThanOrEqual(t Tester, threshold O) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if a.not {
		if a.actual >= threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be greater than or equal to ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	} else {
		if a.actual < threshold {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be greater than or equal to ―――\n%+v\n", threshold)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBeBetweenOrEqual asserts that the actual values is between two threshold values.
// The assertion succeeds if minimum <= value <= maximum.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToBeBetweenOrEqual(t Tester, minimum, maximum O) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if minimum > maximum {
		a.describeActual("Impossible test%s %T: minimum %v > maximum %v.\n",
			preS(a.info), a.actual, minimum, maximum)
		return a.conjunction(t, false)
	} else if a.not {
		if minimum <= a.actual && a.actual <= maximum {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be between ―――\n%+v … %v (inclusive)\n", minimum, maximum)
			return a.conjunction(t, false)
		}
	} else {
		if a.actual < minimum || a.actual > maximum {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be between ―――\n%+v … %v (inclusive)\n", minimum, maximum)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBeBetween asserts that the actual values is between two threshold values.
// The assertion succeeds if minimum < value < maximum.
// The tester is normally [*testing.T].
func (a *OrderedType[O]) ToBeBetween(t Tester, minimum, maximum O) *OrderedOr[O] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if minimum >= maximum {
		a.describeActual("Impossible test%s %T: minimum %v >= maximum %v.\n",
			preS(a.info), a.actual, minimum, maximum)
		return a.conjunction(t, false)
	} else if a.not {
		if minimum < a.actual && a.actual < maximum {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be between ―――\n%+v … %v (exclusive)\n", minimum, maximum)
			return a.conjunction(t, false)
		}
	} else {
		if a.actual <= minimum || a.actual >= maximum {
			a.describeActualExpectedM("%T ―――\n%+v\n", a.actual, a.actual)
			a.addExpectation("to be between ―――\n%+v … %v (exclusive)\n", minimum, maximum)
			return a.conjunction(t, false)
		}
	}

	return a.conjunction(t, true)
}

//=================================================================================================

func (a *OrderedType[O]) conjunction(t Tester, pass bool) *OrderedOr[O] {
	if pass {
		a.passes++
	}

	if t == nil {
		return &OrderedOr[O]{main: a, passes: a.passes} // defer evaluation
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}
	a.applyAll(t)
	return &OrderedOr[O]{main: a, passes: a.passes, unwantedTester: t}
}

//-------------------------------------------------------------------------------------------------

func (or *OrderedOr[O]) Or() *OrderedType[O] {
	if or != nil {
		if or.unwantedTester == nil {
			return or.main // following assertions are active
		}
		or.unwantedTester.Fatal(incorrectTestConjunction)
	}
	return nil // following assertions are no-op
}
