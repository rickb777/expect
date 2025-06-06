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
	original any
	weights  []Weight32
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
	c.shouldHaveCalledErrorf(t, "Expected weight int ―――\n101\n――― to be ―――\n100\n")

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

	m1 := MoreInfo{Extra: "ok", original: Info{Yin: "a", yang: "b"}, weights: []Weight32{1, 2, 3}}
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
- 	weights: nil,
+ 	weights: []expect_test.Weight32{1, 2, 3},
  }
`)

	var fa = 0.01347258873283863
	var fb = 0.013473
	expect.Any(fa).ToBe(c, fb)
	c.shouldNotHaveHadAnError(t)

	var undefined any
	expect.Any(undefined).Not().ToBe(c, 123)
	c.shouldNotHaveHadAnError(t)
}

func TestAnyNotToBe(t *testing.T) {
	c := &capture{}

	weight := 100
	expect.Any(weight).Not().ToBe(c, 101)
	c.shouldNotHaveHadAnError(t)

	weight = 101
	expect.Any(weight).I("weight").Not().ToBe(c, 101)
	c.shouldHaveCalledErrorf(t, "Expected weight int not to be ―――\n101\n")

	var fa = 0.01347258873283863
	var fb = 0.013573
	expect.Any(fa).Not().ToBe(c, fb)
	c.shouldNotHaveHadAnError(t)

	expect.Any(boolTest(errors.New("bang"))).I("data").Not().ToBe(c, false)
	c.shouldHaveCalledFatalf(t,
		"Expected data not to pass a non-nil error but got error parameter 2 ―――\nbang\n",
		"Expected data bool not to be ―――\nfalse\n",
	)
}

func TestAnyToBeBytes(t *testing.T) {
	c := &capture{}

	data := []byte("hello world")
	expect.Any(data).ToBe(c, []byte("hello world"))
	c.shouldNotHaveHadAnError(t)

	data = []byte("hello world")
	expect.Any(data).I("data").ToBe(c, []byte("hello dlrow"))
	c.shouldHaveCalledErrorf(t, "Expected data []uint8 ―――\n"+
		"[104 101 108 108 111 32 119 111 114 108 100]\n"+
		"[]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}\n"+
		"――― to be ―――\n"+
		"[104 101 108 108 111 32 100 108 114 111 119]\n"+
		"[]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x64, 0x6c, 0x72, 0x6f, 0x77}\n")
}

func TestAnyToBeNilOrNot(t *testing.T) {
	c := &capture{}

	var weight *int
	expect.Any(weight).ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Any("hello").I("weight").ToBeNil(c)
	c.shouldHaveCalledErrorf(t, "Expected weight string ―――\nhello\n\"hello\"\n――― to be nil.\n")

	expect.Any(1).Not().ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Any(weight).I("weight").Not().ToBeNil(c)
	c.shouldHaveCalledErrorf(t, "Expected weight *int not to be nil.\n")
}

func TestAnyToEqualOrNot(t *testing.T) {
	c := &capture{}

	var weight Weight32 = 1000
	expect.Any(weight).ToEqual(c, 1000)
	c.shouldNotHaveHadAnError(t)

	weight = 1001
	expect.Any(weight).I("weight").ToEqual(c, 1000)
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 ―――\n1001\n0x3e9\n――― to equal int ―――\n1000\n")

	expect.Any(weight).Not().ToEqual(c, 1000)
	c.shouldNotHaveHadAnError(t)

	weight = 1001
	expect.Any(weight).I("weight").Not().ToEqual(c, 1001)
	c.shouldHaveCalledErrorf(t, "Expected weight expect_test.Weight32 not to equal int ―――\n1001\n")
}

func ExampleAnyType_ToBe() {
	var t *testing.T

	type pair struct {
		a, b int
	}

	// Any matching is most useful for structs (but can test anything)
	v := pair{1, 2}
	expect.Any(v).ToBe(t, pair{1, 2})
	expect.Any(v).Not().ToBe(t, pair{3, 4})

	var i int // some loop counter

	// Info gives more information when the test fails, such as within a loop
	expect.Any(v).Info("loop %d", i).ToBe(t, pair{1, 2})
}
