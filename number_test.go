package expect_test

import (
	"github.com/rickb777/expect"
	"testing"
)

type Seconds32 uint32

func TestNumberToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime).ToBeGreaterThan(1710000000)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeGreaterThan(1710000000)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000000\n... to be greater than ...\n  1710000000\n" {
		t.Error(c.message)
	}
}

func TestNumberNotToBeGreaterThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime).Not().ToBeGreaterThan(1710000001)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000001
	expect.Number(c, utcTime, "utcTime").Not().ToBeGreaterThan(1710000000)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000001\n... not to be greater than ...\n  1710000000\n" {
		t.Error(c.message)
	}
}

func TestNumberToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeLessThan(1710000001)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeLessThan(1710000000)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000000\n... to be less than ...\n  1710000000\n" {
		t.Error(c.message)
	}
}

func TestNumberNotToBeLessThan(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime, "utcTime").Not().ToBeLessThan(1710000000)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").Not().ToBeLessThan(1710000001)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000000\n... not to be less than ...\n  1710000001\n" {
		t.Error(c.message)
	}
}

func TestNumberToBeGreaterThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime, "utcTime").ToBeGreaterThanOrEqualTo(1710000000)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").ToBeGreaterThanOrEqualTo(1710000001)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000000\n... to be greater than or equal to ...\n  1710000001\n" {
		t.Error(c.message)
	}
}

func TestNumberNotToBeGreaterThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime, "utcTime").Not().ToBeGreaterThanOrEqualTo(1710000001)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000001
	expect.Number(c, utcTime, "utcTime").Not().ToBeGreaterThanOrEqualTo(1710000000)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000001\n... not to be greater than or equal to ...\n  1710000000\n" {
		t.Error(c.message)
	}
}

func TestNumberToBeLessThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000000
	expect.Number(c, utcTime).ToBeLessThanOrEqualTo(1710000001)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000001
	expect.Number(c, utcTime, "utcTime").ToBeLessThanOrEqualTo(1710000000)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000001\n... to be less than or equal to ...\n  1710000000\n" {
		t.Error(c.message)
	}
}

func TestNumberNotToBeLessThanOrEqualTo(t *testing.T) {
	c := &capture{}

	var utcTime Seconds32 = 1710000001
	expect.Number(c, utcTime).Not().ToBeLessThanOrEqualTo(1710000000)
	if c.called {
		t.Error("failed")
	}

	c.reset()

	utcTime = 1710000000
	expect.Number(c, utcTime, "utcTime").Not().ToBeLessThanOrEqualTo(1710000001)
	if !c.called {
		t.Error("failed")
	}
	if c.message != "Expected utcTime expect_test.Seconds32 ...\n  1710000000\n... not to be less than or equal to ...\n  1710000001\n" {
		t.Error(c.message)
	}
}
