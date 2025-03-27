package expect_test

import (
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rickb777/expect"
	"testing"
)

func TestMapToBe_string_int(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).Using(cmpopts.EquateEmpty()).ToBe(c, m)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("Table").ToBe(c, map[string]int{"a": 1, "c": 3})
	c.shouldHaveCalledErrorfRE(t, "^Expected Table map\\[string\\]int len:2 ―――\n"+
		"  map\\[[ab]:[12] [ab]:[12]\\]\n"+
		"――― to be len:2 ―――\n"+
		"  map\\[[ac]:[13] [ac]:[13]\\]\n$")

}

type Thing struct {
	A int
}

func TestMapToBe_struct_struct(t *testing.T) {
	c := &capture{}

	m := map[Thing]Info{{A: 1}: {Yin: "i"}}
	expect.Map(m).Using(cmpopts.EquateEmpty()).ToBe(c, m)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("foo").ToBe(c, map[Thing]Info{})
	c.shouldHaveCalledErrorf(t, "Expected foo map[expect_test.Thing]expect_test.Info len:1 ―――\n"+
		"  map[{A:1}:{Yin:i yang:}]\n"+
		"――― to be len:0 ―――\n"+
		"  map[]\n")

}

func TestMapNotToBe(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).I("a1b2").Using(cmpopts.EquateEmpty()).Not().ToBe(c, nil)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("a1b2").Not().ToBe(c, m)
	c.shouldHaveCalledErrorfRE(t, "^Expected a1b2 map\\[string\\]int len:2 ―――\n"+
		"  map\\[[ab]:[12] [ab]:[12]\\]\n"+
		"――― not to be len:2 ―――\n"+
		"  map\\[[ac]:[13] [ab]:[12]\\]\n$")
}

func TestMapToHaveLength(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1}
	expect.Map(m).ToHaveLength(c, 1)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToHaveLength(c, 5)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToHaveLength(c, 5)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:1 ―――\n"+
		"  map[a:1]\n"+
		"――― to have length 5\n")

	expect.Map(m).Not().ToHaveLength(c, 1)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:1 ―――\n"+
		"  map[a:1]\n"+
		"――― not to have length 1\n")
}

func TestMapToBeEmpty(t *testing.T) {
	c := &capture{}

	empty := map[int]int{}
	expect.Map(empty).ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	var n map[int]int
	expect.Map(n).ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	m := map[string]int{"a": 1}
	expect.Map(m).Not().ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:1 ―――\n"+
		"  map[a:1]\n"+
		"――― to be empty\n")

	expect.Map(empty).Not().ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected map[int]int len:0 not to be empty\n")
}

func TestMapToContain(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).Using(cmpopts.EquateEmpty()).ToContain(c, "a")
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Using(cmpopts.EquateEmpty()).ToContain(c, "a", 1)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContain(c, "c")
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:2 to contain \"c\"; keys are ―――\n"+
		"  [a, b]\n")

	expect.Map(m).ToContain(c, "a", 7)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:2 ―――\n"+
		"  \"a\": 1\n"+
		"――― to contain \"a\" and it should match ―――\n"+
		"  \"a\": 7\n")
}

func TestMapNotToContain(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).Using(cmpopts.EquateEmpty()).Not().ToContain(c, "z")
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Using(cmpopts.EquateEmpty()).ToContain(c, "a", 1)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToContain(c, "a")
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:2 not to contain \"a\"; keys are ―――\n"+
		"  [a, b]\n")

	expect.Map(m).Not().ToContain(c, "a", 7)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:2 contains ―――\n"+
		"  \"a\": 1\n"+
		"――― but should contain ―――\n"+
		"  \"a\": 7\n")
}
