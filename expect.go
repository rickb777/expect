package expect

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strings"
	"unicode"

	gocmp "github.com/google/go-cmp/cmp"
)

// Tester reports test errors and failures. Notably, [testing.T] implements this interface.
type Tester interface {
	Error(args ...any)
	Fatal(args ...any)
}

type helper interface {
	Helper()
}

// JustLogIt is a tester that calls log.Fatalf on all test errors and failures.
var JustLogIt = SimpleTester(log.Fatal, log.Fatal)

// SimpleTester is a tester that calls errorf on test errors and fatalf on test failures.
func SimpleTester(errorFn, fatalFn func(v ...any)) Tester {
	return &simpleTester{errorFn: errorFn, fatalFn: fatalFn}
}

type simpleTester struct {
	errorFn, fatalFn func(v ...any)
}

func (c *simpleTester) Error(args ...any) {
	c.errorFn(args...)
}

func (c *simpleTester) Fatal(args ...any) {
	c.fatalFn(args...)
}

//-------------------------------------------------------------------------------------------------

type assertion struct {
	info              string
	otherActual       []any
	not               bool
	passes            int
	actualDescription string
	actualSeparator   bool
	moreMessages      []string
}

func (a *assertion) describeActual(message string, args ...any) {
	a.actualDescription = fmt.Sprintf(message, args...)
	a.actualSeparator = false
}

func (a *assertion) describeActualExpected1(message string, args ...any) {
	expected := fmt.Sprintf("Expected%s ", preS(a.info))
	a.actualDescription = expected + fmt.Sprintf(message, args...)
	a.actualSeparator = false
}

func (a *assertion) describeActualExpectedM(message string, args ...any) {
	expected := fmt.Sprintf("Expected%s ", preS(a.info))
	a.actualDescription = expected + fmt.Sprintf(message, args...)
	a.actualSeparator = true
}

func (a *assertion) addExpectation(message string, args ...any) {
	a.moreMessages = append(a.moreMessages, fmt.Sprintf(message, args...))
}

func (a *assertion) applyAll(t Tester) {
	if t != nil && a.passes == 0 {
		if h, ok := t.(helper); ok {
			h.Helper()
		}

		as := ""
		if a.actualSeparator {
			as = "――― "
		}

		if a.not {
			t.Error(a.actualDescription + join(as+"not ", a.moreMessages, "\n――― and not "))
		} else {
			t.Error(a.actualDescription + join(as, a.moreMessages, "\n――― or "))
		}
	}
}

//-------------------------------------------------------------------------------------------------

func join(before string, messages []string, separator string) string {
	if len(messages) == 0 {
		return ""
	}
	return before + strings.Join(messages, separator)
}

//-------------------------------------------------------------------------------------------------

func (a *assertion) allOtherArgumentsMustNotBeError(t Tester) {
	if a != nil {
		if h, ok := t.(helper); ok {
			h.Helper()
		}

		for i, o := range a.otherActual {
			switch o.(type) {
			case error:
				t.Fatal(fmt.Sprintf("Expected%s not to pass a non-nil error but got error parameter %d ―――\n%v\n",
					preS(a.info), i+2, o))
			}
		}
	}
}

//=================================================================================================

const incorrectTestConjunction = "Incorrect test conjunction.\n" +
	"――― Only the last assertion should have a non-nil tester.\n" +
	"――― Use nil for the preceding assertions."

//=================================================================================================

func makeInfo(info any, other ...any) string {
	format, isString := info.(string)

	if len(other) == 0 {
		if isString {
			return format
		} else {
			return fmt.Sprintf("%v", info)
		}
	}

	if isString {
		return fmt.Sprintf(format, other...)
	}

	args := append([]any{info}, other...)
	return fmt.Sprint(args...)
}

func prefix(pfx, s string) string {
	if s == "" {
		return ""
	}
	return pfx + s
}

func preS(s string) string {
	return prefix(" ", s)
}

func notS(not bool) string {
	if not {
		return "not "
	}
	return ""
}

func verbatim1(v any) string {
	return fmt.Sprintf("%+v\n", v)
}

func verbatim2(v any) string {
	t := reflect.TypeOf(v)
	k0 := t.Kind()

	if isBuiltIn(k0) {
		return fmt.Sprintf("%+v\n", v)
	}

	switch k0 {
	case reflect.Map:
		//k1 := t.Key().Kind()
		//k2 := t.Elem().Kind()
		return gocmpDiffAsVerbatimString(v)

	case reflect.Slice:
		k1 := t.Elem().Kind()
		switch k1 {
		case reflect.Uint8:
			return fmt.Sprintf("%+v\n%q\n", v, v)
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return fmt.Sprintf("%+v\n%#v\n", v, v)
		case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map, reflect.Pointer:
			return gocmpDiffAsVerbatimString(v)
		}
	}

	return fmt.Sprintf("%+v\n", v)
}

type private struct{}

func gocmpDiffAsVerbatimString(v any) string {
	// this uses gocmp.Diff just as a value renderer because it is rather good
	s := gocmp.Diff(private{}, v, DefaultOptions())
	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "+") {
			r1 := strings.TrimLeftFunc(line[1:], unicode.IsSpace)
			return strings.TrimRight(r1, ",") + "\n"
		}
	}
	return ""
}

func isBuiltIn(k reflect.Kind) bool {
	return (k >= reflect.Bool && k <= reflect.Complex128) || k == reflect.String
}

//-------------------------------------------------------------------------------------------------

func findFirstRuneDiff(a, b []rune) (diff, line, column int) {
	line = 1
	column = 1
	shortest := min(len(a), len(b))
	for i := 0; i < shortest; i++ {
		if a[i] != b[i] {
			return i, line, column
		}
		column++
		if a[i] == '\n' {
			line++
			column = 1
		}
	}
	return math.MinInt, math.MinInt, math.MinInt
}
