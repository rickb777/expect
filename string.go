package expect

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// String creates a string assertion. It accepts all string subtypes.
//
// If present, the third parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func String[T Stringy](t Tester, value T, info ...any) StringyType[T] {
	return StringyType[T]{t: t, actual: value, info: makeInfo(info...)}
}

// Trim shortens the error message for very long strings.
func (a StringyType[T]) Trim(at int) StringyType[T] {
	a.trim = at
	return a
}

// Not inverts the assertion.
func (a StringyType[T]) Not() StringyType[T] {
	a.not = !a.not
	return a
}

// ToContain asserts that the actual string contains the substring.
func (a StringyType[T]) ToContain(substring T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	ac := fmt.Sprint(a.actual)
	ex := fmt.Sprint(substring)
	match := strings.Contains(ac, ex)

	if (!a.not && !match) || (a.not && match) {
		a.t.Errorf("Expected%s ―――\n  %s\n――― %sto contain ―――\n  %s\n", preS(a.info), trim(ac, a.trim), notS(a.not), trim(ex, a.trim))
	}
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected strings have the same values and types.
func (a StringyType[T]) ToBe(expected T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	a.toEqual("to be", fmt.Sprint(expected))
}

// ToEqual asserts that the actual and expected strings have the same values.
// Unlike [StringyType.ToBe], the concrete type may differ.
func (a StringyType[T]) ToEqual(expected string) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	a.toEqual("to equal", expected)
}

func (a StringyType[T]) toEqual(what, expected string) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	actual := fmt.Sprint(a.actual)

	if (!a.not && actual != expected) || (a.not && actual == expected) {
		ac := []rune(actual)
		ex := []rune(expected)
		diff := findFirstDiff(ac, ex)
		pointer := diff
		if diff > 100 || (100 > a.trim && a.trim > 0) {
			rem := 70
			if 70 > a.trim && a.trim > 0 {
				rem = a.trim
				a.trim = 0
			}
			chop := diff - rem
			actual = "…" + string(ac[chop:])
			expected = "…" + string(ex[chop:])
			pointer = rem + 1
		}
		a.t.Errorf("Expected%s ―――\n  %s\n%s\n  %s\n%s", preS(a.info),
			trim(actual, a.trim), arrowMarker(notS(a.not), what, pointer), trim(expected, a.trim), firstDiff(diff))
	}
}

//=================================================================================================

func findFirstDiff(a, b []rune) int {
	shortest := min(len(a), len(b))
	for i := 0; i < shortest; i++ {
		ra := a[i]
		rb := b[i]
		if ra != rb {
			return i + 1
		}
	}
	return -1
}

func arrowMarker(not, label string, i int) string {
	indicator := fmt.Sprintf("――― %s%s ―――", not, label)
	iLength := utf8.RuneCountInString(indicator)
	if i <= iLength {
		return indicator
	}
	nSpaces := 1 + i - iLength
	return indicator + strings.Repeat(" ", nSpaces) + "↕"
}

func firstDiff(i int) string {
	if i < 0 {
		return ""
	}
	return fmt.Sprintf("――― the first difference is at character %d\n", i)
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
