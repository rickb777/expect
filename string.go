package expect

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Stringy interface {
	~string | []byte | []rune
}

// StringType is used for assertions about strings.
type StringType[S Stringy] struct {
	actual S
	assertion
	trim int
}

// String creates a string assertion. Strings must contain valid UTF8 encodings.
//
// It accepts all string subtypes and []byte, []rune.
//
// If more than one argument is passed, all subsequent arguments will be required to be nil/zero.
// This is convenient if you want to make an assertion on a method/function that returns a value and an error,
// a common pattern in Go.
func String[S Stringy](value S, other ...any) StringType[S] {
	return StringType[S]{actual: value, assertion: assertion{other: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a StringType[S]) Info(info any, other ...any) StringType[S] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a StringType[S]) I(info any, other ...any) StringType[S] {
	return a.Info(info, other...)
}

// Trim shortens the error message for very long strings.
// Trimming is disabled by default.
func (a StringType[S]) Trim(at int) StringType[S] {
	a.trim = at
	return a
}

// Not inverts the assertion.
func (a StringType[S]) Not() StringType[S] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToContain asserts that the actual string contains the substring.
// The tester is normally [*testing.T].
func (a StringType[S]) ToContain(t Tester, substring S) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	ac := fmt.Sprint(a.actual)
	ex := fmt.Sprint(substring)
	match := strings.Contains(ac, ex)

	if (!a.not && !match) || (a.not && match) {
		t.Errorf("Expected%s ―――\n  %s\n――― %sto contain ―――\n  %s\n",
			preS(a.info), trim(ac, a.trim), notS(a.not), trim(ex, a.trim))
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected strings have the same values and types.
// The tester is normally [*testing.T].
func (a StringType[S]) ToBe(t Tester, expected S) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toEqual(t, "to be", string(expected))
}

//-------------------------------------------------------------------------------------------------

// ToEqual asserts that the actual and expected strings have the same values and similar types.
// Unlike [StringType.ToBe], the concrete type may differ.
// The tester is normally [*testing.T].
func (a StringType[S]) ToEqual(t Tester, expected string) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toEqual(t, "to equal", expected)
}

//-------------------------------------------------------------------------------------------------

func (a StringType[S]) toEqual(t Tester, what, expected string) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

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
		t.Errorf("Expected%s ―――\n  %s\n%s\n  %s\n%s",
			preS(a.info),
			trim(actual, a.trim),
			arrowMarker(notS(a.not), what, pointer),
			trim(expected, a.trim),
			firstDifferenceInfo(diff))
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

//=================================================================================================

func arrowMarker(not, label string, i int) string {
	indicator := fmt.Sprintf("――― %s%s ―――", not, label)
	iLength := utf8.RuneCountInString(indicator)
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
	return fmt.Sprintf("――― the first difference is at index %d\n", i)
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
