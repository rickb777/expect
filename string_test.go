package expect_test

import (
	"errors"
	"github.com/rickb777/expect"
	"regexp"
	"strings"
	"testing"
)

type MyString string

func stringTest(e error) (string, error) { return "", e }

func TestStringToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToBe(t, "hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBe(c, "")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― to be ―――\n  \"\"\n")

	expect.String("abcµdef-0123456789").ToBe(c, "abcµdfe-0123456789")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  abcµdef-0123456789\n"+
		"――― to be ―――\n"+
		"  abcµdfe-0123456789\n"+
		"――― the first difference is at index 5\n")

	numbers1 := strings.Repeat("01234µ6789", 6)
	expect.String(numbers1+"<").Trim(50).ToBe(c, numbers1+">")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ6789<\n"+
		"――― to be ―――                                       ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ6789>\n"+
		"――― the first difference is at index 60\n")
}

func TestStringToEqual(t *testing.T) {
	c := &capture{}

	expect.String([]byte("hello")).ToEqual(t, "hello")
	c.shouldNotHaveHadAnError(t)

	numbers1 := strings.Repeat("01234µ6789", 5) + "_" + "01234«-»6789"

	expect.String(numbers1).ToEqual(t, numbers1)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcµdef-0123456789").ToEqual(c, "abcµdfe-0123456789")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  abcµdef-0123456789\n"+
		"――― to equal ―――\n"+
		"  abcµdfe-0123456789\n"+
		"――― the first difference is at index 5\n")

	numbers2 := strings.Repeat("01234µ6789", 7)

	expect.String(numbers1).ToEqual(c, numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  01234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                    ↕\n"+
		"  01234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789\n"+
		"――― the first difference is at index 50\n")

	expect.String(numbers2+numbers1).ToEqual(c, numbers2+numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                                        ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789\n"+
		"――― the first difference is at index 120\n")

	expect.String(numbers2+numbers1+numbers2).Trim(100).ToEqual(c, numbers2+numbers2+numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789_01234«-»678901234µ678901234µ6…\n"+
		"――― to equal ―――                                                        ↕\n"+
		"  …1234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ678901234µ6789…\n"+
		"――― the first difference is at index 120\n")
}

func TestStringNotToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToBe(t, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToBe(c, "hello")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to be ―――\n  hello\n")

	expect.String(stringTest(errors.New("bang"))).I("data").Not().ToBe(c, "")
	c.shouldHaveCalledFatalf(t, "Expected data not to pass a non-nil error but got parameter 2 (*errors.errorString) ―――\n  bang\n")
}

func TestStringNotToEqual(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToEqual(t, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToEqual(c, "hello")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to equal ―――\n  hello\n")
}

func TestStringToHaveLength(t *testing.T) {
	c := &capture{}

	expect.String("abcdef").ToHaveLength(c, 6)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcdef").Not().ToHaveLength(c, 5)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcdef").ToHaveLength(c, 5)
	c.shouldHaveCalledErrorf(t, "Expected string len:6 ―――\n"+
		"  abcdef\n"+
		"――― to have length 5\n")

	expect.String("abcdef").Not().ToHaveLength(c, 6)
	c.shouldHaveCalledErrorf(t, "Expected string len:6 ―――\n"+
		"  abcdef\n"+
		"――― not to have length 6\n")

	var longString = strings.Repeat("0123456789", 10)
	expect.String(longString).ToHaveLength(c, 90)
	c.shouldHaveCalledErrorf(t, "Expected string len:100 ―――\n"+
		"  0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789\n"+
		"――― to have length 90\n")

	expect.String(longString).Trim(80).ToHaveLength(c, 90)
	c.shouldHaveCalledErrorf(t, "Expected string len:100 ―――\n"+
		"  01234567890123456789012345678901234567890123456789012345678901234567890123456789…\n"+
		"――― to have length 90\n")
}

func TestStringToBeEmpty(t *testing.T) {
	c := &capture{}

	expect.String([]byte{}).ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	expect.String("foo").Not().ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcdef").ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected string len:6 ―――\n"+
		"  abcdef\n"+
		"――― to be empty\n")

	expect.String([]byte{}).Not().ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:0 not to be empty\n")
}

func TestStringToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToContain(t, "ell")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToContain(c, "world")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― to contain ―――\n  world\n")
}

func TestStringNotToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToContain(t, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToContain(c, "ell")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to contain ―――\n  ell\n")
}

func TestStringToMatch(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToMatch(t, regexp.MustCompile("^.*ll.*$"))
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToMatch(c, regexp.MustCompile("^x-ll-$"))
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― to match ―――\n  ^x-ll-$\n")
}

func TestStringNotToMatch(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToMatch(t, regexp.MustCompile("world"))
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToMatch(c, regexp.MustCompile("^.*ll.*$"))
	c.shouldHaveCalledErrorf(t, "Expected ―――\n  hello\n――― not to match ―――\n  ^.*ll.*$\n")
}
