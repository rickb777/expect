package expect

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Stringy interface {
	~string | []byte | []rune
}

// StringType is used for assertions about strings.
type StringType[S Stringy] struct {
	actual S
	*assertion
	trim int
}

type StringOr[S Stringy] struct {
	main   *StringType[S]
	passes int
}

// String creates a string assertion. Strings must contain valid UTF8 encodings.
//
// It accepts all string subtypes and []byte, []rune.
//
// If more than one argument is passed, all subsequent arguments will be required to be nil/zero.
// This is convenient if you want to make an assertion on a method/function that returns a value and an error,
// a common pattern in Go.
func String[S Stringy](value S, other ...any) *StringType[S] {
	return &StringType[S]{actual: value, assertion: &assertion{otherActual: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a *StringType[S]) Info(info any, other ...any) *StringType[S] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a *StringType[S]) I(info any, other ...any) *StringType[S] {
	return a.Info(info, other...)
}

// Trim shortens the error message for very long strings.
// Trimming is disabled by default.
func (a *StringType[S]) Trim(at int) *StringType[S] {
	a.trim = at
	return a
}

// Not inverts the assertion.
func (a *StringType[S]) Not() *StringType[S] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToBeEmpty asserts that the string has zero length.
// The tester is normally [*testing.T].
func (a *StringType[S]) ToBeEmpty(t Tester) *StringOr[S] {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	return a.toHaveLength(t, 0, "to be empty.")
}

//-------------------------------------------------------------------------------------------------

// ToHaveLength asserts that the string has the expected length.
// The tester is normally [*testing.T].
func (a *StringType[S]) ToHaveLength(t Tester, expected int) *StringOr[S] {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	return a.toHaveLength(t, expected, fmt.Sprintf("to have length %d.", expected))
}

//-------------------------------------------------------------------------------------------------

func (a *StringType[S]) toHaveLength(t Tester, expected int, what string) *StringOr[S] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	actual := len(a.actual)

	if (!a.not && actual != expected) || (a.not && actual == expected) {
		if len(a.actual) > 0 {
			a.describeActualExpectedM("%T len:%d ―――\n  %v\n", a.actual, len(a.actual),
				trim(string(a.actual), a.trim))
			a.addExpectation("%s\n", what)
		} else {
			a.describeActualExpected1("%T len:%d ", a.actual, len(a.actual))
			a.addExpectation("%s\n", what)
		}
		return a.conjunction(t, false)
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToContain asserts that the actual string contains the substring.
// The tester is normally [*testing.T].
func (a *StringType[S]) ToContain(t Tester, substring S) *StringOr[S] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	ac := string(a.actual)
	ex := string(substring)
	match := strings.Contains(ac, ex)

	if (!a.not && !match) || (a.not && match) {
		a.describeActualExpectedM("%T len:%d ―――\n  %s\n", a.actual, len(a.actual), trim(ac, a.trim))
		a.addExpectation("to contain ―――\n  %s\n", trim(ex, a.trim))
		return a.conjunction(t, false)
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToMatch asserts that the actual string matches a regular expression.
// The tester is normally [*testing.T].
func (a *StringType[S]) ToMatch(t Tester, pattern *regexp.Regexp) *StringOr[S] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	ac := string(a.actual)
	match := pattern.MatchString(ac)

	if (!a.not && !match) || (a.not && match) {
		a.describeActualExpectedM("―――\n  %s\n", trim(ac, a.trim))
		a.addExpectation("to match ―――\n  %s\n", pattern)
		return a.conjunction(t, false)
	}

	return a.conjunction(t, true)
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected strings have the same values and types.
// The tester is normally [*testing.T].
func (a *StringType[S]) ToBe(t Tester, expected S) *StringOr[S] {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	return a.toEqual(t, "to be", string(expected))
}

//-------------------------------------------------------------------------------------------------

// ToEqual asserts that the actual and expected strings have the same values and similar types.
// Unlike [StringType.ToBe], the concrete type may differ.
// The tester is normally [*testing.T].
func (a *StringType[S]) ToEqual(t Tester, expected string) *StringOr[S] {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	return a.toEqual(t, "to equal", expected)
}

//-------------------------------------------------------------------------------------------------

func (a *StringType[S]) toEqual(t Tester, what, expected string) *StringOr[S] {
	if a == nil {
		return nil
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustBeNil(t)

	actual := string(a.actual)

	if (!a.not && actual != expected) || (a.not && actual == expected) {
		ac := []rune(actual)
		ex := []rune(expected)
		diff := findFirstDiff(ac, ex)
		pointer := diff + 1
		if diff > 100 || (100 > a.trim && a.trim > 0) {
			rem := 70
			if 70 > a.trim && a.trim > 0 {
				rem = a.trim
				a.trim = 0
			}
			chop := (diff - rem) + 1
			actual = "…" + string(ac[chop:])
			expected = "…" + string(ex[chop:])
			pointer = rem + 1
		}
		a.describeActualExpectedM("―――\n  %s\n", trim(actual, a.trim))
		a.addExpectation("%s\n  %s\n%s",
			arrowMarker(what, pointer),
			trim(expected, a.trim),
			firstDifferenceInfo(diff))
		return a.conjunction(t, false)
	}

	return a.conjunction(t, true)
}

//=================================================================================================

func (a *StringType[S]) conjunction(t Tester, pass bool) *StringOr[S] {
	if pass {
		a.passes++
	}

	if t == nil {
		return &StringOr[S]{main: a, passes: a.passes} // defer evaluation
	}

	if h, ok := t.(helper); ok {
		h.Helper()
	}
	a.applyAll(t)
	return nil
}

//-------------------------------------------------------------------------------------------------

func (or *StringOr[S]) Or() *StringType[S] {
	if or == nil {
		return nil
	}
	return or.main
}

//=================================================================================================

func arrowMarker(label string, i int) string {
	indicator := fmt.Sprintf("%s ―――", label)
	iLength := utf8.RuneCountInString("――― " + indicator)
	if i <= iLength {
		return indicator
	}
	nSpaces := 1 + i - iLength
	return indicator + strings.Repeat(" ", nSpaces) + "↕"
}

func firstDifferenceInfo(i int) string {
	if i < 0 {
		return ""
	}
	return fmt.Sprintf("――― the first difference is at index %d.\n", i)
}

func trim(s string, trim int) string {
	if utf8.RuneCountInString(s) > trim && trim > 0 {
		rs := []rune(s)
		return string(rs[:trim]) + "…"
	}
	return blank(s)
}

func blank(s string) string {
	if len(s) == 0 {
		return `""`
	}
	return s
}
