package expect

import (
	"strings"
)

// ErrorType is used for assertions about errors.
type ErrorType struct {
	actual error
	assertion
}

// Error creates an error assertion. This considers the last error it finds in the supplied parameters.
// All other parameters are ignored.
func Error(value any, other ...any) ErrorType {
	foundNil := false

	for i := len(other) - 1; i >= 0; i-- {
		if other[i] == nil {
			foundNil = true
		} else if err, ok := other[i].(error); ok {
			return ErrorType{actual: err}
		}
	}

	if foundNil {
		return ErrorType{}
	}

	if value == nil {
		return ErrorType{}
	} else if err, ok := value.(error); ok {
		return ErrorType{actual: err}
	}

	panic("No parameter was an error.")
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a ErrorType) Info(info any, other ...any) ErrorType {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a ErrorType) I(info any, other ...any) ErrorType {
	return a.Info(info, other...)
}

// Not inverts the assertion.
func (a ErrorType) Not() ErrorType {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToBeNil asserts that the error did not occur.
// The tester is normally [*testing.T].
func (a ErrorType) ToBeNil(t Tester) {
	a.toHaveOccurred(t, !a.not)
}

//-------------------------------------------------------------------------------------------------

// ToHaveOccurred asserts that the error occurred.
// The tester is normally [*testing.T].
func (a ErrorType) ToHaveOccurred(t Tester) {
	a.toHaveOccurred(t, a.not)
}

//-------------------------------------------------------------------------------------------------

// ToHaveOccurred asserts that the error occurred.
func (a ErrorType) toHaveOccurred(t Tester, not bool) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if not {
		if a.actual != nil {
			t.Fatalf("Expected%s error ―――\n  %s\n――― not to have occurred.\n",
				preS(a.info), blank(a.actual.Error()))
		}
	} else {
		if a.actual == nil {
			t.Errorf("Expected%s error to have occurred.\n", preS(a.info))
		}
	}
}

//-------------------------------------------------------------------------------------------------
// TODO ToMatch
//-------------------------------------------------------------------------------------------------

// ToContain asserts that the error occurred and its message contains the substring.
// The tester is normally [*testing.T].
func (a ErrorType) ToContain(t Tester, substring string) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if a.actual == nil {
		t.Errorf("Expected%s error to have occurred but there was no error.\n", preS(a.info))
	} else {
		msg := a.actual.Error()
		match := strings.Contains(msg, substring)
		if (!a.not && !match) || (a.not && match) {
			t.Errorf("Expected%s error ―――\n  %s\n――― %sto contain ―――\n  %s\n",
				preS(a.info), blank(msg), notS(a.not), substring)
		}
	}
}
