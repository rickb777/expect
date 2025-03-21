package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

type Seconds32 uint32

func TestNumberToBe(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime).ToBe(1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBe(1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be ―――\n  1710000001\n")
}

func TestNumberNotToBe(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime).Not().ToBe(1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").Not().ToBe(1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be ―――\n  1710000000\n")
}

func TestNumberToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime).ToBeGreaterThan(1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeGreaterThan(1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be greater than ―――\n  1710000000\n")
}

func TestNumberNotToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime).Not().ToBeGreaterThan(1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(c, utcTime, "utcTime").Not().ToBeGreaterThan(1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― not to be greater than ―――\n  1710000000\n")
}

func TestNumberToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeLessThan(1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeLessThan(1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be less than ―――\n  1710000000\n")
}

func TestNumberNotToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime, "utcTime").Not().ToBeLessThan(1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").Not().ToBeLessThan(1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be less than ―――\n  1710000001\n")
}

func TestNumberToBeGreaterThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime, "utcTime").ToBeGreaterThanOrEqualTo(1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeGreaterThanOrEqualTo(1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be greater than or equal to ―――\n  1710000001\n")
}

func TestNumberNotToBeGreaterThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime, "utcTime").Not().ToBeGreaterThanOrEqualTo(1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(c, utcTime, "utcTime").Not().ToBeGreaterThanOrEqualTo(1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― not to be greater than or equal to ―――\n  1710000000\n")
}

func TestNumberToBeLessThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime).ToBeLessThanOrEqualTo(1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(c, utcTime, "utcTime").ToBeLessThanOrEqualTo(1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― to be less than or equal to ―――\n  1710000000\n")
}

func TestNumberNotToBeLessThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime).Not().ToBeLessThanOrEqualTo(1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").Not().ToBeLessThanOrEqualTo(1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be less than or equal to ―――\n  1710000001\n")
}
