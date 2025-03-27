package expect

import (
	"fmt"
	gocmp "github.com/google/go-cmp/cmp"
	"slices"
	"strings"
)

// MapType is used for assertions about maps.
type MapType[K comparable, V any] struct {
	opts   gocmp.Options
	actual map[K]V
	assertion
}

// Map creates an assertion for deep value comparison of maps of any type.
//
// This uses [gocmp.Equal] so the manner of comparison can be tweaked using that API - see also [MapType.Using]
func Map[K comparable, V any](value map[K]V, other ...any) MapType[K, V] {
	return MapType[K, V]{actual: value, opts: DefaultOptions(), assertion: assertion{other: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number. If this
// is a format string, more parameters can follow and will be formatted accordingly (see [fmt.Sprintf]).
func (a MapType[K, V]) Info(info any, other ...any) MapType[K, V] {
	a.info = makeInfo(info, other...)
	return a
}

// I is a synonym for [Info].
func (a MapType[K, V]) I(info any, other ...any) MapType[K, V] {
	return a.Info(info, other...)
}

// Using replaces the default comparison options with those specified here.
// You can also set [DefaultOptions] instead.
func (a MapType[K, V]) Using(opt ...gocmp.Option) MapType[K, V] {
	a.opts = opt
	return a
}

// Not inverts the assertion.
func (a MapType[K, V]) Not() MapType[K, V] {
	a.not = !a.not
	return a
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected maps have the same values and types.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToBe(t Tester, expected map[K]V) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	opts := append(a.opts, allowUnexported(gatherTypes(nil, a.actual, expected)))

	match := gocmp.Equal(a.actual, expected, opts)

	if (!a.not && !match) || (a.not && match) {
		t.Errorf("Expected%s %T len:%d ―――\n  %+v\n――― %sto be len:%d ―――\n  %+v\n",
			preS(a.info), a.actual, len(a.actual), a.actual,
			notS(a.not), len(expected), expected)
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

//-------------------------------------------------------------------------------------------------

// ToContain asserts that the map contains a particular key. If present, the expected value must also match.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToContain(t Tester, expectedKey K, expectedValue ...V) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	value, present := a.actual[expectedKey]

	types := gatherTypes(nil, a.actual, expectedKey)

	expectedKeyValue := ""
	evi := ""
	expectedKeyS := quotedString(expectedKey)

	if len(expectedValue) > 0 {
		expectedKeyValue = fmt.Sprintf("  %s: %+v\n", expectedKeyS, expectedValue[0])
		if a.not {
			evi = " but it should not match"
		} else {
			evi = " and it should match"
		}

		types = gatherTypes(types, expectedValue[0])
	}

	opts := append(a.opts, allowUnexported(types))

	if !a.not {
		if !present {
			t.Errorf("Expected%s %T len:%d to contain %s; keys are ―――\n  [%s]\n",
				preS(a.info), a.actual, len(a.actual),
				expectedKeyS, strings.Join(toString(keys(a.actual)), ", "))

		} else if len(expectedValue) > 0 && !gocmp.Equal(value, expectedValue[0], opts) {
			t.Errorf("Expected%s %T len:%d ―――\n  %s: %+v\n――― to contain %s%s ―――\n%s",
				preS(a.info), a.actual, len(a.actual), expectedKeyS, value,
				expectedKeyS, evi, expectedKeyValue)
		}

	} else if present {
		if len(expectedValue) > 0 && !gocmp.Equal(value, expectedValue[0], opts) {
			t.Errorf("Expected%s %T len:%d contains ―――\n"+
				"  %s: %+v\n"+
				"――― but should contain ―――\n%s",
				preS(a.info), a.actual, len(a.actual),
				expectedKeyS, value, expectedKeyValue)

		} else {
			t.Errorf("Expected%s %T len:%d not to contain %s; keys are ―――\n  [%s]\n",
				preS(a.info), a.actual, len(a.actual),
				expectedKeyS, strings.Join(toString(keys(a.actual)), ", "))
		}
	}

	allOtherArgumentsMustBeNil(t, a.info, a.other...)
}

func quotedString(v any) string {
	switch s := v.(type) {
	case string:
		return fmt.Sprintf("%q", s)
	default:
		return fmt.Sprintf("%v", v)
	}
}

//-------------------------------------------------------------------------------------------------

// ToBeEmpty asserts that the map has zero length.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToBeEmpty(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toHaveLength(t, 0, "to be empty")
}

//-------------------------------------------------------------------------------------------------

// ToHaveSize is a synonym for ToHaveLength.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToHaveSize(t Tester, expected int) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}
	a.toHaveLength(t, expected, fmt.Sprintf("to have size %d", expected))
}

// ToHaveLength asserts that the map has the expected length.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToHaveLength(t Tester, expected int) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}
	a.toHaveLength(t, expected, fmt.Sprintf("to have length %d", expected))
}

// ToHaveLength asserts that the map has the expected length.
// The tester is normally [*testing.T].
func (a MapType[K, V]) toHaveLength(t Tester, expected int, what string) {
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

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func toString[K comparable](keys []K) []string {
	ss := make([]string, 0, len(keys))
	for _, k := range keys {
		ss = append(ss, fmt.Sprintf("%v", k))
	}
	slices.Sort(ss)
	return ss
}
