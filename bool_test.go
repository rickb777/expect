package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

func TestBoolToBeTrue(t *testing.T) {
	c := &capture{}

	expect.Bool(c, true, "data").ToBeTrue()
	c.shouldNotHaveHadAnError(t)

	expect.Bool(c, false, "data").ToBeTrue()
	c.shouldHaveCalledErrorf(t, "Expected data to be true\n")
}

func TestBoolNotToBeTrue(t *testing.T) {
	c := &capture{}

	expect.Bool(c, false, "data").Not().ToBeTrue()
	c.shouldNotHaveHadAnError(t)

	expect.Bool(c, true, "data").Not().ToBeTrue()
	c.shouldHaveCalledErrorf(t, "Expected data not to be true\n")
}

func TestBoolToBeFalse(t *testing.T) {
	c := &capture{}

	expect.Bool(c, false, "data").ToBeFalse()
	c.shouldNotHaveHadAnError(t)

	expect.Bool(c, true, "data").ToBeFalse()
	c.shouldHaveCalledErrorf(t, "Expected data to be false\n")
}

func TestBoolNotToBeFalse(t *testing.T) {
	c := &capture{}

	expect.Bool(c, true, "data").Not().ToBeFalse()
	c.shouldNotHaveHadAnError(t)

	expect.Bool(c, false, "data").Not().ToBeFalse()
	c.shouldHaveCalledErrorf(t, "Expected data not to be false\n")
}
