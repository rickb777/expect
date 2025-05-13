package expect_test

import (
	"errors"
	"github.com/rickb777/expect"
	"regexp"
	"strings"
	"testing"
)

type MyString string

func stringTestE(e error) (string, error) { return "", e }
func stringTestOK() (string, bool, error) { return "", false, nil }

func TestStringOr(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"

	//----- match early -----

	expect.String(s).ToBe(nil, "hello").Or().ToBe(c, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToEqual(nil, "hello").Or().ToBe(c, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBe(nil, "hello").Or().ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBe(nil, "hello").Or().ToHaveLength(c, 1)
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBe(nil, "hello").Or().ToContain(c, "zzz")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBe(nil, "hello").Or().ToMatch(c, regexp.MustCompile("zzz"))
	c.shouldNotHaveHadAnError(t)

	//----- match late -----

	expect.String(s).ToBe(nil, "goodbye").Or().ToBe(c, "hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToEqual(nil, "goodbye").Or().ToBe(c, "hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBeEmpty(nil).Or().ToBe(c, "hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToHaveLength(nil, 1).Or().ToBe(c, "hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToContain(nil, "xyz").Or().ToBe(c, "hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToMatch(nil, regexp.MustCompile("zzz")).Or().ToBe(c, "hello")
	c.shouldNotHaveHadAnError(t)

	//----- mis-match -----

	expect.String(s).ToBe(c, "hello").Or().ToBe(c, "goodbye")
	c.shouldHaveCalledFatalf(t, "Incorrect test conjunction.\n"+
		"――― Only the last assertion should have a non-nil tester.\n"+
		"――― Use nil for the preceding assertions.")

	expect.String("Ron").ToBe(nil, "Fred").Or().ToBe(nil, "George").Or().ToBe(c, "Ginny")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"Ron\n"+
		"――― to be ―――\n"+
		"Fred\n"+
		"――― the first difference is at rune 0.\n"+
		"\n"+
		"――― or to be ―――\n"+
		"George\n"+
		"――― the first difference is at rune 0.\n"+
		"\n"+
		"――― or to be ―――\n"+
		"Ginny\n"+
		"――― the first difference is at rune 0.\n")

	expect.String("abcµdef-0123456789").ToBe(nil, "abcµdfe-0123456789").Or().ToBe(c, "bacµdef-0123456789")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"abcµdef-0123456789\n"+
		"――― to be ―――\n"+
		"abcµdfe-0123456789\n"+
		"――― the first difference is at rune 5 (line 1:6).\n"+
		"\n"+
		"――― or to be ―――\n"+
		"bacµdef-0123456789\n"+
		"――― the first difference is at rune 0.\n")
}

func TestStringToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToBe(c, "hello")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToBe(c, "")
	c.shouldHaveCalledErrorf(t, "Expected ―――\nhello\n――― to be blank.\n")

	expect.String(stringTestOK()).I("data").ToBe(c, "")
	c.shouldNotHaveHadAnError(t)

	expect.String("abcµdef-°123456789").ToBe(c, "abcµdfe-°123456789")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"abcµdef-°123456789\n"+
		"――― to be ―――\n"+
		"abcµdfe-°123456789\n"+
		"――― the first difference is at rune 5 (line 1:6).\n")

	numbers1 := strings.Repeat("°1234µ6789", 6)

	expect.String(numbers1+"<vwxyz").Trim(50).ToBe(c, numbers1+">vwxyz")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"…6789°1234µ6789°1234µ6789<vwxyz\n"+
		"――― to be ―――            ↕\n"+
		"…6789°1234µ6789°1234µ6789>vwxyz\n"+
		"――― the first difference is at rune 60 (line 1:61).\n")

	lines1 := "A" + numbers1 + "\nB" + numbers1 + "\nC" + numbers1 + "\nvwxyz"
	lines2 := "D" + numbers1 + "\nE" + numbers1 + "\nF" + numbers1

	expect.String(lines1+"<"+lines2).ToBe(c, lines1+">"+lines2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"A°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"B°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"C°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"vwxyz<D°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"E°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"F°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"――― to be ―――\n"+
		"A°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"B°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"C°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"vwxyz>D°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"E°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"F°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"――― the first difference is at rune 191 (line 4:6).\n")

	expect.String(lines1+"<"+lines2).Trim(100).ToBe(c, lines1+">"+lines2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"…789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"vwxyz<D°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ67…\n"+
		"――― to be ―――\n"+
		"…789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"vwxyz>D°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ67…\n"+
		"――― the first difference is at rune 191 (line 4:6).\n")
}

func TestStringToEqual(t *testing.T) {
	c := &capture{}

	expect.String([]byte("hello")).ToEqual(c, "hello")
	c.shouldNotHaveHadAnError(t)

	numbers1 := strings.Repeat("°1234µ6789", 5) + "_" + "01234«-»6789"

	expect.String(numbers1).ToEqual(c, numbers1)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcµdef-°123456789").ToEqual(c, "abcµdfe-°123456789")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"abcµdef-°123456789\n"+
		"――― to equal ―――\n"+
		"abcµdfe-°123456789\n"+
		"――― the first difference is at rune 5 (line 1:6).\n")

	numbers2 := strings.Repeat("°1234µ6789", 7)

	expect.String(numbers1).ToEqual(c, numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                  ↕\n"+
		"°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"――― the first difference is at rune 50 (line 1:51).\n")

	expect.String(numbers2+numbers1).ToEqual(c, numbers2+numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789_01234«-»6789\n"+
		"――― to equal ―――                                                                                                        ↕\n"+
		"°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789\n"+
		"――― the first difference is at rune 120 (line 1:121).\n")

	expect.String(numbers2+numbers1+numbers2).Trim(100).ToEqual(c, numbers2+numbers2+numbers2)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+
		"…1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789_01234«-»6789°1234µ6789°1234µ6789°1234µ6789°1234µ6…\n"+
		"――― to equal ―――                                  ↕\n"+
		"…1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789°1234µ6789…\n"+
		"――― the first difference is at rune 120 (line 1:121).\n")
}

func TestStringNotToBe(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToBe(c, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String("").Not().ToBe(c, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToBe(c, "")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToBe(c, "hello")
	c.shouldHaveCalledErrorf(t, "Expected ―――\nhello\n――― not to be this value.\n")

	expect.String("").Not().ToBe(c, "")
	c.shouldHaveCalledErrorf(t, "Expected ―――\n\"\"\n――― not to be blank.\n")

	numbers1 := strings.Repeat("°1234µ56789", 6)
	lines := "A" + numbers1 + "\nB" + numbers1 + "\nC" + numbers1 + "\nD" + numbers1 + "\nE" + numbers1

	expect.String(lines).Not().ToBe(c, lines)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+lines+"\n――― not to be this value.\n")

	expect.String(lines).Trim(100).Not().ToBe(c, lines)
	c.shouldHaveCalledErrorf(t, "Expected ―――\n"+lines[:118]+"…\n――― not to be this value.\n")

	expect.String(stringTestE(errors.New("bang"))).I("data").Not().ToBe(c, "")
	c.shouldHaveCalledFatalf(t,
		"Expected data not to pass a non-nil error but got error parameter 2 ―――\nbang\n",
		"Expected data ―――\n\"\"\n――― not to be blank.\n",
	)
}

func TestStringNotToEqual(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToEqual(c, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToEqual(c, "hello")
	c.shouldHaveCalledErrorf(t, "Expected ―――\nhello\n――― not to equal this value.\n")
}

func TestStringToHaveLength(t *testing.T) {
	c := &capture{}

	expect.String("abcdef").ToHaveLength(c, 6)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcdef").Not().ToHaveLength(c, 5)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcdef").ToHaveLength(c, 5)
	c.shouldHaveCalledErrorf(t, "Expected string len:6 ―――\n"+
		"abcdef\n"+
		"――― to have length 5.\n")

	expect.String("abcdef").Not().ToHaveLength(c, 6)
	c.shouldHaveCalledErrorf(t, "Expected string len:6 ―――\n"+
		"abcdef\n"+
		"――― not to have length 6.\n")

	var longString = strings.Repeat("°123456789", 10)
	expect.String(longString).ToHaveLength(c, 90)
	c.shouldHaveCalledErrorf(t, "Expected string len:110 ―――\n"+
		"°123456789°123456789°123456789°123456789°123456789°123456789°123456789°123456789°123456789°123456789\n"+
		"――― to have length 90.\n")

	expect.String(longString).Trim(80).ToHaveLength(c, 90)
	c.shouldHaveCalledErrorf(t, "Expected string len:110 ―――\n"+
		"°123456789°123456789°123456789°123456789°123456789°123456789°123456789°123456789…\n"+
		"――― to have length 90.\n")
}

func TestStringToBeEmpty(t *testing.T) {
	c := &capture{}

	expect.String([]byte{}).ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	expect.String("foo").Not().ToBeEmpty(c)
	c.shouldNotHaveHadAnError(t)

	expect.String("abcdef").ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected string len:6 ―――\n"+
		"abcdef\n"+
		"――― to be empty.\n")

	expect.String([]byte{}).Not().ToBeEmpty(c)
	c.shouldHaveCalledErrorf(t, "Expected []uint8 len:0 not to be empty.\n")
}

func TestStringToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToContain(t, "ell")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToContain(c, "world")
	c.shouldHaveCalledErrorf(t, "Expected expect_test.MyString len:5 ―――\nhello\n――― to contain ―――\nworld\n")
}

func TestStringNotToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToContain(c, "world")
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToContain(c, "ell")
	c.shouldHaveCalledErrorf(t, "Expected expect_test.MyString len:5 ―――\nhello\n――― not to contain ―――\nell\n")
}

func TestStringToMatch(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).ToMatch(c, regexp.MustCompile("^.*ll.*$"))
	c.shouldNotHaveHadAnError(t)

	expect.String(s).ToMatch(c, regexp.MustCompile("^x-ll-$"))
	c.shouldHaveCalledErrorf(t, "Expected ―――\nhello\n――― to match ―――\n^x-ll-$\n")
}

func TestStringNotToMatch(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(s).Not().ToMatch(c, regexp.MustCompile("world"))
	c.shouldNotHaveHadAnError(t)

	expect.String(s).Not().ToMatch(c, regexp.MustCompile("^.*ll.*$"))
	c.shouldHaveCalledErrorf(t, "Expected ―――\nhello\n――― not to match ―――\n^.*ll.*$\n")
}

func ExampleStringType_ToBe() {
	var t *testing.T

	// string matching can use any string, or subtype of string
	s := "hello"
	expect.String(s).ToBe(t, "hello")

	var i int // some loop counter
	expect.String(s).Info("loop %d", i).Not().ToBe(t, "goodbye")
}

func ExampleStringType_ToMatch() {
	var t *testing.T

	s := "hello"
	expect.String(s).ToMatch(t, regexp.MustCompile("^he"))

	var i int // some loop counter
	expect.String(s).Info("loop %d", i).Not().ToMatch(t, regexp.MustCompile(".*bye$"))
}

func ExampleStringType_ToContain() {
	var t *testing.T

	s := "Once more unto the breach"
	expect.String(s).ToContain(t, "unto")

	var i int // some loop counter
	expect.String(s).Info("loop %d", i).Not().ToContain(t, "foobar")
}
