package expect_test

import (
	"errors"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rickb777/expect"
	"testing"
)

type Weight32 uint32

type Info struct {
	Yin  string
	yang string
}

type MoreInfo struct {
	Extra    string
	original Info
}

func TestAnyToBe(t *testing.T) {
	c := &capture{}

	var nilBytes []byte
	emptyBytes := make([]byte, 0, 100)
	expect.Any(emptyBytes).Using(cmpopts.EquateEmpty()).ToBe(c, nilBytes)
	c.shouldNotHaveHadAnError(t)

	expect.Any(nilBytes).Using(cmpopts.EquateEmpty()).ToBe(c, emptyBytes)
	c.shouldNotHaveHadAnError(t)

	someNumbers := []int{1, 2, 3, 4, 5}
	expect.Any(someNumbers).Using(cmpopts.EquateEmpty()).ToBe(c, someNumbers)
	c.shouldNotHaveHadAnError(t)

	data := map[string]int{"a": 1, "b": 2}
	expect.Any(data).Using(cmpopts.EquateEmpty()).ToBe(c, data)
	c.shouldNotHaveHadAnError(t)

	weight := 100
	expect.Any(weight).ToBe(c, 100)
	c.shouldNotHaveHadAnError(t)

	weight = 101
	expect.Any(weight).I("weight").ToBe(c, 100)
	c.shouldHaveCalledErrorf(t, "Expected weight int ―――\n  101\n――― to be ―――\n  100\n")

	i1 := Info{Yin: "a", yang: "b"}
	expect.Any(i1).ToBe(c, i1)
	c.shouldNotHaveHadAnError(t)

	i2 := Info{Yin: "a", yang: "U"}
	expect.Any(i1).Info("foo").ToBe(c, i2)
	c.shouldHaveCalledErrorf(t, "Expected foo struct to be as shown (-want, +got) ―――\n"+
		"  expect_test.Info{\n"+
		"  \tYin:  \"a\",\n"+
		"- \tyang: \"U\",\n"+
		"+ \tyang: \"b\",\n"+
		"  }\n")

	m1 := MoreInfo{Extra: "ok", original: Info{Yin: "a", yang: "b"}}
	expect.Any(i1).ToBe(c, i1)
	c.shouldNotHaveHadAnError(t)

	m2 := MoreInfo{Extra: "ok", original: Info{Yin: "a", yang: "c"}}
	expect.Any(m1).Info("foo").ToBe(c, m2)
	c.shouldHaveCalledErrorf(t, `Expected foo struct to be as shown (-want, +got) ―――
  expect_test.MoreInfo{
  	Extra: "ok",
  	original: expect_test.Info{
  		Yin:  "a",
- 		yang: "c",
+ 		yang: "b",
  	},
  }
`)

	var fa = 0.01347258873283863
	var fb = 0.013473
	expect.Any(fa).ToBe(c, fb)
	c.shouldNotHaveHadAnError(t)
}

func TestAnyNotToBe(t *testing.T) {
	c := &capture{}

	weight := 100
	expect.Any(weight).Not().ToBe(c, 101)
	c.shouldNotHaveHadAnError(t)

	weight = 101
	expect.Any(weight).I("weight").Not().ToBe(c, 101)
	c.shouldHaveCalledErrorf(t, "Expected weight int not to be ―――\n  101\n")

	var fa = 0.01347258873283863
	var fb = 0.013573
	expect.Any(fa).Not().ToBe(c, fb)
	c.shouldNotHaveHadAnError(t)

	expect.Any(boolTest(errors.New("bang"))).I("data").Not().ToBe(c, false)
	c.shouldHaveCalledFatalf(t, "Expected data not to pass a non-nil error but got parameter 2 (*errors.errorString) ―――\n  bang\n")
}

func TestAnyToBeBytes(t *testing.T) {
	c := &capture{}

	data := []byte("hello world")
	expect.Any(data).ToBe(c, []byte("hello world"))
	c.shouldNotHaveHadAnError(t)

	data = []byte("hello world")
	expect.Any(data).I("data").ToBe(c, []byte("hello dlrow"))
	c.shouldHaveCalledErrorf(t, "Expected data []uint8 ―――\n"+
		"  [104 101 108 108 111 32 119 111 114 108 100]\n"+
		"  []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}\n"+
		"――― to be ―――\n"+
		"  [104 101 108 108 111 32 100 108 114 111 119]\n"+
		"  []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x64, 0x6c, 0x72, 0x6f, 0x77}\n")
}

func TestAnyToEqual(t *testing.T) {
	c := &capture{}

	var weight Weight32 = 1000
	expect.Any(weight).ToEqual(c, 1000)
	c.shouldNotHaveHadAnError(t)

	weight = 1001
	expect.Any(weight).I("weight").ToEqual(c, int(1000))
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 ―――\n  1001\n  0x3e9\n――― to equal int ―――\n  1000\n")
}

func TestAnyNotToEqual(t *testing.T) {
	c := &capture{}

	var weight Weight32 = 1000
	expect.Any(weight).Not().ToEqual(c, int(1001))
	c.shouldNotHaveHadAnError(t)

	weight = 1001
	expect.Any(weight).I("weight").Not().ToEqual(c, int(1001))
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 not to equal int ―――\n  1001\n")
}
