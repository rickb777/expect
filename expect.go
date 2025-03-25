package expect

import (
	"fmt"
	"log"
	"math"
)

// Tester reports test errors and failures. Notably, [testing.T] implements this interface.
type Tester interface {
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

type helper interface {
	Helper()
}

// JustLogIt is a tester that calls log.Fatalf on all test errors and failures.
var JustLogIt = SimpleTester(log.Fatalf, log.Fatalf)

// SimpleTester is a tester that calls errorf on test errors and fatalf on test failures.
func SimpleTester(errorf, fatalf func(format string, v ...any)) Tester {
	return &simpleTester{errorf: errorf, fatalf: fatalf}
}

type simpleTester struct {
	errorf, fatalf func(format string, v ...any)
}

func (c *simpleTester) Errorf(message string, args ...any) {
	c.errorf(message, args...)
}

func (c *simpleTester) Fatalf(message string, args ...any) {
	c.fatalf(message, args...)
}

//-------------------------------------------------------------------------------------------------

func deepEqual() {

}

//-------------------------------------------------------------------------------------------------

type assertion struct {
	info  string
	other []any
	not   bool
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

//-------------------------------------------------------------------------------------------------

func allOtherArgumentsMustBeNil(t Tester, info string, other ...any) {
	if h, ok := t.(helper); ok {
		h.Helper()
	}

	for i, o := range other {
		if o != nil {
			v := "value"
			if _, ok := o.(error); ok {
				v = "error"
			}
			t.Fatalf("Expected%s not to pass a non-nil %s but got parameter %d (%T) ―――\n  %v\n", preS(info), v, i+2, o, o)
		}
	}
}
