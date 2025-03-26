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

// Slice creates an assertion for deep value comparison of slices of any type.
//
// This uses [gocmp.Equal] so the manner of comparison can be tweaked using that API - see also [SliceType.Using]
func Slice[T comparable](value []T, other ...any) SliceType[T] {
	return SliceType[T]{actual: value, opts: DefaultOptions(), assertion: assertion{other: other}}
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
// TODO ToBeNil
//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected slices have the same values and types.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToBe(t Tester, expected ...T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	types := gatherTypes(nil, a.actual, expected)
	opts := append(a.opts, allowUnexported(types))

	match := gocmp.Equal(a.actual, expected, opts)

	if (!a.not && !match) || (a.not && match) {
		diff := findFirstDiff(a.actual, expected)
		t.Errorf("Expected%s %T len:%d ―――\n%s――― %sto be len:%d ―――\n%s%s",
			preS(a.info), a.actual, len(a.actual), verbatim(a.actual),
			notS(a.not), len(expected), verbatim(expected), firstDifferenceInfo(diff))
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

//-------------------------------------------------------------------------------------------------

// ToBeEmpty asserts that the slice has zero length.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToBeEmpty(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toHaveLength(t, 0, "to be empty")
}

//-------------------------------------------------------------------------------------------------

// ToHaveLength asserts that the slice has the expected length.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToHaveLength(t Tester, expected int) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toHaveLength(t, expected, fmt.Sprintf("to have length %d", expected))
}

//-------------------------------------------------------------------------------------------------

func (a SliceType[T]) toHaveLength(t Tester, expected int, what string) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	actual := len(a.actual)

	as := ""
	if len(a.actual) > 0 {
		as = fmt.Sprintf("―――\n  %v\n――― ", a.actual)
	}

	if (!a.not && actual != expected) || (a.not && actual == expected) {
		t.Errorf("Expected%s %T len:%d %s%s%s\n",
			preS(a.info), a.actual, len(a.actual), as, notS(a.not), what)
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

//-------------------------------------------------------------------------------------------------

// ToContainAll asserts that the slice contains all of the values listed.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToContainAll(t Tester, expected ...T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

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
			t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― to contain all %d but none were found\n",
				preS(a.info), a.actual, len(a.actual), a.actual, len(expected))
		} else if len(found) < len(expected)/2 {
			t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― to contain all %d but only these %d were found\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual, len(expected), len(found), found)
		} else {
			t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― to contain all %d but these %d were missing\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual, len(expected), len(missing), missing)
		}
	} else if a.not && len(missing) == 0 {
		t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― not to contain all %d but they were all present\n",
			preS(a.info), a.actual, len(a.actual), a.actual, len(expected))
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

//-------------------------------------------------------------------------------------------------

// ToContainAny asserts that the slice contains any of the values listed.
// The tester is normally [*testing.T].
func (a SliceType[T]) ToContainAny(t Tester, expected ...T) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

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
		t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― to contain any of %d but none were present\n",
			preS(a.info), a.actual, len(a.actual), a.actual, len(expected))
	} else if a.not && len(found) > 0 {
		if len(missing) == 0 {
			t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― not to contain any of %d but they were all present\n",
				preS(a.info), a.actual, len(a.actual), a.actual, len(expected))
		} else if len(missing) < len(expected)/2 {
			t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― not to contain any of %d but only these %d were missing\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual,
				len(expected), len(missing), missing)
		} else {
			t.Errorf("Expected%s %T len:%d ―――\n  %v\n――― not to contain any of %d but these %d were found\n  %v\n",
				preS(a.info), a.actual, len(a.actual), a.actual,
				len(expected), len(found), found)
		}
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}
