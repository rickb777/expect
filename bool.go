package expect

// Bool creates a boolean assertion.
//
// If more than one argument is passed, all subsequent arguments will be required to be nil/zero.
// This is convenient if you want to make an assertion on a method/function that returns a value and an error,
// a common pattern in Go.
func Bool[B ~bool](value B, other ...any) BoolType[B] {
	return BoolType[B]{actual: value, assertion: assertion{other: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a BoolType[B]) Info(info any, other ...any) BoolType[B] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a BoolType[B]) I(info any, other ...any) BoolType[B] {
	return a.Info(info, other...)
}

// Not inverts the assertion.
func (a BoolType[B]) Not() BoolType[B] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToBeTrue asserts that the actual value is true.
// The tester is normally [*testing.B].
func (a BoolType[B]) ToBeTrue(t Tester) {
	a.ToBe(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBeFalse asserts that the actual value is true.
// The tester is normally [*testing.B].
func (a BoolType[B]) ToBeFalse(t Tester) {
	a.ToBe(t, false)
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected items have the same values and types.
// The tester is normally [*testing.B].
func (a BoolType[B]) ToBe(t Tester, expected B) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.ToEqual(t, bool(expected))
}

//-------------------------------------------------------------------------------------------------

// ToEqual asserts that the actual and expected items have the same values and similar types.
// The tester is normally [*testing.B].
func (a BoolType[B]) ToEqual(t Tester, expected bool) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if (!a.not && bool(a.actual) != expected) || (a.not && bool(a.actual) == expected) {
		t.Errorf("Expected%s %sto be %v\n", preS(a.info), notS(bool(a.not)), expected)
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}
