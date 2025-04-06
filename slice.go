package expect

import (
	"fmt"
	gocmp "github.com/google/go-cmp/cmp"
)

// SliceType is used for assertions about slices.
type SliceType[T comparable] struct {
	opts   gocmp.Options
	actual []T
	assertion
}

// Slice creates an assertion for deep value comparison of slices of any comparable type.
//
// This uses [gocmp.Equal] so the manner of comparison can be tweaked using that API - see also [SliceType.Using]
func Slice[T comparable](value []T, other ...any) SliceType[T] {
	return SliceType[T]{actual: value, opts: DefaultOptions(), assertion: assertion{otherActual: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a SliceType[T]) Info(info any, other ...any) SliceType[T] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a SliceType[T]) I(info any, other ...any) SliceType[T] {
	return a.Info(info, other...)
}

// Using replaces the default comparison options with those specified here.
// You can also set [DefaultOptions] instead.
func (a SliceType[T]) Using(opt ...gocmp.Option) SliceType[T] {
	a.opts = opt
	return a
}

// Not inverts the assertion.
func (a SliceType[T]) Not() SliceType[T] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToBeNil asserts that the actual value is nil / is not nil.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToBeNil(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	if !a.not && !isNilish(a.actual) {
		a.describeActualMulti("Expected%s %T len:%d ―――\n%s", preS(a.info), a.actual, len(a.actual), verbatim(a.actual))
		a.addExpectation("to be nil.\n")
	} else if a.not && isNilish(a.actual) {
		a.describeActual1line("Expected%s %T ", preS(a.info), a.actual)
		a.addExpectation("to be nil.\n")
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected slices have the same values and types.
// The values must be in the same order.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToBe(t Tester, expected ...T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	types := gatherTypes(nil, a.actual, expected)
	opts := append(a.opts, allowUnexported(types))

	match := gocmp.Equal(a.actual, expected, opts)

	if (!a.not && !match) || (a.not && match) {
		diff := findFirstDiff(a.actual, expected)
		a.describeActualMulti("Expected%s %T len:%d ―――\n%s", preS(a.info), a.actual, len(a.actual), verbatim(a.actual))
		a.addExpectation("to be len:%d ―――\n%s%s", len(expected), verbatim(expected), firstDifferenceInfo(diff))
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeEmpty asserts that the slice has zero length.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToBeEmpty(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toHaveLength(t, 0, "to be empty.", true)
}

//-------------------------------------------------------------------------------------------------

// ToHaveLength asserts that the slice has the expected length.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToHaveLength(t Tester, expected int) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toHaveLength(t, expected, fmt.Sprintf("to have length %d.", expected), true)
}

//-------------------------------------------------------------------------------------------------

func (a SliceType[T]) toHaveLength(t Tester, expected int, what string, showActual bool) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	actual := len(a.actual)

	if (!a.not && actual != expected) || (a.not && actual == expected) {
		if showActual && len(a.actual) > 0 {
			a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n", preS(a.info), a.actual, len(a.actual), a.actual)
			a.addExpectation("%s\n", what)
		} else {
			a.describeActual1line("Expected%s %T len:%d ", preS(a.info), a.actual, len(a.actual))
			a.addExpectation("%s\n", what)
		}
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToContainAll asserts that the slice contains all of the values listed.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToContainAll(t Tester, expected ...T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	found := make([]T, 0, len(expected))
	missing := make([]T, 0, len(expected))
	for _, v := range expected {
		if sliceContains(a.actual, v) {
			found = append(found, v)
		} else {
			missing = append(missing, v)
		}
	}

	if !a.not && len(missing) > 0 {
		if len(found) == 0 {
			a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but none were found.\n", allN.FormatInt(len(expected)))
		} else if len(found) < len(expected)/2 {
			a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but only %s found ―――\n  %v\n",
				allN.FormatInt(len(expected)), theseWere.FormatInt(len(found)), found)
		} else {
			a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but %s missing ―――\n  %v\n",
				allN.FormatInt(len(expected)), theseWere.FormatInt(len(missing)), missing)
		}
	} else if a.not && len(missing) == 0 {
		a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
			preS(a.info), a.actual, len(a.actual), a.actual)
		a.addExpectation("to contain %s but they were all present.\n", allN.FormatInt(len(expected)))
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToContainAny asserts that the slice contains any of the values listed.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToContainAny(t Tester, expected ...T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	found := make([]T, 0, len(expected))
	missing := make([]T, 0, len(expected))
	for _, v := range expected {
		if sliceContains(a.actual, v) {
			found = append(found, v)
		} else {
			missing = append(missing, v)
		}
	}

	if !a.not && len(found) == 0 {
		a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
			preS(a.info), a.actual, len(a.actual), a.actual)
		a.addExpectation("to contain %s but none were present.\n",
			anyOfN.FormatInt(len(expected)))
	} else if a.not && len(found) > 0 {
		if len(missing) == 0 {
			a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but %s present.\n",
				anyOfN.FormatInt(len(expected)), theyWereAll.FormatInt(len(expected)))
		} else if len(missing) < len(expected)/2 {
			a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but only %s missing ―――\n  %v\n",
				anyOfN.FormatInt(len(expected)), theseWere.FormatInt(len(missing)), missing)
		} else {
			a.describeActualMulti("Expected%s %T len:%d ―――\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but %s found ―――\n  %v\n",
				anyOfN.FormatInt(len(expected)), theseWere.FormatInt(len(found)), found)
		}
	} else {
		a.passes++
	}

	a.applyAll(t)
}
