package expect

import (
	"io"
	"testing"
)

func TestSimpleTester(t *testing.T) {
	var a, b int

	st := SimpleTester(
		func(v ...any) {
			a++
		},
		func(v ...any) {
			b++
		})

	String("foo").ToBe(st, "bar")
	Error(io.EOF).Not().ToHaveOccurred(st)

	if a != 1 {
		t.Errorf("a = %d; want 1", a)
	}
	if b != 1 {
		t.Errorf("b = %d; want 1", b)
	}
}

func TestMakeInfo_justNumbers(t *testing.T) {
	s := makeInfo(1, 2, 3)
	if s != "1 2 3" {
		t.Errorf("makeInfo(1, 2, 3) returned %s", s)
	}
}

func TestMakeInfo_oneNumber(t *testing.T) {
	s := makeInfo(1)
	if s != "1" {
		t.Errorf("makeInfo(1) returned %s", s)
	}
}

func TestMakeInfo2(t *testing.T) {
	s := makeInfo("foo")
	if s != "foo" {
		t.Errorf("makeInfo(foo) returned %s", s)
	}
}

func TestMakeInfo3(t *testing.T) {
	s := makeInfo("foo %d", 1)
	if s != "foo 1" {
		t.Errorf("makeInfo(foo, 1) returned %s", s)
	}
}
