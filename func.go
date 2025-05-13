package expect

import "strings"

// FuncType is used for assertions about functions.
type FuncType struct {
	actual func()
	assertion
}

// Func wraps a function that can test for panics etc.
func Func(value func()) FuncType {
	return FuncType{actual: value}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a FuncType) Info(info any, other ...any) FuncType {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a FuncType) I(info any, other ...any) FuncType {
	return a.Info(info, other...)
}

// Not inverts the assertion.
func (a FuncType) Not() FuncType {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToPanic asserts that the function did / did not panic.
// The tester is normally [*testing.B].
func (a FuncType) ToPanic(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	defer func() {
		if e := recover(); e != nil {
			if a.not {
				a.describeActualExpected1("not to panic.\n")
				a.applyAll(t)
			} else {
				a.passes++
			}
		}
	}()

	a.actual() // function under test

	if !a.not {
		a.describeActualExpected1("to panic.\n")
		a.applyAll(t)
	} else {
		a.passes++
	}
}

//-------------------------------------------------------------------------------------------------

// ToPanicWithMessage asserts that the function did panic.
// It is not useful to use [FuncType.Not] with this.
// The substring is used to check that the panic passed a string containing that value.
// The tester is normally [*testing.B].
func (a FuncType) ToPanicWithMessage(t Tester, substring string) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	if a.not {
		panic("Func().ToPanicWithMessage() does not allow Not() because of ambiguous meaning")
	}

	defer func() {
		if e := recover(); e != nil {
			if s, ok := e.(string); ok {
				if !strings.Contains(s, substring) {
					a.describeActualExpected1("to panic with a message containing ―――\n%s\n――― but got ―――\n%s\n",
						substring, s)
				}
			} else {
				a.describeActualExpected1("to panic with a string containing ―――\n%s\n――― but got %T ―――\n%v\n",
					substring, e, e)
			}
			a.applyAll(t)
		}
	}()

	a.actual() // function under test

	a.describeActualExpected1("to panic.\n")
	a.applyAll(t)
}
