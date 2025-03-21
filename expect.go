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

type AnyType struct {
	t      Tester
	info   string
	opts   gocmp.Options
	actual any
	not    bool
}

type BoolType struct {
	t      Tester
	info   string
	actual bool
	not    bool
}

type OrderedType[T cmp.Ordered] struct {
	t      Tester
	info   string
	actual T
	not    bool
}

type Stringy interface {
	~string
}

type StringyType[T Stringy] struct {
	t      Tester
	info   string
	actual T
	not    bool
	trim   int
}

type ErrorType struct {
	t      Tester
	info   string
	actual error
	not    bool
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
