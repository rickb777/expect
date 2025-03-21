package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

type MyString string

func TestStringToEqual(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).ToBe("hello")
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.String(c, s).ToBe("world")
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected ...\n  hello\n... to equal ...\n  world\n" {
		t.Error(c.message)
	}
}

func TestStringNotToEqual(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToBe("world")
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.String(c, s).Not().ToBe("hello")
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected ...\n  hello\n... not to equal ...\n  hello\n" {
		t.Error(c.message)
	}
}

func TestStringToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).ToContain("ell")
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.String(c, s).ToContain("world")
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected ...\n  hello\n... to contain ...\n  world\n" {
		t.Error(c.message)
	}
}

func TestStringNotToContain(t *testing.T) {
	c := &capture{}

	var s MyString = "hello"
	expect.String(c, s).Not().ToContain("world")
	if c.called {
		t.Error("failed")
	}

	c.reset()

	expect.String(c, s).Not().ToContain("ell")
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected ...\n  hello\n... not to contain ...\n  ell\n" {
		t.Error(c.message)
	}
}
