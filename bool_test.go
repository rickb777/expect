package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

func TestToBeTrue(t *testing.T) {
	c := &capture{}

	expect.Bool(c, true, "data").ToBeTrue()
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.Bool(c, false, "data").ToBeTrue()
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected data to be true\n" {
		t.Error(c.message)
	}
}

func TestNotToBeTrue(t *testing.T) {
	c := &capture{}

	expect.Bool(c, false, "data").Not().ToBeTrue()
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.Bool(c, true, "data").Not().ToBeTrue()
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected data not to be true\n" {
		t.Error(c.message)
	}
}

func TestToBeFalse(t *testing.T) {
	c := &capture{}

	expect.Bool(c, false, "data").ToBeFalse()
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.Bool(c, true, "data").ToBeFalse()
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected data to be false\n" {
		t.Error(c.message)
	}
}

func TestNotToBeFalse(t *testing.T) {
	c := &capture{}

	expect.Bool(c, true, "data").Not().ToBeFalse()
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.Bool(c, false, "data").Not().ToBeFalse()
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected data not to be false\n" {
		t.Error(c.message)
	}
}
