package expect

import (
	"fmt"
	gocmp "github.com/google/go-cmp/cmp"
	"github.com/rickb777/plural"
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
	return MapType[K, V]{actual: value, opts: DefaultOptions(), assertion: assertion{otherActual: other}}
}

// Info adds a description of the assertion to be included in any error message.
// The first parameter should be some information such as a string or a number or even a struct.
// If info is a format string, more parameters can follow and will be formatted accordingly (see
// [fmt.Sprintf]).
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

// ToBeNil asserts that the actual value is nil / is not nil.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToBeNil(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if !a.not && !isNilish(a.actual) {
		a.describeActualExpected1("%T len:%d ―――\n%s――― to be nil\n",
			a.actual, len(a.actual), verbatim(a.actual))
	} else if a.not && isNilish(a.actual) {
		a.describeActualExpected1("%T not to be nil\n", a.actual)
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBe asserts that the actual and expected maps have the same values and types.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToBe(t Tester, expected map[K]V) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	opts := append(a.opts, allowUnexported(gatherTypes(nil, a.actual, expected)))

	match := gocmp.Equal(a.actual, expected, opts)

	if (!a.not && !match) || (a.not && match) {
		a.describeActualExpected1("%T len:%d ―――\n%+v\n――― %sto be len:%d ―――\n%+v\n",
			a.actual, len(a.actual), a.actual,
			notS(a.not), len(expected), expected)
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToBeEmpty asserts that the map has zero length.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToBeEmpty(t Tester) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.toHaveLength(t, 0, "to be empty.")
}

//-------------------------------------------------------------------------------------------------

// ToHaveSize is a synonym for ToHaveLength.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToHaveSize(t Tester, expected int) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}
	a.toHaveLength(t, expected, fmt.Sprintf("to have size %d.", expected))
}

// ToHaveLength asserts that the map has the expected length.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToHaveLength(t Tester, expected int) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}
	a.toHaveLength(t, expected, fmt.Sprintf("to have length %d.", expected))
}

// ToHaveLength asserts that the map has the expected length.
// The tester is normally [*testing.T].
func (a MapType[K, V]) toHaveLength(t Tester, expected int, what string) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	actual := len(a.actual)

	if (!a.not && actual != expected) || (a.not && actual == expected) {
		if len(a.actual) > 0 {
			a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
			a.addExpectation("%s\n", what)
		} else {
			a.describeActualExpected1("%T len:%d ", a.actual, len(a.actual))
			a.addExpectation("%s\n", what)
		}
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToContain asserts that the map contains a particular key. If present, the expected value must also match.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToContain(t Tester, expectedKey K, expectedValue ...V) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	value, present := a.actual[expectedKey]

	types := gatherTypes(nil, a.actual, expectedKey)

	evi := ""
	expectedKeyS := quotedString(expectedKey)

	if len(expectedValue) > 0 {
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
			a.describeActualExpected1("%T len:%d to contain %s; keys are ―――\n[%s]\n",
				a.actual, len(a.actual),
				expectedKeyS, strings.Join(toString(keys(a.actual)), ", "))
		} else if len(expectedValue) > 0 && !gocmp.Equal(value, expectedValue[0], opts) {
			a.describeActualExpectedM("%T len:%d ―――\n%s: %+v\n", a.actual, len(a.actual), expectedKeyS, value)
			a.addExpectation("to contain %s%s ―――\n%s: %+v\n", expectedKeyS, evi, expectedKeyS, expectedValue[0])
		} else {
			a.passes++
		}

	} else if present {
		if len(expectedValue) > 0 && gocmp.Equal(value, expectedValue[0], opts) {
			a.describeActualExpected1("%T len:%d not to contain ―――\n%s: %+v\n――― but it does.\n",
				a.actual, len(a.actual), expectedKeyS, value)

		} else {
			a.describeActualExpected1("%T len:%d not to contain %s; keys are ―――\n[%s]\n",
				a.actual, len(a.actual),
				expectedKeyS, strings.Join(toString(keys(a.actual)), ", "))
		}

	} else {
		a.passes++
	}

	a.applyAll(t)
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

var (
	theseWere   = plural.FromOne("this was", "these %d were")
	theyWereAll = plural.FromOne("it was", "they were all")
	allN        = plural.FromOne("it", "both", "all %d")
	anyOfN      = plural.FromOne("it", "both", "any of %d")
)

// ToContainAll asserts that the map contains all the expected keys.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToContainAll(t Tester, expectedKey ...K) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if len(expectedKey) == 1 {
		a.ToContain(t, expectedKey[0])
		return
	}

	found := make([]K, 0, len(expectedKey))
	missing := make([]K, 0, len(expectedKey))
	for _, k := range expectedKey {
		_, present := a.actual[k]
		if present {
			found = append(found, k)
		} else {
			missing = append(missing, k)
		}
	}

	if !a.not && len(missing) > 0 {
		if len(found) == 0 {
			a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but none were found.\n", allN.FormatInt(len(expectedKey)))
		} else if len(found) < len(expectedKey)/2 {
			a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but only %s found ―――\n%v\n",
				allN.FormatInt(len(expectedKey)), theseWere.FormatInt(len(found)), found)
		} else {
			a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but %s missing ―――\n%v\n",
				allN.FormatInt(len(expectedKey)), theseWere.FormatInt(len(missing)), missing)
		}
	} else if a.not && len(missing) == 0 {
		a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
		a.addExpectation("to contain %s but %s present.\n",
			allN.FormatInt(len(expectedKey)), theyWereAll.FormatInt(len(expectedKey)))
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

// ToContainAny asserts that the map contains any the expected keys.
// The tester is normally [*testing.T].
func (a MapType[K, V]) ToContainAny(t Tester, expectedKey ...K) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	a.allOtherArgumentsMustNotBeError(t)

	if len(expectedKey) == 1 {
		a.ToContain(t, expectedKey[0])
		return
	}

	found := make([]K, 0, len(expectedKey))
	missing := make([]K, 0, len(expectedKey))
	for _, k := range expectedKey {
		_, present := a.actual[k]
		if present {
			found = append(found, k)
		} else {
			missing = append(missing, k)
		}
	}

	if !a.not && len(found) == 0 {
		a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
		a.addExpectation("to contain %s but none were present.\n", anyOfN.FormatInt(len(expectedKey)))
	} else if a.not && len(found) > 0 {
		if len(missing) == 0 {
			a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but %s present.\n",
				anyOfN.FormatInt(len(expectedKey)), theyWereAll.FormatInt(len(expectedKey)))
		} else if len(missing) < len(expectedKey)/2 {
			a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but only %s missing ―――\n%v\n",
				anyOfN.FormatInt(len(expectedKey)), theseWere.FormatInt(len(missing)), missing)
		} else {
			a.describeActualExpectedM("%T len:%d ―――\n%v\n", a.actual, len(a.actual), a.actual)
			a.addExpectation("to contain %s but %s found ―――\n%v\n",
				anyOfN.FormatInt(len(expectedKey)), theseWere.FormatInt(len(found)), found)
		}
	} else {
		a.passes++
	}

	a.applyAll(t)
}

//-------------------------------------------------------------------------------------------------

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

//-------------------------------------------------------------------------------------------------

func toString[K comparable](keys []K) []string {
	ss := make([]string, 0, len(keys))
	for _, k := range keys {
		ss = append(ss, fmt.Sprintf("%v", k))
	}
	slices.Sort(ss)
	return ss
}
