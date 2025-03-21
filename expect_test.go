package expect_test

import (
	"fmt"
)

type capture struct {
	called  bool
	message string
}

func (c *capture) reset() {
	c.called = false
	c.message = ""
}

func (c *capture) Helper() {}

func (c *capture) Errorf(message string, args ...any) {
	c.called = true
	c.message = fmt.Sprintf(message, args...)
}

func (c *capture) Fatalf(message string, args ...any) {
	c.called = true
	c.message = fmt.Sprintf(message, args...)
}
