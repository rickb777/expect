package expect

// Bool creates a boolean assertion.
// If present, the second parameter should be a string. If this is a format
// string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func Bool(t Tester, value bool, info ...any) BoolType {
	return BoolType{t: t, actual: value, info: makeInfo(info...)}
}

// Not inverts the assertion.
func (a BoolType) Not() BoolType {
	a.not = !a.not
	return a
}

// ToBeTrue asserts that the actual value is true.
func (a BoolType) ToBeTrue() {
	a.ToBe(true)
}

// ToBeFalse asserts that the actual value is true.
func (a BoolType) ToBeFalse() {
	a.ToBe(false)
}

// ToBe asserts that the actual value is as expected.
func (a BoolType) ToBe(expected bool) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if (!a.not && a.actual != expected) || (a.not && a.actual == expected) {
		a.t.Errorf("Expected%s %sto be %v\n", preS(a.info), notS(a.not), expected)
	}
}
