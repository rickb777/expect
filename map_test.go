package expect_test

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rickb777/expect"
)

func TestMapToBe_string_int(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).Using(cmpopts.EquateEmpty()).ToBe(c, m)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("Table").ToBe(c, map[string]int{"a": 1, "c": 3})
	c.shouldHaveCalledErrorf(t, `Expected Table map len:2 (-want, +got) ―――
  map[string]int{
  	"a": 1,
+ 	"b": 2,
- 	"c": 3,
  }
`)

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
	c.shouldHaveCalledErrorf(t, `Expected foo map len:1 to be len:0 (-want, +got) ―――
  map[expect_test.Thing]expect_test.Info{
+ 	{A: 1}: {Yin: "i"},
  }
`)
}

func TestMapNotToBe(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).I("a1b2").Using(cmpopts.EquateEmpty()).Not().ToBe(c, nil)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("a1b2").Not().ToBe(c, m)
	c.shouldHaveCalledErrorfRE(t, "^Expected a1b2 map\\[string\\]int not to be len:2 ―――\n"+
		"map\\[[ab]:[12] [ab]:[12]\\]\n$")
}

func TestMapToBeNilOrNot(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1}
	var empty map[int]int

	expect.Map(empty).ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).I("stuff").ToBeNil(c)
	c.shouldHaveCalledErrorf(t, `Expected stuff map[string]int len:1 ―――
map[string]int{"a": 1}
――― to be nil
`)

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
	c.shouldHaveCalledErrorf(t, `Expected map[string]int len:1 ―――
map[a:1]
――― to have length 5.
`)

	expect.Map(m).Not().ToHaveLength(c, 1)
	c.shouldHaveCalledErrorf(t, `Expected map[string]int len:1 ―――
map[a:1]
――― not to have length 1.
`)
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
	c.shouldHaveCalledErrorf(t, `Expected map[string]int len:1 ―――
map[a:1]
――― to be empty.
`)

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
	c.shouldHaveCalledErrorf(t, `Expected map[string]int len:2 to contain "c"; keys are ―――
[a, b]
`)

	expect.Map(m).ToContain(c, "a", 7)
	c.shouldHaveCalledErrorf(t, `Expected map[string]int len:2 ―――
"a": 1
――― to contain "a" and it should match ―――
"a": 7
`)
}

func TestMapNotToContain(t *testing.T) {
	c := &capture{}

	m := map[string]int{"a": 1, "b": 2}
	expect.Map(m).Using(cmpopts.EquateEmpty()).Not().ToContain(c, "z")
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Using(cmpopts.EquateEmpty()).ToContain(c, "a", 1)
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToContain(c, "a")
	c.shouldHaveCalledErrorf(t, `Expected map[string]int len:2 not to contain "a"; keys are ―――
[a, b]
`)

	expect.Map(m).Not().ToContain(c, "a", 1)
	c.shouldHaveCalledErrorf(t, `Expected map[string]int len:2 not to contain ―――
"a": 1
――― but it does.
`)
}

func TestMapToContainAll(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).ToContainAll(c, 'a', 'b', 'c')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAll(c, 'a')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAll(c, 'z', 'y', 'x', 'w', 'v')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― to contain all 5 but none were found.
`)

	expect.Map(m).ToContainAll(c, 'b', 'c', 'a', 'h', 'j')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― to contain all 5 but these 2 were missing ―――
[104 106]
`)

	expect.Map(m).ToContainAll(c, 'c', 'f', 'a', 'j', 'l', 'n')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― to contain all 6 but only these 2 were found ―――
[99 97]
`)
}

func TestMapNotToContainAll(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).Not().ToContainAll(c, 'b', 'd', 'f', 'z')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToContainAll(c, 'a', 'b', 'c')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― not to contain all 3 but they were all present.
`)
}

func TestMapToContainAny(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).ToContainAny(c, 'z', 'b', 'd', 'f', 'w')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAny(c, 'b')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).ToContainAny(c, 'z', 'y', 'x', 'w', 'v')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― to contain any of 5 but none were present.
`)
}

func TestMapNotToContainAny(t *testing.T) {
	c := &capture{}

	m := map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
	expect.Map(m).Not().ToContainAny(c, 'm', 'n', 'o', 'p')
	c.shouldNotHaveHadAnError(t)

	expect.Map(m).Not().ToContainAny(c, 'a', 'j', 'z')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― not to contain any of 3 but this was found ―――
[97]
`)

	expect.Map(m).Not().ToContainAny(c, 'a', 'c', 'e')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― not to contain any of 3 but they were all present.
`)

	expect.Map(m).Not().ToContainAny(c, 'd', 'a', 'l', 'b')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― not to contain any of 4 but only this was missing ―――
[108]
`)

	expect.Map(m).Not().ToContainAny(c, 'b', 'm', 'c', 'h', 'j')
	c.shouldHaveCalledErrorf(t, `Expected map[uint8]int len:5 ―――
map[97:1 98:2 99:3 100:4 101:5]
――― not to contain any of 5 but these 2 were found ―――
[98 99]
`)
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
