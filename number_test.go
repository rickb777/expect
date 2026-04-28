package expect_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rickb777/expect"
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
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― to be ―――\n1710000001\n")

	expect.Number(1.1).ToBe(c, 1.1)
	c.shouldNotHaveHadAnError(t)

	expect.ApproximateFloatFraction = 1e-6
	expect.Number(1.1).ToBe(c, 1.100001) // almost the same
	c.shouldNotHaveHadAnError(t)

	expect.Number(1.1).Using(cmpopts.EquateApprox(1e-7, 0)).ToBe(c, 1.100001) // almost the same
	c.shouldHaveCalledErrorf(t, "Expected float64 ―――\n1.1\n――― to be ―――\n1.100001\n")

	expect.Number(1.1).ToBe(c, 1.10001)
	c.shouldHaveCalledErrorf(t, "Expected float64 ―――\n1.1\n――― to be ―――\n1.10001\n")
}

func TestNumberToEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).ToEqual(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	expect.Number(uint64(1)).ToEqual(c, float32(1.0))
	c.shouldNotHaveHadAnError(t)

	expect.Number("aardvark").ToEqual(c, 1)
	c.shouldHaveCalledErrorf(t, "Expected string ―――\naardvark\n――― to equal ―――\n1\n")

	expect.Number("aardvark").ToEqual(c, false)
	c.shouldHaveCalledErrorf(t, "Expected bool ―――\nfalse\n――― type must be int, uint, or float (of any length) ―――\n")

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToEqual(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― to equal ―――\n1710000001\n")

	expect.Number(1.1).ToEqual(c, 1.1)
	c.shouldNotHaveHadAnError(t)

	expect.ApproximateFloatFraction = 1e-6
	expect.Number(1.1).ToEqual(c, 1.100001) // almost the same
	c.shouldNotHaveHadAnError(t)

	expect.Number(1.1).Using(cmpopts.EquateApprox(1e-7, 0)).ToEqual(c, 1.100001) // almost the same
	c.shouldHaveCalledErrorf(t, "Expected float64 ―――\n1.1\n――― to equal ―――\n1.100001\n")

	expect.Number(1.1).ToEqual(c, 1.10001)
	c.shouldHaveCalledErrorf(t, "Expected float64 ―――\n1.1\n――― to equal ―――\n1.10001\n")

	//----- mis-match -----

	expect.Number(utcTime).ToEqual(c, 1710000000).Or().ToEqual(c, 1710000001)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberNotToBe(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).Not().ToBe(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBe(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― not to be ―――\n1710000000\n")

	expect.Number(numberTest(errors.New("bang"))).I("data").Not().ToBe(c, 0)
	c.shouldHaveCalledFatalf(t,
		"Expected data not to pass a non-nil error but got error parameter 2 ―――\nbang\n",
		"Expected data int ―――\n0\n――― not to be ―――\n0\n",
	)

	expect.ApproximateFloatFraction = 1e-6
	expect.Number(1.1).Not().ToBe(c, 1.100001)
	c.shouldHaveCalledErrorf(t, "Expected float64 ―――\n1.1\n――― not to be ―――\n1.100001\n")

	expect.Number(1.1).Not().ToBe(c, 1.10001)
	c.shouldNotHaveHadAnError(t)
}

func TestNumberNotToEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).Not().ToEqual(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToEqual(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― not to equal ―――\n1710000000\n")

	expect.Number("aardvark").Not().ToEqual(c, 1)
	c.shouldNotHaveHadAnError(t)

	expect.Number(numberTest(errors.New("bang"))).I("data").Not().ToEqual(c, 0)
	c.shouldHaveCalledFatalf(t,
		"Expected data not to pass a non-nil error but got error parameter 2 ―――\nbang\n",
		"Expected data int ―――\n0\n――― not to equal ―――\n0\n",
	)
}

func TestNumberToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).ToBeGreaterThan(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeGreaterThan(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― to be greater than ―――\n1710000000\n")

	//----- mis-match -----

	utcTime = 1710000002
	expect.Number(utcTime).ToBeGreaterThan(c, 1710000000).Or().ToBeGreaterThan(c, 1710000001)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberNotToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).Not().ToBeGreaterThan(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThan(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000001\n――― not to be greater than ―――\n1710000000\n")
}

func TestNumberToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeLessThan(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeLessThan(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― to be less than ―――\n1710000000\n")

	//----- mis-match -----

	expect.Number(utcTime).ToBeLessThan(c, 1710000002).Or().ToBeLessThan(c, 1710000001)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberNotToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThan(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThan(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― not to be less than ―――\n1710000001\n")
}

func TestNumberToBeGreaterThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).I("utcTime").ToBeGreaterThanOrEqual(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeGreaterThanOrEqual(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― to be greater than or equal to ―――\n1710000001\n")

	//----- mis-match -----

	utcTime = 1710000002
	expect.Number(utcTime).ToBeGreaterThanOrEqual(c, 1710000000).Or().ToBeGreaterThanOrEqual(c, 1710000001)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberNotToBeGreaterThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThanOrEqual(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeGreaterThanOrEqual(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000001\n――― not to be greater than or equal to ―――\n1710000000\n")
}

func TestNumberToBeLessThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(utcTime).ToBeLessThanOrEqual(c, 1710000001)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").ToBeLessThanOrEqual(c, 1710000000)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000001\n――― to be less than or equal to ―――\n1710000000\n")

	//----- mis-match -----

	utcTime = 1710000000
	expect.Number(utcTime).ToBeLessThanOrEqual(c, 1710000002).Or().ToBeLessThanOrEqual(c, 1710000003)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberNotToBeLessThanOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).Not().ToBeLessThanOrEqual(c, 1710000000)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").Not().ToBeLessThanOrEqual(c, 1710000001)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n1710000000\n――― not to be less than or equal to ―――\n1710000001\n")
}

func TestNumberToBeBetween(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).ToBeBetween(c, 1710000000, 1710000002)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000000
	expect.Number(utcTime).I("utcTime").ToBeBetween(c, 1710000001, 1710000002)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n"+
		"1710000000\n"+
		"――― to be between ―――\n"+
		"1710000001 … 1710000002 (exclusive)\n")

	expect.Number(utcTime).I("utcTime").ToBeBetween(c, 1710000002, 1710000000)
	c.shouldHaveCalledErrorf(t, "Impossible test utcTime expect_test.Seconds32: minimum 1710000002 >= maximum 1710000000.\n")

	//----- mis-match -----

	utcTime = 1710000002
	expect.Number(utcTime).ToBeBetween(c, 1710000000, 1710000004).Or().ToBeBetween(c, 1710000005, 1710000007)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberNotToBeBetween(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).Not().ToBeBetween(c, 1710000002, 1710000003)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeBetween(c, 1710000000, 1710000002)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n"+
		"1710000001\n"+
		"――― not to be between ―――\n"+
		"1710000000 … 1710000002 (exclusive)\n")

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
		"1710000000\n"+
		"――― to be between ―――\n"+
		"1710000001 … 1710000002 (inclusive)\n")

	expect.Number(utcTime).I("utcTime").ToBeBetweenOrEqual(c, 1710000002, 1710000000)
	c.shouldHaveCalledErrorf(t, "Impossible test utcTime expect_test.Seconds32: minimum 1710000002 > maximum 1710000000.\n")

	//----- mis-match -----

	utcTime = 1710000002
	expect.Number(utcTime).ToBeBetweenOrEqual(c, 1710000000, 1710000004).Or().ToBeBetweenOrEqual(c, 1710000005, 1710000007)
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")
}

func TestNumberNotToBeBetweenOrEqual(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(utcTime).Not().ToBeBetweenOrEqual(c, 1710000002, 1710000003)
	c.shouldNotHaveHadAnError(t)

	utcTime = 1710000001
	expect.Number(utcTime).I("utcTime").Not().ToBeBetweenOrEqual(c, 1710000000, 1710000002)
	c.shouldHaveCalledErrorf(t, "Expected utcTime expect_test.Seconds32 ―――\n"+
		"1710000001\n"+
		"――― not to be between ―――\n"+
		"1710000000 … 1710000002 (inclusive)\n")

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
