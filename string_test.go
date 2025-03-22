package expect_test

import (
	"errors"
	"github.com/rickb777/expect"
	"strings"
	"testing"
)

type MyString string

func stringTest(e error) (string, error) { return "", e }

func TestStringToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToBe("hello", t)
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBe("", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― to be ―――\n  \"\"\n")

	expect.String("abcµdef-0123456789").ToBe("abcµdfe-0123456789", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  abcµdef-0123456789\n"+
		"――― to be ―――\n"+
		"  abcµdfe-0123456789\n"+
		"――― the first difference is at character 6\n")

	numbers1 := strings.Repeat("01234µ6789", 6)
	expect.String(numbers1+"<").Trim(50).ToBe(numbers1+">", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ6789<\n"+
		"――― to be ―――                                       ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ6789>\n"+
		"――― the first difference is at character 61\n")
}

func TestStringToEqual(t *testing.T) {
	c := &capture{}

	expect.String([]byte("hello")).ToEqual("hello", t)
	c.shouldNotHaveHadAnError(t)

	numbers1 := strings.Repeat("01234µ6789", 5) + "_" + "01234«-»6789"

	expect.String(numbers1).ToEqual(numbers1, t)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcµdef-0123456789").ToEqual("abcµdfe-0123456789", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  abcµdef-0123456789\n"+
		"――― to equal ―――\n"+
		"  abcµdfe-0123456789\n"+
		"――― the first difference is at character 6\n")

	numbers2 := strings.Repeat("01234µ6789", 7)

	expect.String(numbers1).ToEqual(numbers2, c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  01234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                    ↕\n"+
		"  01234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789\n"+
		"――― the first difference is at character 51\n")

	expect.String(numbers2+numbers1).ToEqual(numbers2+numbers2, c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                                        ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789\n"+
		"――― the first difference is at character 121\n")

	expect.String(numbers2+numbers1+numbers2).Trim(100).ToEqual(numbers2+numbers2+numbers2, c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»678901234µ678901234µ6…\n"+
		"――― to equal ―――                                                        ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789…\n"+
		"――― the first difference is at character 121\n")
}

func TestStringNotToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToBe("world", t)
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToBe("hello", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to be ―――\n  hello\n")

	expect.String(stringTest(errors.New("bang"))).I("data").Not().ToBe("", c)
	c.shouldHaveCalledFatalf(t, "Expected data not to pass a non-nil error but got parameter 2 (*errors.errorString) ―――\n  bang\n")
}

func TestStringNotToEqual(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToEqual("world", t)
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToEqual("hello", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to equal ―――\n  hello\n")
}

func TestStringToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToContain("ell", t)
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToContain("world", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― to contain ―――\n  world\n")
}

func TestStringNotToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToContain("world", t)
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToContain("ell", c)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to contain ―――\n  ell\n")
}
