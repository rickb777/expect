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
	c.shouldHaveCalledErrorf(t, "Expected ...\n  hello\n... to be ...\n  \"\"\n")
}

func TestStringToEqual(t *testing.T) {
	c := &capture{}

	numbers1 := strings.Repeat("0123456789", 10) + "_" + "0123456789"

	expect.String(c, numbers1).ToEqual(numbers1)
	c.shouldNotHaveHadAnError(t)

	numbers2 := strings.Repeat("0123456789", 12)

	expect.String(c, numbers1).ToEqual(numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ...\n"+
		"  0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789...\n"+
		"... to equal ...\n"+
		"  0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789...\n")
}

func TestStringNotToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToBe("world")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).Not().ToBe("hello")
	c.shouldHaveCalledErrorf(t, "Expected ...\n  hello\n... not to be ...\n  hello\n")
}

func TestStringNotToEqual(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToEqual("world")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).Not().ToEqual("hello")
	c.shouldHaveCalledErrorf(t, "Expected ...\n  hello\n... not to equal ...\n  hello\n")
}

func TestStringToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).ToContain("ell")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).ToContain("world")
	c.shouldHaveCalledErrorf(t, "Expected ...\n  hello\n... to contain ...\n  world\n")
}

func TestStringNotToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToContain("world")
	c.shouldNotHaveHadAnError(t)

	expect.String(c, s).Not().ToContain("ell")
	c.shouldHaveCalledErrorf(t, "Expected ...\n  hello\n... not to contain ...\n  ell\n")
}
