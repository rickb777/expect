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
// If present, the third parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a BoolType[B]) Info(info ...any) BoolType[B] {
	a.info = makeInfo(info...)
	return a
}

// I is a synonym for [Info].
func (a BoolType[B]) I(info ...any) BoolType[B] {
	return a.Info(info...)
}

// Not inverts the assertion.
func (a BoolType[B]) Not() BoolType[B] {
	a.not = !a.not
	return a
}

// ToBeTrue asserts that the actual value is true.
// The tester is normally [*testing.B].
func (a BoolType[B]) ToBeTrue(t Tester) {
	a.ToBe(true, t)
}

// ToBeFalse asserts that the actual value is true.
// The tester is normally [*testing.B].
func (a BoolType[B]) ToBeFalse(t Tester) {
	a.ToBe(false, t)
}

// ToBe asserts that the actual value is as expected.
// The tester is normally [*testing.B].
func (a BoolType[B]) ToBe(expected B, t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if (!a.not && a.actual != expected) || (a.not && a.actual == expected) {
		t.Errorf("Expected%s %sto be %v\n", preS(a.info), notS(bool(a.not)), expected)
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}
