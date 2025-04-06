package expect

import (
	"fmt"
	"log"
	"math"
	"strings"
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

func (a *assertion) describeActual1line(message string, args ...any) {
	a.actualDescription = fmt.Sprintf(message, args...)
	a.actualSeparator = false
}

func (a *assertion) describeActualMulti(message string, args ...any) {
	a.actualDescription = fmt.Sprintf(message, args...)
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

func (a *assertion) allOtherArgumentsMustBeNil(t Tester) {
	if a != nil {
		if h, ok := t.(helper); ok {
			h.Helper()
		}

		for i, o := range a.otherActual {
			if o != nil {
				v := "value"
				if _, ok := o.(error); ok {
					v = "error"
				}
				t.Fatal(fmt.Sprintf("Expected%s not to pass a non-nil %s but got parameter %d (%T) ―――\n  %v\n", preS(a.info), v, i+2, o, o))
			}
		}
	}
}

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
		return fmt.Sprintf(info.(string), other...)
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

func verbatim(v any) string {
	a := fmt.Sprintf("  %+v\n", v)
	b := fmt.Sprintf("  %#v\n", v)
	if a == b {
		return blank(a)
	}
	return a + b
}

//-------------------------------------------------------------------------------------------------

func findFirstDiff[T comparable](a, b []T) int {
	shortest := min(len(a), len(b))
	for i := 0; i < shortest; i++ {
		ra := a[i]
		rb := b[i]
		if ra != rb {
			return i
		}
	}
	return math.MinInt
}

//-------------------------------------------------------------------------------------------------

func sliceContains[T comparable](list []T, wanted T) bool {
	for _, v := range list {
		if v == wanted {
			return true
		}
	}
	return false
}
