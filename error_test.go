package expect_test

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"testing"

	"github.com/rickb777/expect"
)

var e1 = errors.New("something bad happened")

func TestErrorNoErrorProvided(t *testing.T) {
	c := &capture{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	expect.Error(1, 2, "hello", true).I("xyz").ToBeNil(c)
	c.shouldNotHaveHadAnError(t)
}

func TestErrorToBeNil(t *testing.T) {
	c := &capture{}

	var err error
	expect.Error(err).I("xyz").ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Error(1, 2, err).I("xyz").ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Error(e1).I("xyz").ToBeNil(c)
	c.shouldHaveCalledFatalf(t, `Expected xyz error ―――
something bad happened
――― not to have occurred.
`)

	thingUnderTest1 := func() (string, bool, error) { return "foo", true, nil }
	expect.Error(thingUnderTest1()).I("xyz").ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	thingUnderTest2 := func() (string, error, bool, error) { return "foo", e1, true, nil }
	expect.Error(thingUnderTest2()).I("xyz").ToBeNil(c)
	c.shouldHaveCalledFatalf(t, `Expected xyz error ―――
something bad happened
――― not to have occurred.
`)
}

func TestErrorToHaveOccurred(t *testing.T) {
	c := &capture{}

	expect.Error(1, 2, 3, e1).I("xyz").ToHaveOccurred(c)
	c.shouldNotHaveHadAnError(t)

	expect.Error(nil).I("xyz").ToHaveOccurred(c)
	c.shouldHaveCalledErrorf(t, "Expected xyz error to have occurred.\n")

	thingUnderTest := func() (string, error, bool, error) { return "foo", e1, true, nil }
	expect.Error(thingUnderTest()).I("xyz").ToHaveOccurred(c)
	c.shouldNotHaveHadAnError(t)
}

func TestErrorNotToBeNil(t *testing.T) {
	c := &capture{}

	expect.Error(e1).I("xyz").Not().ToBeNil(c)
	c.shouldNotHaveHadAnError(t)

	expect.Error(nil).I("xyz").Not().ToBeNil(c)
	c.shouldHaveCalledErrorf(t, "Expected xyz error to have occurred.\n")
}

func TestErrorNotToHaveOccurred(t *testing.T) {
	c := &capture{}

	expect.Error(nil).I("xyz").Not().ToHaveOccurred(c)
	c.shouldNotHaveHadAnError(t)

	expect.Error(e1).I("xyz").Not().ToHaveOccurred(c)
	c.shouldHaveCalledFatalf(t, `Expected xyz error ―――
something bad happened
――― not to have occurred.
`)
}

func TestErrorToWrap(t *testing.T) {
	c := &capture{}

	expect.Error(nil).I("xyz").ToWrap(c, io.EOF)
	c.shouldHaveCalledErrorf(t, "Expected xyz error to have occurred but there was no error.\n")

	expect.Error(e1).I("xyz").Not().ToWrap(c, io.EOF)
	c.shouldNotHaveHadAnError(t)

	e2 := fmt.Errorf("foo %w", io.EOF)
	e3 := fmt.Errorf("bar %w", e2)
	expect.Error(e3).I("xyz").ToWrap(c, io.EOF)
	c.shouldNotHaveHadAnError(t)

	expect.Error(e1).I("xyz").ToWrap(c, io.EOF)
	c.shouldHaveCalledErrorf(t, `Expected xyz error ―――
something bad happened
――― to wrap ―――
*errors.errorString "EOF"
`)

	expect.Error(e3).I("xyz").Not().ToWrap(c, io.EOF)
	c.shouldHaveCalledErrorf(t, `Expected xyz error ―――
bar foo EOF
――― not to wrap ―――
*errors.errorString "EOF"
`)
}

func TestErrorToContain(t *testing.T) {
	c := &capture{}

	expect.Error(e1).I("xyz").ToContain(c, "something bad happened")
	c.shouldNotHaveHadAnError(t)

	expect.Error(e1).I("xyz").ToContain(c, "missing")
	c.shouldHaveCalledErrorf(t, `Expected xyz error ―――
something bad happened
――― to contain ―――
missing
`)

	expect.Error(0, nil).ToContain(c, "something bad happened")
	c.shouldHaveCalledErrorf(t, "Expected error to have occurred but there was no error.\n")
}

func TestErrorToMatch(t *testing.T) {
	c := &capture{}

	expect.Error(e1).I("xyz").ToMatch(c, regexp.MustCompile("something bad happened"))
	c.shouldNotHaveHadAnError(t)

	expect.Error(e1).I("xyz").ToMatch(c, regexp.MustCompile("missing"))
	c.shouldHaveCalledErrorf(t, `Expected xyz error ―――
something bad happened
――― to match ―――
missing
`)

	expect.Error(0, nil).ToMatch(c, regexp.MustCompile("something bad happened"))
	c.shouldHaveCalledErrorf(t, "Expected error to have occurred but there was no error.\n")
}

func ExampleErrorType_ToBeNil() {
	var t *testing.T

	var err error
	// ... something under test goes here
	expect.Error(err).ToBeNil(t)
	expect.Error(err).Not().ToHaveOccurred(t)

	// if there's a function that returns various results and an error...
	thingUnderTest := func() (string, error) { return "", nil }
	// ...the function return parameters can be passed straight in
	expect.Error(thingUnderTest()).Not().ToHaveOccurred(t)
}
