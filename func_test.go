package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

func TestFuncToPanic(t *testing.T) {
	c := &capture{}

	expect.Func(func() { panic("ouch") }).Info("my func").ToPanic(c)
	c.shouldNotHaveHadAnError(t)

	expect.Func(func() {}).Info("my func").ToPanic(c)
	c.shouldHaveCalledErrorf(t, "Expected my func to panic.\n")

	expect.Func(func() { panic("happy") }).Info("my func").ToPanicWithMessage(c, "ouch")
	c.shouldHaveCalledErrorf(t, "Expected my func to panic with a message containing ―――\n"+
		"  ouch\n"+
		"――― but got ―――\n"+
		"  happy\n")

	expect.Func(func() { panic(123) }).Info("my func").ToPanicWithMessage(c, "ouch")
	c.shouldHaveCalledErrorf(t, "Expected my func to panic with a string containing ―――\n"+
		"  ouch\n"+
		"――― but got int ―――\n"+
		"  123\n")
}

func TestFuncNotToPanic(t *testing.T) {
	c := &capture{}

	expect.Func(func() {}).I("my func").Not().ToPanic(c)
	c.shouldNotHaveHadAnError(t)

	expect.Func(func() { panic("ouch") }).I("my func").Not().ToPanic(c)
	c.shouldHaveCalledErrorf(t, "Expected my func not to panic.\n")
}
