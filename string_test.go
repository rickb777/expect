package expect_test

import (
	"github.com/rickb777/expect"
	"strings"
	"testing"
)

type MyString string

func TestStringToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).ToBe("hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).ToBe("")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― to be ―――\n  \"\"\n")

	expect.String(c, "abcµdef-0123456789").ToBe("abcµdfe-0123456789")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  abcµdef-0123456789\n"+
		"――― to be ―――\n"+
		"  abcµdfe-0123456789\n"+
		"――― the first difference is at character 6\n")

	numbers1 := strings.Repeat("01234µ6789", 6)
	expect.String(c, numbers1+"<").Trim(50).ToBe(numbers1 + ">")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ6789<\n"+
		"――― to be ―――                                       ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ6789>\n"+
		"――― the first difference is at character 61\n")
}

func TestStringToEqual(t *testing.T) {
	c := &capture{}

	numbers1 := strings.Repeat("01234µ6789", 5) + "_" + "01234«-»6789"

	expect.String(c, numbers1).ToEqual(numbers1)
	c.shouldNotHaveHadAnError(t)

	expect.String(c, "abcµdef-0123456789").ToEqual("abcµdfe-0123456789")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  abcµdef-0123456789\n"+
		"――― to equal ―――\n"+
		"  abcµdfe-0123456789\n"+
		"――― the first difference is at character 6\n")

	numbers2 := strings.Repeat("01234µ6789", 7)

	expect.String(c, numbers1).ToEqual(numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  01234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                    ↕\n"+
		"  01234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789\n"+
		"――― the first difference is at character 51\n")

	expect.String(c, numbers2+numbers1).ToEqual(numbers2 + numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                                        ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789\n"+
		"――― the first difference is at character 121\n")
}

func TestStringNotToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToBe("world")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).Not().ToBe("hello")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to be ―――\n  hello\n")
}

func TestStringNotToEqual(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToEqual("world")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).Not().ToEqual("hello")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to equal ―――\n  hello\n")
}

func TestStringToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).ToContain("ell")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).ToContain("world")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― to contain ―――\n  world\n")
}

func TestStringNotToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToContain("world")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).Not().ToContain("ell")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to contain ―――\n  ell\n")
}
