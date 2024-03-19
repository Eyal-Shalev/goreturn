package goreturn_test

import (
	"testing"
	"time"

	"github.com/Eyal-Shalev/goreturn"
)

func TestReturn0(t *testing.T) {
	called := false
	fn := func() {
		called = true
	}
	c := goreturn.Return0(fn)
	select {
	case <-c:
		if !called {
			t.Errorf("Return0() did not call the function")
		}
	case <-time.After(time.Millisecond):
		t.Errorf("Return0() took too long to return")
	}
}

func TestReturn1(t *testing.T) {
	called := false
	fn := func() int {
		called = true
		return 42
	}
	c := goreturn.Return1(fn)
	select {
	case v := <-c:
		if !called {
			t.Errorf("Return1() did not call the function")
		}
		if v != 42 {
			t.Errorf("Return1() did not return the expected value")
		}
	case <-time.After(time.Millisecond):
		t.Errorf("Return1() took too long to return")
	}
}

func TestReturn2(t *testing.T) {
	called := false
	fn := func() (int, string) {
		called = true
		return 42, "hello"
	}
	c := goreturn.Return2(fn)
	select {
	case v := <-c:
		if !called {
			t.Errorf("Return2() did not call the function")
		}
		if v.V1 != 42 || v.V2 != "hello" {
			t.Errorf("Return2() did not return the expected values")
		}
	case <-time.After(time.Millisecond):
		t.Errorf("Return2() took too long to return")
	}
}

func TestReturn3(t *testing.T) {
	called := false
	fn := func() (int, string, bool) {
		called = true
		return 42, "hello", true
	}
	c := goreturn.Return3(fn)
	select {
	case v := <-c:
		if !called {
			t.Errorf("Return3() did not call the function")
		}
		if v.V1 != 42 || v.V2 != "hello" || !v.V3 {
			t.Errorf("Return3() did not return the expected values")
		}
	case <-time.After(time.Millisecond):
		t.Errorf("Return3() took too long to return")
	}
}
