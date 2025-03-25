package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

func TestFuncToPanic(t *testing.T) {
	c := &capture{}

	expect.Func(func() { panic("ouch") }).Info("my func").ToPanic(c)
	c.shouldNotHaveHadAnError(t)

	expect.Func(func() { panic("happy") }).Info("my func").ToPanicWithMessage(c, "ouch")
	c.shouldHaveCalledErrorf(t, "Expected my func to panic with message containing ―――\n"+
		"  ouch\n"+
		"――― but got ―――\n"+
		"  happy\n")

	expect.Func(func() {}).Info("my func").ToPanic(c)
	c.shouldHaveCalledErrorf(t, "Expected my func to panic\n")
}

func TestFuncNotToPanic(t *testing.T) {
	c := &capture{}

	expect.Func(func() {}).I("my func").Not().ToPanic(c)
	c.shouldNotHaveHadAnError(t)

	expect.Func(func() { panic("ouch") }).I("my func").Not().ToPanic(c)
	c.shouldHaveCalledErrorf(t, "Expected my func not to panic\n")
}
