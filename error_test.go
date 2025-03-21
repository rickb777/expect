package expect_test

import (
	"errors"
	"github.com/rickb777/expect"
	"testing"
)

var e1 = errors.New("something bad happened")

func TestErrorToBeNil(t *testing.T) {
	c := &capture{}

	expect.Error(c, nil, "xyz").ToBeNil()
	c.shouldNotHaveHadAnError(t)

	expect.Error(c, e1, "xyz").ToBeNil()
	c.shouldHaveCalledFatalf(t, "Expected xyz error ―――\n  something bad happened\n――― not to have occurred.\n")
}

func TestErrorToHaveOccurred(t *testing.T) {
	c := &capture{}

	expect.Error(c, e1, "xyz").ToHaveOccurred()
	c.shouldNotHaveHadAnError(t)

	expect.Error(c, nil, "xyz").ToHaveOccurred()
	c.shouldHaveCalledErrorf(t, "Expected xyz error to have occurred.\n")
}

func TestErrorNotToBeNil(t *testing.T) {
	c := &capture{}

	expect.Error(c, e1, "xyz").Not().ToBeNil()
	c.shouldNotHaveHadAnError(t)

	expect.Error(c, nil, "xyz").Not().ToBeNil()
	c.shouldHaveCalledErrorf(t, "Expected xyz error to have occurred.\n")
}

func TestErrorNotToHaveOccurred(t *testing.T) {
	c := &capture{}

	expect.Error(c, nil, "xyz").Not().ToHaveOccurred()
	c.shouldNotHaveHadAnError(t)

	expect.Error(c, e1, "xyz").Not().ToHaveOccurred()
	c.shouldHaveCalledFatalf(t, "Expected xyz error ―――\n  something bad happened\n――― not to have occurred.\n")
}

func TestErrorToContain(t *testing.T) {
	c := &capture{}

	expect.Error(c, e1, "xyz").ToContain("something bad happened")
	c.shouldNotHaveHadAnError(t)

	expect.Error(c, e1, "xyz").ToContain("missing")
	c.shouldHaveCalledErrorf(t, "Expected xyz error ―――\n  something bad happened\n――― to contain ―――\n  missing\n")
}
