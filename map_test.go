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
		"map\\[[ab]:[12] [ab]:[12]\\]\n"+
		"――― to be len:2 ―――\n"+
		"map\\[[ac]:[13] [ac]:[13]\\]\n$")

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
		"map[{A:1}:{Yin:i yang:}]\n"+
		"――― to be len:0 ―――\n"+
		"map[]\n")

}

func TestMapNotToBe(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).I("a1b2").Using(cmpopts.EquateEmpty()).Not().ToBe(c, nil)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("a1b2").Not().ToBe(c, m)
	c.shouldHaveCalledErrorfRE(t, "^Expected a1b2 map\\[string\\]int len:2 ―――\n"+
		"map\\[[ab]:[12] [ab]:[12]\\]\n"+
		"――― not to be len:2 ―――\n"+
		"map\\[[ac]:[13] [ab]:[12]\\]\n$")
}

func TestMapToBeNilOrNot(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1}
	var empty map[int]int

	expect.Map(empty).ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("stuff").ToBeNil(c)
	c.shouldHaveCalledErrorf(t, "Expected stuff map[string]int len:1 ―――\n"+
		"map[a:1]\n"+
		"map[string]int{\"a\":1}\n"+
		"――― to be nil\n")

	expect.Map(m).Not().ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Map(empty).I("stuff").Not().ToBeNil(c)
	c.shouldHaveCalledErrorf(t, "Expected stuff map[int]int not to be nil\n")
}

func TestMapToHaveLength(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1}
	expect.Map(m).ToHaveLength(c, 1)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToHaveSize(c, 1)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToHaveLength(c, 5)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToHaveLength(c, 5)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:1 ―――\n"+
		"map[a:1]\n"+
		"――― to have length 5.\n")

	expect.Map(m).Not().ToHaveLength(c, 1)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:1 ―――\n"+
		"map[a:1]\n"+
		"――― not to have length 1.\n")
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
		"map[a:1]\n"+
		"――― to be empty.\n")

	expect.Map(empty).Not().ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected map[int]int len:0 not to be empty.\n")
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
		"[a, b]\n")

	expect.Map(m).ToContain(c, "a", 7)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:2 ―――\n"+
		"\"a\": 1\n"+
		"――― to contain \"a\" and it should match ―――\n"+
		"\"a\": 7\n")
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
		"[a, b]\n")

	expect.Map(m).Not().ToContain(c, "a", 1)
	c.shouldHaveCalledErrorf(t, "Expected map[string]int len:2 not to contain ―――\n"+
		"\"a\": 1\n"+
		"――― but it does.\n")
}

func TestMapToContainAll(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).ToContainAll(c, 'a', 'b', 'c')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAll(c, 'a')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAll(c, 'z', 'y', 'x', 'w', 'v')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― to contain all 5 but none were found.\n")

	expect.Map(m).ToContainAll(c, 'b', 'c', 'a', 'h', 'j')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― to contain all 5 but these 2 were missing ―――\n"+
		"[104 106]\n")

	expect.Map(m).ToContainAll(c, 'c', 'f', 'a', 'j', 'l', 'n')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― to contain all 6 but only these 2 were found ―――\n"+
		"[99 97]\n")
}

func TestMapNotToContainAll(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).Not().ToContainAll(c, 'b', 'd', 'f', 'z')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToContainAll(c, 'a', 'b', 'c')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― not to contain all 3 but they were all present.\n")
}

func TestMapToContainAny(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).ToContainAny(c, 'z', 'b', 'd', 'f', 'w')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAny(c, 'b')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAny(c, 'z', 'y', 'x', 'w', 'v')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― to contain any of 5 but none were present.\n")
}

func TestMapNotToContainAny(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).Not().ToContainAny(c, 'm', 'n', 'o', 'p')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToContainAny(c, 'a', 'j', 'z')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― not to contain any of 3 but this was found ―――\n"+
		"[97]\n")

	expect.Map(m).Not().ToContainAny(c, 'a', 'c', 'e')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― not to contain any of 3 but they were all present.\n")

	expect.Map(m).Not().ToContainAny(c, 'd', 'a', 'l', 'b')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― not to contain any of 4 but only this was missing ―――\n"+
		"[108]\n")

	expect.Map(m).Not().ToContainAny(c, 'b', 'm', 'c', 'h', 'j')
	c.shouldHaveCalledErrorf(t, "Expected map[uint8]int len:5 ―――\n"+
		"map[97:1 98:2 99:3 100:4 101:5]\n"+
		"――― not to contain any of 5 but these 2 were found ―――\n"+
		"[98 99]\n")
}

func ExampleMapType_ToBe() {
	var t *testing.T

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	// ToBe verifies all the keys and values match
	expect.Map(m).ToBe(t, map[string]int{"a": 1, "b": 2, "c": 3})
}

func ExampleMapType_ToContain() {
	var t *testing.T

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// verify one key is present
	expect.Map(m).ToContain(t, "b")

	// verify one key and its value match
	expect.Map(m).ToContain(t, "b", 2)
}

func ExampleMapType_ToContainAll() {
	var t *testing.T

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	expect.Map(m).ToContainAll(t, "a", "b")
}

func ExampleMapType_ToContainAny() {
	var t *testing.T

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	expect.Map(m).ToContainAny(t, "z", "b")
}
