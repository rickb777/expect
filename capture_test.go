package expect_test

import (
	"fmt"
	"testing"
)

type capture struct {
	errorfCalls int
	fatalfCalls int
	message     string
}

func (c *capture) reset() {
	c.errorfCalls = 0
	c.message = ""
}

func (c *capture) shouldNotHaveHadAnError(t *testing.T) {
	t.Helper()
	if c.errorfCalls > 0 {
		t.Errorf("failed: Errorf called %d times", c.errorfCalls)
	}
	if c.fatalfCalls > 0 {
		t.Errorf("failed: Fatalf called %d times", c.fatalfCalls)
	}
	c.reset()
}

func (c *capture) shouldHaveCalledErrorf(t *testing.T, message string) {
	t.Helper()
	if c.errorfCalls == 0 {
		t.Errorf("failed to call Errorf (and %d calls to Fatalf)", c.fatalfCalls)
	} else if c.message != message {
		t.Error(c.message)
	}
	c.reset()
}

func (c *capture) shouldHaveCalledFatalf(t *testing.T, message string) {
	t.Helper()
	if c.fatalfCalls == 0 {
		t.Errorf("failed to call Fatalf (and %d calls to Errorf)", c.errorfCalls)
	} else if c.message != message {
		t.Error(c.message)
	}
	c.reset()
}

func (c *capture) Helper() {}

func (c *capture) Errorf(message string, args ...any) {
	c.errorfCalls++
	c.message = fmt.Sprintf(message, args...)
}

func (c *capture) Fatalf(message string, args ...any) {
	c.fatalfCalls++
	c.message = fmt.Sprintf(message, args...)
}
