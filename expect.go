package expect

import (
	"cmp"
	"fmt"
	gocmp "github.com/google/go-cmp/cmp"
	"log"
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

type assertion struct {
	info  string
	other []any
	not   bool
}

// AnyType is used for equality assertions for any type.
type AnyType struct {
	opts   gocmp.Options
	actual any
	assertion
}

// BoolType is used for assertions about bools.
type BoolType[B ~bool] struct {
	actual B
	assertion
}

// OrderedType is used for assertions about numbers and other ordered types.
type OrderedType[O cmp.Ordered] struct {
	actual O
	assertion
}

type Stringy interface {
	~string | []byte | []rune
}

// StringType is used for assertions about strings.
type StringType[S Stringy] struct {
	actual S
	assertion
	trim int
}

// ErrorType is used for assertions about errors.
type ErrorType struct {
	actual error
	assertion
}

//=================================================================================================

func makeInfo(info ...any) string {
	if len(info) > 1 {
		return fmt.Sprintf(info[0].(string), info[1:]...)
	} else if len(info) > 0 {
		return fmt.Sprintf("%v", info[0])
	}
	return ""
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

func allOtherArgumentsMustBeNil(t Tester, info string, other ...any) {
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
