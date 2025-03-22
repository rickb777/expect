package expect_test

import (
	"errors"
	"github.com/rickb777/expect"
	"testing"
)

type Seconds32 uint32

func numberTest(e error) (int, error) { return 0, e }

func TestNumberToBe(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).ToBe(1710000000, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBe(1710000001, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be ―――\n  1710000001\n")
}

func TestNumberNotToBe(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).Not().ToBe(1710000001, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBe(1710000000, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be ―――\n  1710000000\n")

	expect.Number(numberTest(errors.New("bang"))).I("data").Not().ToBe(0, c)
	c.shouldHaveCalledFatalf(t, "Expected data not to pass a non-nil error but got parameter 2 (*errors.errorString) ―――\n  bang\n")
}

func TestNumberToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).ToBeGreaterThan(1710000000, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeGreaterThan(1710000000, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be greater than ―――\n  1710000000\n")
}

func TestNumberNotToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).Not().ToBeGreaterThan(1710000001, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThan(1710000000, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― not to be greater than ―――\n  1710000000\n")
}

func TestNumberToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeLessThan(1710000001, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeLessThan(1710000000, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be less than ―――\n  1710000000\n")
}

func TestNumberNotToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThan(1710000000, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThan(1710000001, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be less than ―――\n  1710000001\n")
}

func TestNumberToBeGreaterThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).I("utcTime").ToBeGreaterThanOrEqualTo(1710000000, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeGreaterThanOrEqualTo(1710000001, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be greater than or equal to ―――\n  1710000001\n")
}

func TestNumberNotToBeGreaterThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThanOrEqualTo(1710000001, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThanOrEqualTo(1710000000, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― not to be greater than or equal to ―――\n  1710000000\n")
}

func TestNumberToBeLessThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).ToBeLessThanOrEqualTo(1710000001, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").ToBeLessThanOrEqualTo(1710000000, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― to be less than or equal to ―――\n  1710000000\n")
}

func TestNumberNotToBeLessThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).Not().ToBeLessThanOrEqualTo(1710000000, c)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThanOrEqualTo(1710000001, c)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be less than or equal to ―――\n  1710000001\n")
}
