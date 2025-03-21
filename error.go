package expect

import "strings"

// Error creates an error assertion.
//
// If present, the third parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func Error(t Tester, value error, info ...any) ErrorType {
	return ErrorType{t: t, actual: value, info: makeInfo(info...)}
}

// Not inverts the assertion.
func (a ErrorType) Not() ErrorType {
	a.not = !a.not
	return a
}

// ToBeNil asserts that the error did not occur.
func (a ErrorType) ToBeNil() {
	a.toHaveOccurred(!a.not)
}

// ToHaveOccurred asserts that the error occurred.
func (a ErrorType) ToHaveOccurred() {
	a.toHaveOccurred(a.not)
}

// ToHaveOccurred asserts that the error occurred.
func (a ErrorType) toHaveOccurred(not bool) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if not {
		if a.actual != nil {
			a.t.Fatalf("Expected%s error ...\n  %s\n... not to have occurred.\n", preS(a.info), blank(a.actual.Error()))
		}
	} else {
		if a.actual == nil {
			a.t.Errorf("Expected%s error to have occurred.\n", preS(a.info))
		}
	}
}

// ToContain asserts that the error occurred and its message contains the substring.
func (a ErrorType) ToContain(substring string) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	if a.actual == nil {
		a.t.Errorf("Expected%s error to have occurred.\n", preS(a.info))
	} else {
		msg := a.actual.Error()
		match := strings.Contains(msg, substring)
		if (!a.not && !match) || (a.not && match) {
			a.t.Errorf("Expected%s error ...\n  %s\n... %sto contain ...\n  %s\n", preS(a.info), blank(msg), notS(a.not), substring)
		}
	}
}
