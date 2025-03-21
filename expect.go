package expect

import (
	"cmp"
	"fmt"
	gocmp "github.com/google/go-cmp/cmp"
	"log"
)

type Tester interface {
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

type helper interface {
	Helper()
}

// JustLogIt is a tester that calls log.Fatalf on all test failures.
func JustLogIt() Tester { return &justLogIt{} }

type justLogIt struct{}

func (c *justLogIt) Errorf(message string, args ...any) {
	log.Fatalf(message, args...)
}

func (c *justLogIt) Fatalf(message string, args ...any) {
	log.Fatalf(message, args...)
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
		return info[0].(string)
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
