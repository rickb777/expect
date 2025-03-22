package expect_test

import (
	"errors"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rickb777/expect"
	"testing"
)

type Weight32 uint32

func TestAnyToEqual(t *testing.T) {
	c := &capture{}

	var nilBytes []byte
	emptyBytes := make([]byte, 0, 100)
	expect.Any(emptyBytes).Using(cmpopts.EquateEmpty()).ToBe(nilBytes, c)
	c.shouldNotHaveHadAnError(t)

	expect.Any(nilBytes).Using(cmpopts.EquateEmpty()).ToBe(emptyBytes, c)
	c.shouldNotHaveHadAnError(t)

	someNumbers := []int{1, 2, 3, 4, 5}
	expect.Any(someNumbers).Using(cmpopts.EquateEmpty()).ToBe(someNumbers, c)
	c.shouldNotHaveHadAnError(t)

	data := map[string]int{"a": 1, "b": 2}
	expect.Any(data).Using(cmpopts.EquateEmpty()).ToBe(data, c)
	c.shouldNotHaveHadAnError(t)

	weight := 100
	expect.Any(weight).ToBe(100, c)
	c.shouldNotHaveHadAnError(t)

	weight = 101
	expect.Any(weight).I("weight").ToBe(100, c)
	c.shouldHaveCalledErrorf(t, "Expected weight int ―――\n  101\n――― to equal ―――\n  100\n")

	var fa = 0.01347258873283863
	var fb = 0.013473
	expect.Any(fa).ToBe(fb, c)
	c.shouldNotHaveHadAnError(t)
}

func TestAnyNotToEqual(t *testing.T) {
	c := &capture{}

	weight := 100
	expect.Any(weight).Not().ToBe(101, c)
	c.shouldNotHaveHadAnError(t)

	weight = 101
	expect.Any(weight).I("weight").Not().ToBe(101, c)
	c.shouldHaveCalledErrorf(t, "Expected weight int ―――\n  101\n――― not to equal ―――\n  101\n")

	var fa = 0.01347258873283863
	var fb = 0.013573
	expect.Any(fa).Not().ToBe(fb, c)
	c.shouldNotHaveHadAnError(t)

	expect.Any(boolTest(errors.New("bang"))).I("data").Not().ToBe(false, c)
	c.shouldHaveCalledFatalf(t, "Expected data not to pass a non-nil error but got parameter 2 (*errors.errorString) ―――\n  bang\n")
}

func TestAnyToEqualBytes(t *testing.T) {
	c := &capture{}

	data := []byte("hello world")
	expect.Any(data).ToBe([]byte("hello world"), c)
	c.shouldNotHaveHadAnError(t)

	data = []byte("hello world")
	expect.Any(data).I("data").ToBe([]byte("hello dlrow"), c)
	c.shouldHaveCalledErrorf(t, "Expected data []uint8 ―――\n"+
		"  [104 101 108 108 111 32 119 111 114 108 100]\n"+
		"  []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}\n"+
		"――― to equal ―――\n"+
		"  [104 101 108 108 111 32 100 108 114 111 119]\n"+
		"  []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x64, 0x6c, 0x72, 0x6f, 0x77}\n")
}

func TestAnyToBeEquivalentTo(t *testing.T) {
	c := &capture{}

	var weight Weight32 = 1000
	expect.Any(weight).ToEqual(1000, c)
	c.shouldNotHaveHadAnError(t)

	weight = 1001
	expect.Any(weight).I("weight").ToEqual(1000, c)
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 ―――\n  1001\n  0x3e9\n――― to be equivalent to int ―――\n  1000\n")
}

func TestAnyNotToBeEquivalentTo(t *testing.T) {
	c := &capture{}

	var weight Weight32 = 1000
	expect.Any(weight).Not().ToEqual(1001, c)
	c.shouldNotHaveHadAnError(t)

	weight = 1001
	expect.Any(weight).I("weight").Not().ToEqual(1001, c)
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 ―――\n  1001\n  0x3e9\n――― not to be equivalent to int ―――\n  1001\n")
}
