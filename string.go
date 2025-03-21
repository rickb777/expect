package expect

import (
	"fmt"
	"strings"
)

// String creates a string assertion. It accepts all string subtypes.
//
// If present, the third parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func String[T Stringy](t Tester, value T, info ...any) StringyType[T] {
	return StringyType[T]{t: t, actual: value, info: makeInfo(info...)}
}

// Not inverts the assertion.
func (a StringyType[T]) Not() StringyType[T] {
	a.not = !a.not
	return a
}

// ToBe asserts that the actual and expected strings have the same values and types.
func (a StringyType[T]) ToBe(expected T) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	ac := fmt.Sprint(a.actual)
	ex := fmt.Sprint(expected)
	match := ac == ex
	if (!a.not && !match) || (a.not && match) {
		a.t.Errorf("Expected%s ...\n  %s\n... %sto be ...\n  %s\n", preS(a.info), trim100(ac), notS(a.not), trim100(ex))
	}
}

// ToEqual asserts that the actual and expected strings have the same values.
// Unlike [StringyType.ToBe], the concrete type may differ.
func (a StringyType[T]) ToEqual(expected string) {
	if h, ok := a.t.(helper); ok {
		h.Helper()
	}

	ac := fmt.Sprint(a.actual)
	match := ac == expected
	if (!a.not && !match) || (a.not && match) {
		a.t.Errorf("Expected%s ...\n  %s\n... %sto equal ...\n  %s\n", preS(a.info), trim100(ac), notS(a.not), trim100(expected))
	}
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
		a.t.Errorf("Expected%s ...\n  %s\n... %sto contain ...\n  %s\n", preS(a.info), trim100(ac), notS(a.not), trim100(ex))
	}
}

//=================================================================================================

func trim100(s string) string {
	if len(s) > 100 {
		return s[:100] + "..."
	}
	return blank(s)
}

func blank(s string) string {
	if len(s) == 0 {
		return `""`
	}
	return s
}
