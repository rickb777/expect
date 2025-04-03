package expect_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

type capture struct {
	errorfCalls int
	fatalfCalls int
	message     []string
}

func (c *capture) reset() {
	c.errorfCalls = 0
	c.message = nil
}

func sliceEquals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (c *capture) shouldNotHaveHadAnError(t *testing.T) {
	t.Helper()
	if c.errorfCalls > 0 {
		t.Errorf("failed: Errorf called %d times\n%s", c.errorfCalls, c.message)
	}
	if c.fatalfCalls > 0 {
		t.Errorf("failed: Fatalf called %d times\n%s", c.fatalfCalls, c.message)
	}
	c.reset()
}

func (c *capture) shouldHaveCalledErrorf(t *testing.T, message ...string) {
	t.Helper()
	if c.errorfCalls == 0 {
		t.Errorf("failed to call Errorf (and %d calls to Fatalf)", c.fatalfCalls)
	} else if !sliceEquals(c.message, message) {
		t.Errorf("%s\n―――%d―――", c.message, len(c.message))
	}
	c.reset()
}

func (c *capture) shouldHaveCalledErrorfRE(t *testing.T, message string) {
	t.Helper()
	if c.errorfCalls != 1 {
		t.Errorf("failed to call Errorf once (and %d calls to Fatalf)", c.fatalfCalls)
	} else if !regexp.MustCompile(message).MatchString(c.message[0]) {
		t.Error(strings.Join(c.message, "\n"))
	}
	c.reset()
}

func (c *capture) shouldHaveCalledFatalf(t *testing.T, message ...string) {
	t.Helper()
	if c.fatalfCalls == 0 {
		t.Errorf("failed to call Fatalf (and %d calls to Errorf)", c.errorfCalls)
	} else if !sliceEquals(c.message, message) {
		t.Error(strings.Join(c.message, "\n"))
	}
	c.reset()
}

//func (c *capture) shouldHaveCalledFatalfRE(t *testing.T, message string) {
//	t.Helper()
//	if c.fatalfCalls == 0 {
//		t.Errorf("failed to call Fatalf (and %d calls to Errorf)", c.errorfCalls)
//	} else if !regexp.MustCompile(message).MatchString(c.message) {
//      t.Error(strings.Join(c.message, "\n"))
//	}
//	c.reset()
//}

func (c *capture) Helper() {}

func (c *capture) Errorf(message string, args ...any) {
	c.message = append(c.message, fmt.Sprintf(message, args...))
	c.errorfCalls++
}

func (c *capture) Fatalf(message string, args ...any) {
	c.message = append(c.message, fmt.Sprintf(message, args...))
	c.fatalfCalls++
}
