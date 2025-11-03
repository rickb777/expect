package expect_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rickb777/expect"
)

type MyBytes []byte

//func stringTestE(e error) (string, error) { return "", e }

func TestSliceToBe_byte(t *testing.T) {
	c := &capture{}

	var s = MyBytes("abcdef")
	expect.Slice(s).Using(cmpopts.EquateEmpty()).ToBe(c, MyBytes("abcdef")...)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).I("MyBytes").ToBe(c, MyBytes(nil)...)
	c.shouldHaveCalledErrorf(t, "Expected MyBytes []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"[]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66}\n"+
		"――― to be len:0 ―――\n"+
		"[]\n"+
		"[]byte(nil)\n")

	expect.Slice(s).ToBe(c, MyBytes("abcµdef")...)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"[]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66}\n"+
		"――― to be len:8 ―――\n"+
		"[97 98 99 194 181 100 101 102]\n"+
		"[]byte{0x61, 0x62, 0x63, 0xc2, 0xb5, 0x64, 0x65, 0x66}\n"+
		"――― the first difference is at index 3.\n")
}

func TestSliceToBe_struct(t *testing.T) {
	c := &capture{}

	i1 := Info{Yin: "a", yang: "b"}
	i2 := Info{Yin: "c", yang: "d"}

	var s = []Info{i1, i2}
	expect.Slice(s).Using(cmpopts.EquateEmpty()).ToBe(c, i1, i2)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).I("Foo").ToBe(c, i2)
	c.shouldHaveCalledErrorf(t, "Expected Foo []expect_test.Info len:2 ―――\n"+
		"[{Yin:a yang:b} {Yin:c yang:d}]\n"+
		"[]expect_test.Info{expect_test.Info{Yin:\"a\", yang:\"b\"}, expect_test.Info{Yin:\"c\", yang:\"d\"}}\n"+
		"――― to be len:1 ―――\n"+
		"[{Yin:c yang:d}]\n"+
		"[]expect_test.Info{expect_test.Info{Yin:\"c\", yang:\"d\"}}\n"+
		"――― the first difference is at index 0.\n")
}

func TestSliceNotToBe(t *testing.T) {
	c := &capture{}

	var s = MyBytes("abcdef")
	expect.Slice(s).Not().ToBe(c, MyBytes("abcdeg")...)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).Not().ToBe(c, s...)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"[]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66}\n"+
		"――― not to be len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"[]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66}\n")
}

func TestSliceToBeNilOrNot(t *testing.T) {
	c := &capture{}

	s := []int{1, 2, 3}
	var empty []int

	expect.Slice(empty).ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).I("stuff").ToBeNil(c)
	c.shouldHaveCalledErrorf(t, "Expected stuff []int len:3 ―――\n"+
		"[1 2 3]\n"+
		"[]int{1, 2, 3}\n"+
		"――― to be nil.\n")

	expect.Slice(s).Not().ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(empty).I("stuff").Not().ToBeNil(c)
	c.shouldHaveCalledErrorf(t, "Expected stuff []int not to be nil.\n")
}

func TestSliceToHaveLength(t *testing.T) {
	c := &capture{}

	var s = MyBytes("abcdef")
	expect.Slice(s).ToHaveLength(c, 6)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).Not().ToHaveLength(c, 5)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).ToHaveLength(c, 5)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― to have length 5.\n")

	expect.Slice(s).Not().ToHaveLength(c, 6)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― not to have length 6.\n")

	var longSlice = MyBytes(strings.Repeat("0123456789", 10))
	expect.Slice(longSlice).ToHaveLength(c, 90)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:100 ―――\n"+
		"[48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54 55 56 57]\n"+
		"――― to have length 90.\n")
}

func TestSliceToBeEmpty(t *testing.T) {
	c := &capture{}

	expect.Slice([]byte{}).ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	var s = MyBytes("abcdef")
	expect.Slice(s).Not().ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― to be empty.\n")

	expect.Slice([]byte{}).Not().ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:0 not to be empty.\n")
}

func TestSliceToContainAll(t *testing.T) {
	c := &capture{}

	var s = MyBytes("abcdef")
	expect.Slice(s).ToContainAll(c, 'b', 'd', 'f')
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).ToContainAll(c, 'z', 'y', 'x', 'w', 'v')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― to contain all 5 but none were found.\n")

	expect.Slice(s).ToContainAll(c, 'b', 'd', 'f', 'h', 'j')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― to contain all 5 but these 2 were missing ―――\n"+
		"[104 106]\n")

	expect.Slice(s).ToContainAll(c, 'd', 'f', 'h', 'j', 'l', 'n')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― to contain all 6 but only these 2 were found ―――\n"+
		"[100 102]\n")

	expect.Slice([]*int{ptr(1), ptr(2)}).ToContainAll(c, ptr(2), ptr(1))
	c.shouldNotHaveHadAnError(t)

	a := mustParseURL("http://x.com/a")
	b := mustParseURL("http://y.com/b")
	expect.Slice([]*url.URL{a, b}).ToContainAll(c, mustParseURL("http://x.com/a"), mustParseURL("http://y.com/b"))
	c.shouldNotHaveHadAnError(t)
}

func TestSliceNotToContainAll(t *testing.T) {
	c := &capture{}

	var s = MyBytes("abcdef")
	expect.Slice(s).Not().ToContainAll(c, 'b', 'd', 'f', 'z')
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).Not().ToContainAll(c, 'a', 'b', 'd', 'f')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― not to contain all 4 but they were all present.\n")
}

func TestSliceToContainAny(t *testing.T) {
	c := &capture{}

	var s = MyBytes("abcdef")
	expect.Slice(s).ToContainAny(c, 'z', 'b', 'd', 'f', 'w')
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).ToContainAny(c, 'z', 'y', 'x', 'w', 'v')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― to contain any of 5 but none were present.\n")

	expect.Slice([]*int{ptr(1), ptr(2)}).ToContainAny(c, ptr(2), ptr(3))
	c.shouldNotHaveHadAnError(t)
}

func TestSliceNotToContainAny(t *testing.T) {
	c := &capture{}

	var s = MyBytes("abcdef")
	expect.Slice(s).Not().ToContainAny(c, 'm', 'n', 'o', 'p')
	c.shouldNotHaveHadAnError(t)

	expect.Slice(s).Not().ToContainAny(c, 'a', 'b', 'd', 'f')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― not to contain any of 4 but they were all present.\n")

	expect.Slice(s).Not().ToContainAny(c, 'd', 'f', 'b', 'a', 'l', 'n')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― not to contain any of 6 but only these 2 were missing ―――\n"+
		"[108 110]\n")

	expect.Slice(s).Not().ToContainAny(c, 'b', 'm', 'f', 'h', 'j')
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:6 ―――\n"+
		"[97 98 99 100 101 102]\n"+
		"――― not to contain any of 5 but these 2 were found ―――\n"+
		"[98 102]\n")
}

func mustParseURL(s string) *url.URL {
	u, e := url.Parse(s)
	if e != nil {
		panic(e)
	}
	return u
}

func ptr[T any](x T) *T { return &x }
