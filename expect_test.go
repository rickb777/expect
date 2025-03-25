package expect

import "testing"

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
