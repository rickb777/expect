package expect_test

import (
	"errors"
	"github.com/rickb777/expect"
	"testing"
)

func boolTest(e error) (bool, error) { return false, e }

func TestBoolToBeTrue(t *testing.T) {
	c := &capture{}

	expect.Bool(true).Info("data").ToBeTrue(t)
	c.shouldNotHaveHadAnError(t)

	expect.Bool(false).Info("data").ToBeTrue(c)
	c.shouldHaveCalledErrorf(t, "Expected data to be true.\n")

	expect.Bool(boolTest(nil)).Info("data").ToBeTrue(c)
	c.shouldHaveCalledErrorf(t, "Expected data to be true.\n")
}

func TestBoolNotToBeTrue(t *testing.T) {
	c := &capture{}

	expect.Bool(false).I("data").Not().ToBeTrue(t)
	c.shouldNotHaveHadAnError(t)

	expect.Bool(true).I("data").Not().ToBeTrue(c)
	c.shouldHaveCalledErrorf(t, "Expected data not to be true.\n")
}

func TestBoolToBeFalse(t *testing.T) {
	c := &capture{}

	expect.Bool(false).I("data").ToBeFalse(t)
	c.shouldNotHaveHadAnError(t)

	expect.Bool(true).I("data").ToBeFalse(c)
	c.shouldHaveCalledErrorf(t, "Expected data to be false.\n")
}

func TestBoolNotToBeFalse(t *testing.T) {
	c := &capture{}

	expect.Bool(true).I("data").Not().ToBeFalse(t)
	c.shouldNotHaveHadAnError(t)

	expect.Bool(false).I("data").Not().ToBeFalse(c)
	c.shouldHaveCalledErrorf(t, "Expected data not to be false.\n")

	expect.Bool(boolTest(errors.New("bang"))).I("data").Not().ToBeFalse(c)
	c.shouldHaveCalledFatalf(t,
		"Expected data not to pass a non-nil error but got parameter 2 (*errors.errorString) ―――\n  bang\n",
		"Expected data not to be false.\n",
	)
}

func ExampleBoolType_ToBe() {
	var t *testing.T

	var v bool
	expect.Bool(v).ToBe(t, false) // or ToBeFalse(t)

	var i int // some loop counter

	// Info gives more information when the test fails, such as within a loop
	expect.Bool(v).Info("loop %d", i).Not().ToBeTrue(t)
}
