package expect_test

import (
	"errors"
	"github.com/rickb777/expect"
	"testing"
)

type Seconds32 uint32

func numberTest(e error) (int, error) { return 0, e }

func TestNumberOr(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001

	//----- match early -----

	expect.Number(utcTime).ToBe(nil, 1710000001).Or().ToBe(c, 1710000002)
	c.shouldNotHaveHadAnError(t)

	//----- match late -----

	expect.Number(utcTime).ToBe(nil, 1710000000).Or().ToBe(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	//----- mis-match -----

	expect.Number(utcTime).ToBe(c, 1710000001).Or().ToBe(c, 1710000002)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberToBe(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).ToBe(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBe(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be ―――\n  1710000001\n")
}

func TestNumberNotToBe(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).Not().ToBe(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBe(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be ―――\n  1710000000\n")

	expect.Number(numberTest(errors.New("bang"))).I("data").Not().ToBe(c, 0)
	c.shouldHaveCalledFatalf(t,
		"Expected data not to pass a non-nil error but got error parameter 2 ―――\n  bang\n",
		"Expected data int ―――\n  0\n――― not to be ―――\n  0\n",
	)
}

func TestNumberToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).ToBeGreaterThan(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeGreaterThan(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be greater than ―――\n  1710000000\n")
}

func TestNumberNotToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).Not().ToBeGreaterThan(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThan(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― not to be greater than ―――\n  1710000000\n")
}

func TestNumberToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeLessThan(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeLessThan(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be less than ―――\n  1710000000\n")
}

func TestNumberNotToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThan(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThan(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be less than ―――\n  1710000001\n")
}

func TestNumberToBeGreaterThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).I("utcTime").ToBeGreaterThanOrEqual(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeGreaterThanOrEqual(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― to be greater than or equal to ―――\n  1710000001\n")
}

func TestNumberNotToBeGreaterThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThanOrEqual(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThanOrEqual(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― not to be greater than or equal to ―――\n  1710000000\n")
}

func TestNumberToBeLessThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).ToBeLessThanOrEqual(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").ToBeLessThanOrEqual(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000001\n――― to be less than or equal to ―――\n  1710000000\n")
}

func TestNumberNotToBeLessThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).Not().ToBeLessThanOrEqual(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThanOrEqual(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n  1710000000\n――― not to be less than or equal to ―――\n  1710000001\n")
}

func TestNumberToBeBetween(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).ToBeBetween(c, 1710000000, 1710000002)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeBetween(c, 1710000001, 1710000002)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n"+
		"  1710000000\n"+
		"――― to be between ―――\n"+
		"  1710000001 … 1710000002 (exclusive)\n")

	expect.Number(utcTime).I("utcTime").ToBeBetween(c, 1710000002, 1710000000)
	c.shouldHaveCalledErrorf(t, "Impossible test utcTime expect_test.Seconds32: minimum 1710000002 >= maximum 1710000000.\n")
}

func TestNumberNotToBeBetween(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).Not().ToBeBetween(c, 1710000002, 1710000003)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeBetween(c, 1710000000, 1710000002)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n"+
		"  1710000001\n"+
		"――― not to be between ―――\n"+
		"  1710000000 … 1710000002 (exclusive)\n")

	expect.Number(utcTime).I("utcTime").Not().ToBeBetween(c, 1710000002, 1710000000)
	c.shouldHaveCalledErrorf(t, "Impossible test utcTime expect_test.Seconds32: minimum 1710000002 >= maximum 1710000000.\n")
}

func TestNumberToBeBetweenOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).ToBeBetweenOrEqual(c, 1710000000, 1710000002)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).ToBeBetweenOrEqual(c, 1710000000, 1710000002)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000002
	expect.Number(utcTime).ToBeBetweenOrEqual(c, 1710000000, 1710000002)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeBetweenOrEqual(c, 1710000001, 1710000002)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n"+
		"  1710000000\n"+
		"――― to be between ―――\n"+
		"  1710000001 … 1710000002 (inclusive)\n")

	expect.Number(utcTime).I("utcTime").ToBeBetweenOrEqual(c, 1710000002, 1710000000)
	c.shouldHaveCalledErrorf(t, "Impossible test utcTime expect_test.Seconds32: minimum 1710000002 > maximum 1710000000.\n")
}

func TestNumberNotToBeBetweenOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).Not().ToBeBetweenOrEqual(c, 1710000002, 1710000003)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeBetweenOrEqual(c, 1710000000, 1710000002)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n"+
		"  1710000001\n"+
		"――― not to be between ―――\n"+
		"  1710000000 … 1710000002 (inclusive)\n")

	expect.Number(utcTime).I("utcTime").Not().ToBeBetweenOrEqual(c, 1710000002, 1710000000)
	c.shouldHaveCalledErrorf(t, "Impossible test utcTime expect_test.Seconds32: minimum 1710000002 > maximum 1710000000.\n")
}

func ExampleOrderedType_ToBe() {
	var t *testing.T

	// Number matching can use any size of int or uint or float, or subtype of any of them.
	// This example allows either of the two expected values.
	v := 123
	expect.Number(v).ToBe(nil, 123).Or().ToBe(t, 125)

	// The `Info` method can be helpful when testing inside a loop, for example.
	var i int // some loop counter
	expect.Number(v).Info("loop %d", i).Not().ToBe(t, 321)
}

func ExampleOrderedType_ToBeBetween() {
	var t *testing.T

	// number matching can use any size of int or uint or float, or subtype of any of them
	v := 123
	expect.Number(v).ToBeBetween(t, 100, 200)

	var i int // some loop counter
	expect.Number(v).Info("loop %d", i).Not().ToBeBetween(t, 1, 10)
}

func ExampleOrderedType_ToBeGreaterThan() {
	var t *testing.T

	// number matching can use any size of int or uint or float, or subtype of any of them
	v := 123
	expect.Number(v).ToBeGreaterThan(t, 100)

	var i int // some loop counter
	expect.Number(v).Info("loop %d", i).Not().ToBeGreaterThan(t, 200)
}
