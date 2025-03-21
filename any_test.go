package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

type Weight32 uint32

func TestAnyToEqual(t *testing.T) {
	c := &capture{}

	weight := 1710000000
	expect.Any(c, weight).ToBe(1710000000)
	c.shouldNotHaveHadAnError(t)

	weight = 1710000001
	expect.Any(c, weight, "weight").ToBe(1710000000)
	c.shouldHaveCalledErrorf(t, "Expected weight int ...\n  1710000001\n... to equal ...\n  1710000000\n")

	var fa = 0.01347258873283863
	var fb = 0.013473
	expect.Any(c, fa).ToBe(fb)
	c.shouldNotHaveHadAnError(t)
}

func TestAnyNotToEqual(t *testing.T) {
	c := &capture{}

	weight := 1710000000
	expect.Any(c, weight).Not().ToBe(1710000001)
	c.shouldNotHaveHadAnError(t)

	weight = 1710000001
	expect.Any(c, weight, "weight").Not().ToBe(1710000001)
	c.shouldHaveCalledErrorf(t, "Expected weight int ...\n  1710000001\n... not to equal ...\n  1710000001\n")

	var fa = 0.01347258873283863
	var fb = 0.013573
	expect.Any(c, fa).Not().ToBe(fb)
	c.shouldNotHaveHadAnError(t)
}

func TestAnyToEqualBytes(t *testing.T) {
	c := &capture{}

	data := []byte("hello world")
	expect.Any(c, data).ToBe([]byte("hello world"))
	c.shouldNotHaveHadAnError(t)

	data = []byte("hello world")
	expect.Any(c, data, "data").ToBe([]byte("hello dlrow"))
	c.shouldHaveCalledErrorf(t, "Expected data []uint8 ...\n"+
		"  [104 101 108 108 111 32 119 111 114 108 100]\n"+
		"  []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}\n"+
		"... to equal ...\n"+
		"  [104 101 108 108 111 32 100 108 114 111 119]\n"+
		"  []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x64, 0x6c, 0x72, 0x6f, 0x77}\n")
}

func TestAnyToBeEquivalentTo(t *testing.T) {
	c := &capture{}

	var weight Weight32 = 1234
	expect.Any(c, weight).ToEqual(1234)
	c.shouldNotHaveHadAnError(t)

	weight = 1235
	expect.Any(c, weight, "weight").ToEqual(1234)
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 ...\n  1235\n  0x4d3\n... to be equivalent to int ...\n  1234\n")
}

func TestAnyNotToBeEquivalentTo(t *testing.T) {
	c := &capture{}

	var weight Weight32 = 1234
	expect.Any(c, weight).Not().ToEqual(1235)
	c.shouldNotHaveHadAnError(t)

	weight = 1235
	expect.Any(c, weight, "weight").Not().ToEqual(1235)
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 ...\n  1235\n  0x4d3\n... not to be equivalent to int ...\n  1235\n")
}
